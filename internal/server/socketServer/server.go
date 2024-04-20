package socketServer

import "C"
import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/sessionManager"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/transportClient"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/rtcSFU/types"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/socketClient"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/serverTypes"
	svc "github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/encoding/protojson"
	"log"
	"net/http"
	"sync"
	"time"
)

var Upgrader = websocket.Upgrader{
	//ReadBufferSize:  1024,
	//WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var _ serverTypes.ISocketServer = (*SocketServer)(nil)

type SocketServer struct {
	sync.Mutex
	Clients    map[string]*socketClient.SocketClient
	register   chan *socketClient.SocketClient
	unRegister chan *socketClient.SocketClient
	multicast  chan []byte
	broadcast  chan []byte

	sessionManager *sessionManager.SessionManager
	testConn       *webrtc.PeerConnection
	//eventListener *listener.SocketEvent
}

func NewSocketServer(iceServerURLs []string) *SocketServer {
	return &SocketServer{
		Clients:        make(map[string]*socketClient.SocketClient),
		register:       make(chan *socketClient.SocketClient),
		unRegister:     make(chan *socketClient.SocketClient),
		multicast:      make(chan []byte),
		broadcast:      make(chan []byte),
		sessionManager: sessionManager.NewSessionManager(),
	}
}

func (s *SocketServer) Start() {
	logx.Info("Starting ws server")
	for {
		select {
		case client := <-s.register:
			logx.Infof("New User is connecting : uuid: %v and name: %v ", client.UUID, client.Name)
			old, ok := s.add(client.UUID, client)

			//TODO: Send a welcome message?
			if ok {
				old.SendMessage(websocket.CloseMessage, []byte("Closed connection")) //TODO: send a close message to client
				old.Closed()                                                         //TODO: close the sending channel of old client
			}

		case client := <-s.unRegister:
			logx.Infof("User %v is leaving.", client.UUID)
			s.remove(client)
		case message := <-s.multicast:
			err := s.multicastMessageHandler(message)
			if err != nil {
				logx.Error("multicast error ", err)
				//send back to the sender?
			}

		case message := <-s.broadcast: //received protoBuffer message -> it need to be decoded
			err := s.broadcastMessageHandler(message)
			if err != nil {
				logx.Error(err)
			}
		}
	}
}

func (s *SocketServer) RegisterClient(client serverTypes.ISocketClient) {
	s.register <- (client).(*socketClient.SocketClient)
}

func (s *SocketServer) UnRegisterClient(client serverTypes.ISocketClient) {
	s.unRegister <- (client).(*socketClient.SocketClient)
}

func (s *SocketServer) BroadcastMessage(message []byte) {
	s.broadcast <- message
}

func (s *SocketServer) MulticastMessage(message []byte) {
	s.multicast <- message
}

func (s *SocketServer) GetOneClient(clientUUId string) (*socketClient.SocketClient, error) {
	for _, client := range s.Clients {
		if client.UUID == clientUUId {
			return client, nil
		}
	}
	return nil, errx.NewCustomErrCode(errx.CLIENT_NOT_FOUND)
}

// Internal function
func (s *SocketServer) sendGroupMessage(message *socket_message.Message, server *SocketServer, svcCtx *svc.ServiceContext) {
	//TODO: GET ALL GROUP MEMBER
	//TODO: Check if group is exist
	ctx := context.Background()
	group, err := svcCtx.DAO.FindOneGroupByUUID(ctx, message.ToUUID)
	if err != nil {
		logx.Error(err.Error())
		return
	}

	//TODO: Get All Group Members
	members, err := svcCtx.DAO.FindOneGroupMembers(ctx, group.Id)
	if err != nil {
		logx.Error(err.Error())
		return

	}
	for _, mem := range members {
		if mem.MemberInfo.Uuid == message.FromUUID && message.ContentType != variable.SYS {
			continue
		}

		conn, ok := server.Clients[mem.MemberInfo.Uuid]

		socketMessage := &socket_message.Message{
			MessageID:      message.MessageID,
			ReplyMessageID: message.ReplyMessageID,
			Avatar:         message.Avatar,
			FromUserName:   message.FromUserName,
			FromUUID:       message.ToUUID,   //From Group UUID
			ToUUID:         message.FromUUID, //To Member UUID
			Content:        message.Content,
			ContentType:    message.ContentType,
			MessageType:    message.MessageType,
			EventType:      message.EventType,
			UrlPath:        message.UrlPath,
			GroupName:      group.GroupName,
			GroupAvatar:    group.GroupAvatar,
			FileName:       message.FileName,
			FileSize:       message.FileSize,
		}

		messageBytes, err := json.MarshalIndent(socketMessage, "", "\t")
		if err != nil {
			logx.Error(err)
			continue
		}

		if !ok {
			logx.Infof("Group %v 's member %v is offline", message.ToUUID, mem.MemberInfo.Uuid)
			ctx := context.Background()
			_, err := variable.RedisConnection.RPush(ctx, mem.MemberInfo.Uuid, messageBytes).Result()
			if err != nil {
				logx.Error("offline message to redis err %s", err.Error())
			}
			continue

		}
		conn.SendMessage(websocket.BinaryMessage, messageBytes)
	}
}

// saveMessage, TEXT:Save directly and other types need to be store to FS
func (s *SocketServer) saveMessage(svcCtx *svc.ServiceContext, message *socket_message.Message) {
	//TODO : Save Message into db
	svcCtx.DAO.InsertOneMessage(context.Background(), message)
}

func (s *SocketServer) sendAcknowledgement(seqID string, toUUID string) {
	logx.Infof("Sending ack message with SeqID : %s", seqID)
	acknowledgement := &socket_message.Message{
		MessageID: seqID,
		ToUUID:    toUUID,
		EventType: variable.MSG_ACK,
	}
	ackMessage, err := json.MarshalIndent(acknowledgement, "", "\t")
	if err != nil {
		logx.Error(err)
		return
	}

	client, ok := s.Clients[toUUID]
	logx.Error(toUUID)
	if ok {
		client.SendMessage(websocket.BinaryMessage, ackMessage)
	} else {
		//TODO: Offline
		logx.Info("user is offline -> ack failed")
	}
}

func (s *SocketServer) onHandleNormalMessage(msg *socket_message.Message) error {
	logx.Info("Normal message handling")
	switch msg.ContentType {
	case variable.TEXT:
		fallthrough
	case variable.IMAGE:
		fallthrough
	case variable.FILE:
		fallthrough
	case variable.AUDIO:
		fallthrough
	case variable.VIDEO:
		fallthrough
	case variable.STORY:
		fallthrough
	case variable.SYS:
		fallthrough
	case variable.REPLY:
		fallthrough
	case variable.STICKER:
		fallthrough
	case variable.SHARED:
		//TODO: save message
		conn, ok := s.Clients[msg.FromUUID]
		if ok && msg.ContentType != variable.SYS {
			//TODO: No need to save system message
			s.saveMessage(conn.SvcCtx, msg)
		}

		bytes, err := protojson.Marshal(msg)
		if err != nil {
			logx.Error(err)
			return err
		}

		if msg.MessageType == variable.MESSAGE_TYPE_USERCHAT {
			client, ok := s.Clients[msg.ToUUID]
			if ok {
				client.SendMessage(websocket.BinaryMessage, bytes)
			} else {
				logx.Info("User not connect.")
				ctx := context.Background()
				_, err := variable.RedisConnection.RPush(ctx, msg.ToUUID, string(bytes)).Result()
				if err != nil {
					logx.Error("offline message to redis err %s", err.Error())
					return err
				}
			}

		} else if msg.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
			logx.Info("Sending Group Message")
			s.sendGroupMessage(msg, s, conn.SvcCtx)
		}

		if msg.ContentType != variable.SYS {
			//MARK: system message no need to ack??
			logx.Infof("Sending ack with seqId %s to userId :%s", msg.MessageID, msg.FromUUID)
			s.sendAcknowledgement(msg.MessageID, msg.FromUUID)
		}

		break
	default:
		logx.Error("Content type no supported")
		break

	}

	return nil
}

func (s *SocketServer) onHandleSFUMessage(msg *socket_message.Message) error {
	switch msg.EventType {
	case variable.SFU_EVENT_PRODUCER_SDP:
		var joinRoomData types.SFUConnectSessionReq
		jsonString := msg.Content //Can be a json string?
		userId := msg.FromUUID
		if err := jsonx.Unmarshal([]byte(jsonString), &joinRoomData); err != nil {
			logx.Error("json unmarshal error", err)
			return err
		}
		logx.Info(joinRoomData)
		sdpType := &types.Signaling{}
		if err := jsonx.Unmarshal([]byte(joinRoomData.SDPType), &sdpType); err != nil {
			logx.Error("json unmarshal error(sdp type)", err)
			return err
		}

		c, err := s.GetOneClient(userId)
		if err != nil {
			logx.Error("SocketClient not found")
			return err
		}
		logx.Info("Joining to session : ", joinRoomData.SessionId)

		//Find a session
		isNewRoom := false
		session, err := s.sessionManager.GetOneSession(joinRoomData.SessionId)
		if err != nil {
			logx.Info("Session not found")
			session = s.sessionManager.CreateOneSession(joinRoomData.SessionId, joinRoomData.CallType)
			isNewRoom = true
		}

		logx.Info("Current session Id : ", session.SessionId)

		//Create transport client
		tc := transportClient.NewTransportClient(userId, joinRoomData.SessionId, c)
		logx.Info("Created transport client for ", userId)

		session.AddNewSessionClient(userId, tc)

		err = tc.NewConnection(c.SvcCtx.Config.IceServer.Urls, sdpType, func(state webrtc.PeerConnectionState) {
			logx.Info("Connection State changed : ", state)
			switch state {
			case webrtc.PeerConnectionStateConnected:
				//Send a message to all client in that room

				ctx := context.Background()

				clients := session.GetSessionClients()
				sessionProducersList := make([]types.SFUProducerUserInfo, 0)

				currentUser, err := c.SvcCtx.DAO.FindOneUserByUUID(ctx, userId)
				if err != nil {
					logx.Error("Get User Info err : ", err)
					break
				}

				if isNewRoom {
					logx.Info("Sending a new session message.")
					group, err := c.SvcCtx.DAO.FindOneGroupByUUID(ctx, session.SessionId)
					if err != nil {
						logx.Error(err.Error())
						return
					}

					//TODO: Get All Group Members
					members, err := c.SvcCtx.DAO.FindOneGroupMembers(ctx, group.Id)
					if err != nil {
						logx.Error(err.Error())
						return

					}

					for _, member := range members {
						newMessage := &socket_message.Message{
							MessageID:    uuid.NewString(),
							FromUserName: currentUser.NickName,
							FromUUID:     session.SessionId,      //From Group UUID
							ToUUID:       member.MemberInfo.Uuid, //To Member UUID
							Content:      fmt.Sprintf("%s started a group %s's call", currentUser.NickName, joinRoomData.CallType),
							ContentType:  variable.TEXT,
							MessageType:  variable.MESSAGE_TYPE_GROUPCHAT,
							EventType:    variable.MESSAGE,
							GroupName:    group.GroupName,
							GroupAvatar:  group.GroupAvatar,
						}

						messageBytes, err := json.MarshalIndent(newMessage, "", "\t")
						if err != nil {
							logx.Error(err)
							continue
						}

						memberClient, err := s.GetOneClient(member.MemberInfo.Uuid)
						if err != nil {
							logx.Error("Client not exist")
							continue
						}

						memberClient.SendMessage(websocket.BinaryMessage, messageBytes)
					}
				}

				time.Sleep(2 * time.Second) //waiting for 2 sec to received all the track from producer.
				currentUserInfo := types.SFUProducerUserInfo{
					ProducerUserId:     currentUser.Uuid,
					ProducerUserName:   currentUser.NickName,
					ProducerUserAvatar: currentUser.Avatar,
				}

				//time.Sleep(1 * time.Second)
				for _, clientId := range clients {
					if clientId != userId {
						//sessionProducersList = append(sessionProducersList, clientId)
						//logx.Infof("Current user %s is Producer", clientId)
						curClient, err := s.GetOneClient(clientId)
						if err != nil {
							logx.Error("Get client err : ", err)
							continue
						}
						logx.Info("Getting user info")
						//TODO: Get Current client info
						ctx := context.Background()
						producerInfo, err := curClient.SvcCtx.DAO.FindOneUserByUUID(ctx, clientId)
						if err != nil {
							logx.Error("Get User Info err : ", err)
							continue
						}

						producerUserInfo := types.SFUProducerUserInfo{
							ProducerUserId:     producerInfo.Uuid,
							ProducerUserName:   producerInfo.NickName,
							ProducerUserAvatar: producerInfo.Avatar,
						}

						//add to session producers list
						sessionProducersList = append(sessionProducersList, producerUserInfo)

						resp := types.SfuNewProducerResp{
							SessionId:    session.SessionId,
							ProducerId:   currentUserInfo.ProducerUserId,
							ProducerInfo: currentUserInfo,
						}

						respStr, err := jsonx.Marshal(resp)
						if err != nil {
							logx.Error("resp marshal error : ", err)
							break
						}

						msg := &socket_message.Message{
							FromUUID:    userId,
							ToUUID:      clientId,
							Content:     string(respStr),
							ContentType: variable.SFU,
							EventType:   variable.SFU_EVENT_SEND_NEW_PRODUCER, //join room.
						}

						msgBytes, err := json.MarshalIndent(msg, "", "\t")
						if err != nil {
							logx.Error(err)
							break
						}
						logx.Info("Sending new producer in to client : ", clientId)
						curClient.SendMessage(websocket.BinaryMessage, msgBytes)
					}
				}

				//Response to producer.
				resp := types.SFUConnectSessionResp{
					SessionId:        session.SessionId,
					ProducerId:       msg.FromUUID,
					SessionProducers: sessionProducersList,
				}

				respStr, err := jsonx.Marshal(resp)
				if err != nil {
					logx.Error("resp marshal error : ", err)
					break
				}

				msg := &socket_message.Message{
					ToUUID:      msg.FromUUID,
					Content:     string(respStr),
					ContentType: variable.SFU,
					EventType:   variable.SFU_EVENT_PRODUCER_CONNECTED, //join room.
				}

				msgBytes, err := json.MarshalIndent(msg, "", "\t")
				if err != nil {
					logx.Error(err)
					break
				}
				c.SendMessage(websocket.BinaryMessage, msgBytes)

				break
			case webrtc.PeerConnectionStateDisconnected:
			case webrtc.PeerConnectionStateClosed:
				logx.Info("Connection State Change : Disconnected")
				//TODO: send a signal to all client in the session

				break
			case webrtc.PeerConnectionStateFailed:
				logx.Info("Connection State Change : Failed")
				//TODO: Close the connection when failed
				if err := tc.Close(); err != nil {
					log.Print(err)
				}
				break
			default:
				break
			}

		}, func(clientId string, track *webrtc.TrackLocalStaticRTP) {
			logx.Infof("Producer %s new track comes in : Kind %s", userId, track.Kind())
			if track == nil {
				logx.Error("Track is nil - new track listener")
				return
			}
			session.OnNewTrack(userId, track)
		})
		if err != nil {
			logx.Error("Create ans error ", err)
			return err
		}

		break
	case variable.SFU_EVENT_CONSUMER_SDP:
		//MARK: Same as Create?
		consumeReq := types.SFUConsumeProducerReq{}
		jsonString := msg.Content //Can be a json string?
		userId := msg.FromUUID

		c, err := s.GetOneClient(userId)
		if err != nil {
			logx.Error("Get SocketClient error : ", err)
			return err
		}

		if err := jsonx.Unmarshal([]byte(jsonString), &consumeReq); err != nil {
			logx.Error("Unmarshal error")
			return err
		}
		logx.Info("Receive offer for consumer : Producer Id : ", consumeReq.ProducerId)

		sdpType := &types.Signaling{}
		if err := jsonx.Unmarshal([]byte(consumeReq.SDPType), &sdpType); err != nil {
			logx.Error("json unmarshal error(sdp type)", err)
			return err
		}

		session, err := s.sessionManager.GetOneSession(consumeReq.SessionId)
		if err != nil {
			logx.Error("Get session error : ", err)
			return err
		}

		transC, err := session.GetTransportClient(userId)
		if err != nil {
			logx.Error("Get transportClient error : ", err)
			return err
		}

		producerClient, err := session.GetTransportClient(consumeReq.ProducerId)
		if err != nil {
			logx.Error(err)
			return err
		}

		producer, err := producerClient.GetProducer()
		if err != nil {
			logx.Error(err)
			return err
		}
		logx.Info("Consuming ID ,", producerClient.GetClientId())
		if err := transC.Consume(
			consumeReq.ProducerId,
			c.SvcCtx.Config.IceServer.Urls,
			sdpType,
			producer,
			func(state webrtc.PeerConnectionState) {
				logx.Info("(Consumer)Connection State changed : ", state)
				switch state {
				case webrtc.PeerConnectionStateConnected:
					logx.Info("(Consumer)Connection State Change : Connected")
					break
				case webrtc.PeerConnectionStateDisconnected:
				case webrtc.PeerConnectionStateClosed:
					logx.Info("(Consumer)Connection State Change : Disconnected")
					if err := transC.CloseConsumer(consumeReq.ProducerId); err != nil {
						logx.Error(err)
					}
					break
				case webrtc.PeerConnectionStateFailed:
					logx.Info("(Consumer)Connection State Change : Failed")
					//TODO: Close the connection when failed
					if err := transC.CloseConsumer(consumeReq.ProducerId); err != nil {

						logx.Error(err)
					}
					break
				default:
					break
				}
			}, func(clientId string, track *webrtc.TrackLocalStaticRTP) {
				if track == nil {
					return
				}
				logx.Infof("Consumer %s new track comes in , Kind %s", userId, track.Kind())
			}); err != nil {
			logx.Errorf("Consume %s error %s", consumeReq.ProducerId, err)
			return err
		}
		break
	case variable.SFU_EVENT_CONSUMER_ICE: //same logic as SFU_EVENT_PRODUCER_ICE but is Producer is false
		fallthrough
	case variable.SFU_EVENT_PRODUCER_ICE:
		//Add to ice candindate info into the peer connection that data is provided
		//MARK: Get All producer -> return a list of producerUserId
		iceCandidateReq := types.SFUSendIceCandidateReq{}
		jsonString := msg.Content //Can be a json string?
		iceCandidateType := types.IceCandidateType{}
		userId := msg.FromUUID
		//logx.Info("Received ice candidate from client")
		_, err := s.GetOneClient(userId)
		if err != nil {
			logx.Error("SocketClient not found")
			return err
		}

		if err := jsonx.Unmarshal([]byte(jsonString), &iceCandidateReq); err != nil {
			logx.Error("Unmarshal get producers request error : ", err)
			return err
		}
		if err := jsonx.Unmarshal([]byte(iceCandidateReq.IceCandidateType), &iceCandidateType); err != nil {
			logx.Error("Unmarshal IceCandidateType error : ", err)
			return err
		}

		session, err := s.sessionManager.GetOneSession(iceCandidateReq.SessionId)
		if err != nil {
			logx.Error(err)
			return err
		}
		transC, err := session.GetTransportClient(userId) //get current user - transport client obj
		if err != nil {
			logx.Error("Get Transport client error,", err)
			return err
		}
		if iceCandidateReq.IsProducer {
			if err := transC.ExchangeIceCandidateForProducer(iceCandidateReq.IceCandidateType); err != nil {
				logx.Error("Exchange ice candzidate for producer error,", err)
				return err
			}
		} else {
			if err := transC.ExchangeIceCandidateForConsumers(iceCandidateReq.ClientId, iceCandidateReq.IceCandidateType); err != nil {
				logx.Error("Exchange ice candidate for consumer error,", err)
				return err
			}
		}
		break
	case variable.SFU_EVENT_PRODUCER_CLOSE:
		closeConnReq := types.SFUCloseConnectionReq{}
		jsonString := msg.Content //Can be a json string?
		userId := msg.FromUUID

		c, err := s.GetOneClient(msg.FromUUID)
		if err != nil {
			logx.Error("SocketClient not found")
			return err
		}

		if err := jsonx.Unmarshal([]byte(jsonString), &closeConnReq); err != nil {
			logx.Error("Unmarshal get producers request error : ", err)
			return err
		}

		session, err := s.sessionManager.GetOneSession(closeConnReq.SessionId)
		if err != nil {
			logx.Error("Get one session error , ", err)
			return err
		}

		//Send a close message to all session client.
		//Disconnect consumer with userId
		for _, clientId := range session.GetSessionClients() {
			if clientId != userId {
				sessionClient, err := s.GetOneClient(clientId)
				if err != nil {
					logx.Error("Socket client error : ", err)
					continue
				}

				closeResp := types.SFUCloseConnectionResp{
					ProducerId: userId,
				}

				respStr, err := jsonx.Marshal(closeResp)
				if err != nil {
					logx.Error(err)
					continue
				}

				msg := &socket_message.Message{
					ToUUID:      clientId,
					Content:     string(respStr),
					ContentType: variable.SFU,
					EventType:   variable.SFU_EVENT_CONSUMER_CLOSE, // a producer is left
				}

				msgBytes, err := json.MarshalIndent(msg, "", "\t")
				if err != nil {
					logx.Error(err)
					return err
				}

				sessionClient.SendMessage(websocket.BinaryMessage, msgBytes)

				//Disconnect consumer
				transClient, err := session.GetTransportClient(clientId)
				if err != nil {
					logx.Error(err)
					return err
				}

				if err := transClient.CloseConsumer(userId); err != nil {
					logx.Error(err)
					return err
				}
			}
		}

		transC, err := session.GetTransportClient(msg.FromUUID)
		if err != nil {
			logx.Error("Get transport client , ", err)
			return err
		}

		//Close and disconnect all consumer connection.
		if err := transC.Close(); err != nil {
			logx.Error("Close connection error ,", err)
			return err
		}

		//Remove Current Client from Session
		session.RemoveSessionClient(msg.FromUUID)

		if session.IsEmpty() {
			logx.Info("Session is empty --- removing.....")
			s.sessionManager.RemoveOneSession(session.SessionId)
			ctx := context.Background()

			currentUser, err := c.SvcCtx.DAO.FindOneUserByUUID(ctx, userId)
			if err != nil {
				logx.Error("Get User Info err : ", err)
				break
			}

			group, err := c.SvcCtx.DAO.FindOneGroupByUUID(ctx, session.SessionId)
			if err != nil {
				logx.Error(err.Error())
				break
			}

			//TODO: Get All Group Members
			members, err := c.SvcCtx.DAO.FindOneGroupMembers(ctx, group.Id)
			if err != nil {
				logx.Error(err.Error())
				break

			}

			for _, member := range members {
				newMessage := &socket_message.Message{
					MessageID:    uuid.NewString(),
					FromUserName: currentUser.NickName,
					FromUUID:     session.SessionId,      //From Group UUID
					ToUUID:       member.MemberInfo.Uuid, //To Member UUID
					Content:      fmt.Sprintf("%s ended a group %s's call", currentUser.NickName, session.CallType),
					ContentType:  variable.TEXT,
					MessageType:  variable.MESSAGE_TYPE_GROUPCHAT,
					EventType:    variable.MESSAGE,
					GroupName:    group.GroupName,
					GroupAvatar:  group.GroupAvatar,
				}

				messageBytes, err := json.MarshalIndent(newMessage, "", "\t")
				if err != nil {
					logx.Error(err)
					continue
				}

				memberClient, err := s.GetOneClient(member.MemberInfo.Uuid)
				if err != nil {
					logx.Error("Client not exist")
					continue
				}

				memberClient.SendMessage(websocket.BinaryMessage, messageBytes)
			}

		}

		break

		//Send from server when producer connected?
	case variable.SFU_EVENT_PRODUCER_MEDIA_STATUS:
		logx.Info("Event : SFU_EVENT_PRODUCER_MEIDA_STATUS")
		mediaStatusReq := types.SFUProducerMediaStatusReq{}
		jsonString := msg.Content //Can be a json string?
		userId := msg.FromUUID

		if err := jsonx.Unmarshal([]byte(jsonString), &mediaStatusReq); err != nil {
			logx.Error("Unmarshal get producers request error : ", err)
			return err
		}

		session, err := s.sessionManager.GetOneSession(mediaStatusReq.SessionId)
		if err != nil {
			logx.Error("Get one session error , ", err)
			return err
		}

		//Send the message to all client in the room
		for _, clientId := range session.GetSessionClients() {
			if clientId != userId {
				sessionClient, err := s.GetOneClient(clientId)
				if err != nil {
					logx.Error("Socket client error : ", err)
					continue
				}

				msg := &socket_message.Message{
					ToUUID:      clientId,
					Content:     jsonString,
					ContentType: variable.SFU,
					EventType:   variable.SFU_EVENT_PRODUCER_MEDIA_STATUS, // a producer is left
				}

				msgBytes, err := json.MarshalIndent(msg, "", "\t")
				if err != nil {
					logx.Error(err)
					return err
				}

				sessionClient.SendMessage(websocket.BinaryMessage, msgBytes)
			}
		}
	default:
		logx.Infof("Event Type not support")
	}

	return nil
}

func (s *SocketServer) multicastMessageHandler(message []byte) error {
	var socketMessage socket_message.Message
	err := protojson.Unmarshal(message, &socketMessage)
	if err != nil {
		logx.Error(err)
		return err
	}

	if socketMessage.ToUUID != "" && socketMessage.ToUUID != "SFU" {
		return s.onHandleNormalMessage(&socketMessage)
	} else if socketMessage.ToUUID == "SFU" {
		return s.onHandleSFUMessage(&socketMessage)
	}
	return nil
}

func (s *SocketServer) broadcastMessageHandler(message []byte) error {
	var socketMessage socket_message.Message
	err := protojson.Unmarshal(message, &socketMessage)
	if err != nil {
		logx.Error(err)
		return err
	}
	//MARK: to all user....
	//MARK: just like system message..
	for _, client := range s.Clients {
		client.SendMessage(websocket.BinaryMessage, message)
	}
	return nil
}

func (s *SocketServer) add(uuid string, client *socketClient.SocketClient) (*socketClient.SocketClient, bool) {
	s.Lock()
	defer s.Unlock()
	old, ok := s.Clients[uuid]
	s.Clients[client.UUID] = client //add to our map
	if ok {
		return old, ok //TODO: Close the older connection
	}
	return nil, false
}

func (s *SocketServer) remove(client *socketClient.SocketClient) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Clients[client.UUID]; ok {
		delete(s.Clients, client.UUID)
	}
}

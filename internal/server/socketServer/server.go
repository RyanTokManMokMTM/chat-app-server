package socketServer

import (
	"context"
	"encoding/json"
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
	"net/http"
	"sync"
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

// Internal function
func (s *SocketServer) multicastMessageHandler(message []byte) error {
	var socketMessage socket_message.Message
	err := protojson.Unmarshal(message, &socketMessage)
	if err != nil {
		logx.Error(err)
		return err
	}

	if socketMessage.ToUUID != "" {
		//TODO: Send it to someone with a specific Uuid
		if socketMessage.ContentType >= variable.TEXT && socketMessage.ContentType <= variable.SYS {
			//TODO: save message
			conn, ok := s.Clients[socketMessage.FromUUID]
			if ok && socketMessage.ContentType != variable.SYS {
				//TODO: No need to save system message
				s.saveMessage(conn.SvcCtx, &socketMessage)
			}

			bytes, err := protojson.Marshal(&socketMessage)
			if err != nil {
				logx.Error(err)
				return err
			}

			if socketMessage.MessageType == variable.MESSAGE_TYPE_USERCHAT {

				client, ok := s.Clients[socketMessage.ToUUID]
				if ok {
					client.SendMessage(websocket.BinaryMessage, bytes)
				} else {
					ctx := context.Background()
					_, err := variable.RedisConnection.RPush(ctx, socketMessage.ToUUID, string(bytes)).Result()
					if err != nil {
						logx.Error("offline message to redis err %s", err.Error())
						return err
					}
				}

			} else if socketMessage.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
				logx.Info("Sending Group Message")
				s.sendGroupMessage(&socketMessage, s, conn.SvcCtx)
			}

		}
	} else {
		//Or communicate with server?
		switch socketMessage.EventType {
		case variable.SFU_EVENT_CONNECT:
			var joinRoomData types.SFUJoinRoomReq
			jsonString := socketMessage.Content //Can be a json string?
			userId := socketMessage.FromUUID
			if err := jsonx.Unmarshal([]byte(jsonString), &joinRoomData); err != nil {
				logx.Error("json unmarshal error", err)
				break
			}

			c, err := s.GetOneClient(userId)
			if err != nil {
				logx.Error("SocketClient not found")
				break
			}
			logx.Info("Joining to session : ", joinRoomData.SessionId)

			//Find a session
			session, err := s.sessionManager.GetOneSession(joinRoomData.SessionId)
			if err != nil {
				logx.Error("Session not found")
				session = s.sessionManager.CreateOneSession(joinRoomData.SessionId)
			}

			logx.Info("Current session Id : ", session.SessionId)

			//Create transport client
			tc := transportClient.NewTransportClient(userId, c)
			logx.Info("Created transport client for ", userId)

			//Create the SFU connection
			err = tc.NewConnection(c.SvcCtx.Config.IceServer.Urls, joinRoomData.Offer, func(state webrtc.PeerConnectionState) {
				switch state {
				case webrtc.PeerConnectionStateNew:
					logx.Info("Connection State Change : New Connection")
					break
				case webrtc.PeerConnectionStateConnecting:
					logx.Info("Connection State Change : Connecting")
					break
				case webrtc.PeerConnectionStateConnected:
					logx.Info("Connection State Change : Connected")
					//TODO: send a signal to all client in the session
					clients := session.GetSessionClients()
					for _, c := range clients {
						if c != userId {
							receiver, err := s.GetOneClient(c)
							if err != nil {
								logx.Error(err)
								continue
							}

							resp := types.SFUResponse{
								Data: userId,
							}

							respStr, err := jsonx.Marshal(resp)
							if err != nil {
								logx.Error("resp marshal error : ", err)
								break
							}

							msg := socket_message.Message{
								ToUUID:      socketMessage.FromUUID,
								Content:     string(respStr),
								ContentType: variable.SFU,
								EventType:   variable.SFU_EVENT_SEND_NEW_PRODUCER,
							}

							msgBytes, err := json.MarshalIndent(msg, "", "\t")
							if err != nil {
								logx.Error(err)
								break
							}

							receiver.SendMessage(websocket.BinaryMessage, msgBytes)

						}
					}
					break
				case webrtc.PeerConnectionStateDisconnected:
				case webrtc.PeerConnectionStateClosed:
					logx.Info("Connection State Change : Disconnected")
					//TODO: send a signal to all client in the session
					clients := session.GetSessionClients()
					for _, c := range clients {
						if c != userId {
							receiver, err := s.GetOneClient(c)
							if err != nil {
								logx.Error(err)
								continue
							}

							resp := types.SFUResponse{
								Data: userId,
							}

							respStr, err := jsonx.Marshal(resp)
							if err != nil {
								logx.Error("resp marshal error : ", err)
								break
							}

							msg := socket_message.Message{
								ToUUID:      socketMessage.FromUUID,
								Content:     string(respStr),
								ContentType: variable.SFU,
								EventType:   variable.SFU_EVENT_SEND_PRODUCER_CLOSE, //join room.
							}

							msgBytes, err := json.MarshalIndent(msg, "", "\t")
							if err != nil {
								logx.Error(err)
								break
							}

							receiver.SendMessage(websocket.BinaryMessage, msgBytes)

						}
					}
					break
				case webrtc.PeerConnectionStateFailed:
					logx.Info("Connection State Change : Failed")
					break
				}
			})
			if err != nil {
				logx.Error("Create ans error ", err)
				break
			}

			break
		case variable.SFU_EVENT_CONSUM:
			//MARK: Same as Create?
			consumeReq := types.SFUConsumeReq{}
			jsonString := socketMessage.Content //Can be a json string?
			userId := socketMessage.FromUUID

			c, err := s.GetOneClient(userId)
			if err != nil {
				logx.Error("Get SocketClient error : ", err)
				break
			}

			if err := jsonx.Unmarshal([]byte(jsonString), &consumeReq); err != nil {
				logx.Error("Unmarshal error")
				break
			}

			session, err := s.sessionManager.GetOneSession(consumeReq.SessionId)
			if err != nil {
				logx.Error("Get session error : ", err)
				break
			}

			transC, err := session.GetTransportClient(userId)
			if err != nil {
				logx.Error("Get transportClient error : ", err)
				break
			}

			if err := transC.Consume(consumeReq.ConsumerId, c.SvcCtx.Config.IceServer.Urls, consumeReq.Offer, func(state webrtc.PeerConnectionState) {
				logx.Error("consumer state : ", state)
			}); err != nil {
				logx.Errorf("Consume %s error %s", consumeReq.ConsumerId, err)
				break
			}
			break
		case variable.SFU_EVENT_GET_PRODUCERS:
			//MARK: Get All producer -> return a list of producerUserId
			getProducersReq := types.SFUGetProducerReq{}
			jsonString := socketMessage.Content //Can be a json string?

			c, err := s.GetOneClient(socketMessage.FromUUID)
			if err != nil {
				logx.Error("SocketClient not found")
				break
			}

			if err := jsonx.Unmarshal([]byte(jsonString), &getProducersReq); err != nil {
				logx.Error("Unmarshal get producers request error : ", err)
				break
			}

			session, err := s.sessionManager.GetOneSession(getProducersReq.SessionId)
			if err != nil {
				logx.Error(err)
				break
			}

			producersList := session.GetSessionClients()
			respList, err := jsonx.Marshal(producersList)

			resp := types.SFUResponse{
				Data: string(respList),
			}

			respStr, err := jsonx.Marshal(resp)
			if err != nil {
				logx.Error("resp marshal error : ", err)
				break
			}

			msg := socket_message.Message{
				ToUUID:      socketMessage.FromUUID,
				Content:     string(respStr),
				ContentType: variable.SFU,
				EventType:   variable.SFU_EVENT_GET_PRODUCERS, //join room.
			}

			msgBytes, err := json.MarshalIndent(msg, "", "\t")
			if err != nil {
				logx.Error(err)
				break
			}

			c.SendMessage(websocket.BinaryMessage, msgBytes)

			break
		case variable.SFU_EVENT_ICE:
			//Add to ice candindate info into the peer connection that data is provided
			//MARK: Get All producer -> return a list of producerUserId
			iceCandindateReq := types.SFUIceCandindateReq{}
			jsonString := socketMessage.Content //Can be a json string?

			_, err := s.GetOneClient(socketMessage.FromUUID)
			if err != nil {
				logx.Error("SocketClient not found")
				break
			}

			if err := jsonx.Unmarshal([]byte(jsonString), &iceCandindateReq); err != nil {
				logx.Error("Unmarshal get producers request error : ", err)
				break
			}

			session, err := s.sessionManager.GetOneSession(iceCandindateReq.SessionId)
			if err != nil {
				logx.Error(err)
				break
			}

			transC, err := session.GetTransportClient(socketMessage.FromUUID) //get current user - transport client obj
			if err != nil {
				logx.Error("Get Transport client error,", err)
				break
			}

			if iceCandindateReq.IsProducer {
				if err := transC.ExchangeIceCandindateForProducer(iceCandindateReq.IceCandidate); err != nil {
					logx.Error("Exchange ice candindate for producer error,", err)
					break
				}
			} else {
				if err := transC.ExchangeIceCandindateForConsumers(iceCandindateReq.ToClientId, iceCandindateReq.IceCandidate); err != nil {
					logx.Error("Exchange ice candindate for consumer error,", err)
					break
				}
			}
			break
		case variable.SFU_EVENT_CLOSE:
			closeConnReq := types.SFUCloseReq{}
			jsonString := socketMessage.Content //Can be a json string?
			userId := socketMessage.FromUUID

			_, err := s.GetOneClient(socketMessage.FromUUID)
			if err != nil {
				logx.Error("SocketClient not found")
				break
			}

			if err := jsonx.Unmarshal([]byte(jsonString), &closeConnReq); err != nil {
				logx.Error("Unmarshal get producers request error : ", err)
				break
			}

			session, err := s.sessionManager.GetOneSession(closeConnReq.SessionId)
			if err != nil {
				logx.Error("Get one session error , ", err)
				break
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

					closeResp := types.SFUProducerClosedResp{
						ProducerId: userId,
					}

					respStr, err := jsonx.Marshal(closeResp)
					if err != nil {
						logx.Error(err)
						continue
					}

					msg := socket_message.Message{
						ToUUID:      clientId,
						Content:     string(respStr),
						ContentType: variable.SFU,
						EventType:   variable.SFU_EVENT_CLOSE,
					}

					msgBytes, err := json.MarshalIndent(msg, "", "\t")
					if err != nil {
						logx.Error(err)
						break
					}

					sessionClient.SendMessage(websocket.BinaryMessage, msgBytes)

					//Disconnect consumer
					transClient, err := session.GetTransportClient(clientId)
					if err != nil {
						logx.Error(err)
						break
					}

					if err := transClient.CloseConsumer(userId); err != nil {
						logx.Error(err)
						break
					}
				}
			}

			transC, err := session.GetTransportClient(socketMessage.FromUUID)
			if err != nil {
				logx.Error("Get transport client , ", err)
				break
			}

			//Close and disconnect all consumer connection.
			if err := transC.Close(); err != nil {
				logx.Error("Close connection error ,", err)
				break
			}

			break

		default:
			logx.Infof("Event Type not support")
		}

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

func (s *SocketServer) GetOneClient(clientUUId string) (*socketClient.SocketClient, error) {
	for _, client := range s.Clients {
		if client.UUID == clientUUId {
			return client, nil
		}
	}
	return nil, errx.NewCustomErrCode(errx.CLIENT_NOT_FOUND)
}

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

		socketMessage := socket_message.Message{
			Avatar:       message.Avatar,
			FromUserName: message.FromUserName,
			FromUUID:     message.ToUUID,   //From Group UUID
			ToUUID:       message.FromUUID, //To Member UUID
			Content:      message.Content,
			ContentType:  message.ContentType,
			MessageType:  message.MessageType,
			EventType:    message.EventType,
			UrlPath:      message.UrlPath,
			GroupName:    group.GroupName,
			GroupAvatar:  group.GroupAvatar,
			FileName:     message.FileName,
			FileSize:     message.FileSize,
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

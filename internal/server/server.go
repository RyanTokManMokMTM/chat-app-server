package server

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/serialization"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/sfuType"
	svc "github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	//ReadBufferSize:  1024,
	//WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketServer struct {
	sync.Mutex
	Clients    map[string]*SocketClient
	Register   chan *SocketClient
	UnRegister chan *SocketClient
	Multicast  chan []byte
	Broadcast  chan []byte

	sfuRooms *SFURooms
	//eventListener *listener.SocketEvent
}

func NewSocketServer(iceServerURLs []string) *SocketServer {
	return &SocketServer{
		Clients:    make(map[string]*SocketClient),
		Register:   make(chan *SocketClient),
		UnRegister: make(chan *SocketClient),
		Multicast:  make(chan []byte),
		Broadcast:  make(chan []byte),
		sfuRooms:   NewSFURooms(iceServerURLs),
	}
}

func (s *SocketServer) Start() {
	logx.Info("Starting ws server")
	for {
		select {
		case client := <-s.Register:
			logx.Infof("New User is connecting : uuid: %v and name: %v ", client.UUID, client.Name)
			old, ok := s.Add(client.UUID, client)

			//TODO: Send a welcome message?
			if ok {
				old.conn.WriteMessage(websocket.CloseMessage, nil) //TODO: send a close message to client
				old.Closed()                                       //TODO: close the sending channel of old client
			}

		case client := <-s.UnRegister:
			logx.Infof("User %v is leaving.", client.UUID)
			s.Remove(client)
		case message := <-s.Multicast:
			err := s.multicastMessageHandler(message)
			if err != nil {
				logx.Error("multicast error ", err)
				//send back to the sender?
			}

		case message := <-s.Broadcast: //received protoBuffer message -> it need to be decoded
			err := s.broadcastMessageHandler(message)
			if err != nil {
				logx.Error(err)
			}
		}
	}
}

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
				saveMessage(conn.svcCtx, &socketMessage)
			}

			bytes, err := protojson.Marshal(&socketMessage)
			if err != nil {
				logx.Error(err)
				return err
			}

			if socketMessage.MessageType == variable.MESSAGE_TYPE_USERCHAT {

				client, ok := s.Clients[socketMessage.ToUUID]
				if ok {
					client.sendChannel <- bytes
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
				sendGroupMessage(&socketMessage, s, conn.svcCtx)
			}

		}
	} else {
		logx.Error("Receiver is empty?")
		//Or communicate with server?
		switch socketMessage.EventType {
		case variable.SFU_JOIN:
			roomUUID := socketMessage.Content //Can be a json string?
			var joinData sfuType.SFUJoinEventDataReq
			if err := serialization.DecodeJSON([]byte(socketMessage.Content), joinData); err != nil {
				logx.Error("join data decode error ", err)
			}
			//MARK: For a new connection has joined/create the channel.
			/*
				STEP:
				1. check if the svc exists
				2.

			*/
			room, err := s.sfuRooms.FindOneRoom(roomUUID)
			if errors.Is(err, errx.NewCustomErrCode(errx.SFU_ROOM_NOT_FOUND)) {
				room = s.sfuRooms.CreateNewRoom(roomUUID)
			}

			//NARK get client data...

			client, err := s.GetOneClient(socketMessage.FromUUID)
			if err != nil {
				logx.Error("Find client error ", err)
				break
			}
			answer, err := room.NewConnection(s.sfuRooms.IceUrls, client, joinData.Offer, func(peerConn *webrtc.PeerConnection, remote *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
				if receiver != nil && receiver.Tracks()[0] != nil {
					//peerConn.AddTrack(webrtc.NewTrackLocalStaticSample())
					stream, err := webrtc.NewTrackLocalStaticSample(remote.Codec().RTPCodecCapability, remote.ID(), remote.StreamID())
					if err != nil {
						logx.Error("Create stream local err : ", err)
						return
					}
					peerConn.AddTrack(stream)
				}
			})
			if err != nil {
				logx.Error("Create Peer connection error ", err)
				break
			}

			answerJSON, err := serialization.EncodeJSON(answer)
			if err != nil {
				logx.Error("Encode json err ", err)
				break
			}

			_ = sfuType.SFUJoinEventDataResp{
				Answer: string(answerJSON),
			}

			break
		case variable.SFU_OFFER:
			//MARK: Same as Create?
			break
		case variable.SFU_ANSWER:
			break
		case variable.SFU_CONSUM:
			break
		case variable.SFU_CONSUM_ICE:
			break
		case variable.SFU_CLOSE:
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
		select {
		case client.sendChannel <- message: //TODO: IF sendChannel is not available -> close and remove from the map
		default:
			close(client.sendChannel)
			s.Remove(client)
		}

	}
	return nil
}

func (s *SocketServer) Add(uuid string, client *SocketClient) (*SocketClient, bool) {
	s.Lock()
	defer s.Unlock()
	old, ok := s.Clients[uuid]
	s.Clients[client.UUID] = client //add to our map
	if ok {
		return old, ok //TODO: Close the older connection
	}
	return nil, false
}

func (s *SocketServer) Remove(client *SocketClient) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.Clients[client.UUID]; ok {
		delete(s.Clients, client.UUID)
	}
}

func (s *SocketServer) GetOneClient(clientUUId string) (*SocketClient, error) {
	for _, client := range s.Clients {
		if client.UUID == clientUUId {
			return client, nil
		}
	}
	return nil, errx.NewCustomErrCode(errx.CLIENT_NOT_FOUND)
}

func sendGroupMessage(message *socket_message.Message, server *SocketServer, svcCtx *svc.ServiceContext) {
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
		conn.sendChannel <- messageBytes
	}
}

// saveMessage, TEXT:Save directly and other types need to be store to FS
func saveMessage(svcCtx *svc.ServiceContext, message *socket_message.Message) {
	//TODO : Save Message into db
	svcCtx.DAO.InsertOneMessage(context.Background(), message)
}

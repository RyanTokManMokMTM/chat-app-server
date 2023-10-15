package server

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
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
	Broadcast  chan []byte

	isDone chan struct{}
}

func NewSocketServer() *SocketServer {
	return &SocketServer{
		Clients:    make(map[string]*SocketClient),
		Register:   make(chan *SocketClient),
		UnRegister: make(chan *SocketClient),
		Broadcast:  make(chan []byte),
		isDone:     make(chan struct{}),
	}
}

func (s *SocketServer) Start() {
	logx.Info("Starting ws router")
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

			logx.Info("Done")
			s.isDone <- struct{}{}

		case client := <-s.UnRegister:
			logx.Infof("User %v is leaving.", client.UUID)
			s.Remove(client)

		case message := <-s.Broadcast: //received protoBuffer message -> it need to be decoded
			var socketMessage socket_message.Message
			err := protojson.Unmarshal(message, &socketMessage)
			if err != nil {
				logx.Error(err)
				continue
			}
			//return

			//TODO: Send To Who?
			//TODO: Send To Nobody , it means broadcast to a specific user/group
			if socketMessage.ToUUID != "" {
				//TODO: Send it to someone with a specific Uuid
				if socketMessage.ContentType >= variable.TEXT && socketMessage.ContentType <= variable.SHARED {
					//TODO: save message
					conn, ok := s.Clients[socketMessage.FromUUID]
					if ok && socketMessage.ContentType != variable.SYS {
						//TODO: No need to save system message
						saveMessage(conn.svcCtx, &socketMessage)
					}

					bytes, err := protojson.Marshal(&socketMessage)
					if err != nil {
						logx.Error(err)
						continue
					}

					if socketMessage.MessageType == variable.MESSAGE_TYPE_USERCHAT {
						logx.Error(s)
						//0x105a1b020
						if client, ok := s.Clients[socketMessage.ToUUID]; ok {
							client.sendChannel <- bytes
						} else {
							logx.Infof("user %s is not online ", socketMessage.ToUUID)
							ctx := context.Background()
							_, err := variable.RedisConnection.RPush(ctx, socketMessage.ToUUID, string(bytes)).Result()
							if err != nil {
								logx.Error("offline message to redis err %s", err.Error())
							}
						}

					} else if socketMessage.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
						logx.Info("Sending Group Message")
						sendGroupMessage(&socketMessage, s, conn.svcCtx)
					}

					if socketMessage.ContentType != variable.SYS {
						//MARK: system message no need to ack??
						sendAcknowledgement(socketMessage.MessageID, s, socketMessage.FromUUID)
					}
				}
			} else {
				//TODO: Send To All User that are online
				for _, client := range s.Clients {
					select {
					case client.sendChannel <- message: //TODO: IF sendChannel is not available -> close and remove from the map
					default:
						close(client.sendChannel)
						s.Remove(client)
					}
				}
			}

		}
	}
}

func (s *SocketServer) Add(uuid string, client *SocketClient) (*SocketClient, bool) {
	s.Lock()
	defer s.Unlock()
	old, ok := s.Clients[uuid]
	s.Clients[client.UUID] = client //add to our map

	logx.Infof("Added user %s", client.UUID)
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
			MessageID:      message.MessageID,
			ReplyMessageID: message.ReplyMessageID,
			Avatar:         message.Avatar,
			FromUserName:   message.FromUserName,
			FromUUID:       message.ToUUID,   //From Group UUID
			ToUUID:         message.FromUUID, //To Member UUID
			Content:        message.Content,
			ContentType:    message.ContentType,
			MessageType:    message.MessageType,
			Type:           message.Type,
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
		conn.sendChannel <- messageBytes
	}
}

// saveMessage, TEXT:Save directly and other types need to be store to FS
func saveMessage(svcCtx *svc.ServiceContext, message *socket_message.Message) {
	svcCtx.DAO.InsertOneMessage(context.Background(), message)
}

func sendAcknowledgement(seqID string, server *SocketServer, toUUID string) {
	logx.Infof("Sending ack message with SeqID : %s", seqID)
	acknowledgement := socket_message.Message{
		MessageID: seqID,
		ToUUID:    toUUID,
		Type:      variable.MSG_ACK,
	}
	ackMessage, err := json.MarshalIndent(acknowledgement, "", "\t")
	if err != nil {
		logx.Error(err)
		return
	}

	client, ok := server.Clients[toUUID]
	logx.Error(toUUID)
	if ok {
		client.sendChannel <- ackMessage
	} else {
		//TODO: Offline
		logx.Info("user is offline -> ack failed")
	}
}

package server

import (
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/ryantokmanmok/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmok/chat-app-server/common/errx"
	"github.com/ryantokmanmok/chat-app-server/common/variable"
	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	socket_message "github.com/ryantokmanmok/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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
}

func NewSocketServer() *SocketServer {
	return &SocketServer{
		Clients:    make(map[string]*SocketClient),
		Register:   make(chan *SocketClient),
		UnRegister: make(chan *SocketClient),
		Broadcast:  make(chan []byte),
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

		case message := <-s.Broadcast: //received protoBuffer message -> it need to be decoded
			logx.Infof("A new message need to be broadcast : %v ", message)
			//decode back to protoBuffer type -Message
			var socketMessage socket_message.Message
			err := proto.UnmarshalMerge(message, &socketMessage)
			if err != nil {
				logx.Error("Unmarshal message error : %v ", err)
			}

			//TODO: Send To Who?
			//TODO: Send To Nobody , it means broadcast to all user who is connected to the server
			if socketMessage.ToUUID != "" {
				//TODO: Send it to someone with a specific UUID
				if socketMessage.MessageType >= variable.TEXT && socketMessage.MessageType <= variable.VIDEO {
					//TODO: save message
					_, ok := s.Clients[socketMessage.FromUUID]
					if ok {
						saveMessage(&socketMessage)
					}

					if socketMessage.MessageType == variable.MESSAGE_TYPE_USERCHAT {
						//TODO: Send Group Message
						client, ok := s.Clients[socketMessage.ToUUID]
						if ok {
							bytes, err := proto.Marshal(&socketMessage)
							if err != nil {
								client.sendChannel <- bytes
							}
						}
					} else if socketMessage.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
						//TODO: Send Peer to Peer Message
						sendGroupMessage(&socketMessage, s)
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

func sendGroupMessage(message *socket_message.Message, server *SocketServer) {
	//TODO: GET ALL GROUP MEMBER
	//TODO: SEND TO ALL ONLINE USER
}

// saveMessage, TEXT:Save directly and other types need to be store to FS
func saveMessage(message *socket_message.Message) {}

func ServeWS(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request, wsServer *SocketServer) {
	//TODO: Upgrade http to websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Websocket upgrade error"))
		return
	}
	//TODO : Get UserID from Context
	userID := ctxtool.GetUserIDFromCTX(r.Context())
	//TODO : Find User Info from DB
	u, err := svcCtx.DAO.FindOneUser(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errx.NewCustomErrCode(errx.USER_NOT_EXIST).GetMessage()))
		return
	}

	client := NewSocketClient(u.Uuid, u.NickName, conn, wsServer)
	wsServer.Register <- client

	go client.ReadLoop()
	go client.WriteLoop()
}

package server

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/ryantokmanmok/chat-app-server/common/variable"
	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	socket_message "github.com/ryantokmanmok/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
	"time"
)

type SocketClient struct {
	once        sync.Once
	UUID        string
	Name        string
	conn        *websocket.Conn
	sendChannel chan []byte
	isClose     chan struct{}
	server      *SocketServer
	svcCtx      *svc.ServiceContext
}

func NewSocketClient(uuid string, name string, conn *websocket.Conn, server *SocketServer, svcCtx *svc.ServiceContext) *SocketClient {
	return &SocketClient{
		UUID:        uuid,
		Name:        name,
		conn:        conn,
		sendChannel: make(chan []byte),
		isClose:     make(chan struct{}),
		server:      server,
		svcCtx:      svcCtx,
	}
}

// ReadLoop from client via its channel/socket
func (c *SocketClient) ReadLoop() {

	defer func() {
		c.server.UnRegister <- c
		c.conn.Close() //TODO: close the connection
	}()

	//TODO: set read init time
	c.conn.SetReadDeadline(time.Now().Add(time.Second * variable.ReadWait)) //TODO: Need to read any message before deadline
	c.conn.SetReadLimit(variable.ReadLimit)                                 //TODO: Size of a message
	c.conn.SetPongHandler(func(appData string) error {                      //TODO: Received a ping message from client, we need to handle it by setting a handle function
		logx.Info(appData)
		return nil
	})

	for {
		//read and message
		//client may send a ping signal
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			logx.Error("Read Client Message Error")
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.Error(err)
			}
			break
		}

		//TODO: Unmarshal message to prototype message
		var socketMessage socket_message.Message
		err = jsonpb.UnmarshalString(string(message), &socketMessage)
		if err != nil {
			logx.Error(err)
		}

		if socketMessage.Type == variable.HEAT_BEAT {
			//ping message -> send pong message
			logx.Info("Get a ping message from client")
			msg := socket_message.Message{
				Content: variable.PONG_MESSAGE,
				Type:    variable.HEAT_BEAT,
			}

			pongBytes, err := proto.Marshal(&msg)
			if err != nil {
				logx.Error("marshal message error : %v", err)
			}

			c.conn.WriteMessage(websocket.BinaryMessage, pongBytes) //TODO: Send it to client directly
		} else {
			//normal message
			c.server.Broadcast <- message
		}
	}
}

func (c *SocketClient) WriteLoop() {
	t := time.NewTicker(time.Second * variable.ReadWait * 9 / 10)
	defer func() {
		c.server.UnRegister <- c
		c.conn.Close()
	}()

	for {
		select {
		case data, ok := <-c.sendChannel:
			//send received message to user
			if !ok {
				break
			}

			//TODO: Set WriteDeadLine
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * variable.WriteWait))
			logx.Info("received a message")

			w, err := c.conn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				logx.Error(err)
				return
			}

			_, _ = w.Write(data)
			n := len(c.sendChannel)
			for i := 0; i < n; i++ {
				_, _ = w.Write(data)
			}

			if err := w.Close(); err != nil {
				logx.Error("writer close err :%v", err)
				return
			}

		case <-t.C:
			c.conn.SetWriteDeadline(time.Now().Add(time.Second * variable.WriteWait))
			logx.Info("received a ping ticker")
			//TODO: Send Ticket message
			msg := socket_message.Message{
				Content: variable.PING_MESSAGE,
				Type:    variable.HEAT_BEAT,
			}

			bytes, _ := proto.Marshal(&msg)

			c.server.Broadcast <- bytes
		case <-c.isClose:
			logx.Info("received a connection closed signal and user is disconnected")
			return
		}
	}
}

func (c *SocketClient) Closed() {
	c.once.Do(func() {
		logx.Info("client close the connection")
		c.isClose <- struct{}{}
	})
}

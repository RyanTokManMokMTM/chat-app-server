package server

import "github.com/gorilla/websocket"

type SocketClient struct {
	UUID string
	Name string
	Conn *websocket.Conn
}

func NewSocketClient(uuid string, name string, conn *websocket.Conn) *SocketClient {
	return &SocketClient{
		UUID: uuid,
		Name: name,
		Conn: conn,
	}
}

func (c *SocketClient) Rec()  {}
func (c *SocketClient) Send() {}

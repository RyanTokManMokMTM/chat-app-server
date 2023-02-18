package server

import (
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

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
	logx.Info("Starting websocket server")
	for {
		select {
		case client := <-s.Register:
			logx.Infof("New User is connecting : uuid: %v and name: %v ", client.UUID, client.Name)

		case client := <-s.UnRegister:
			logx.Infof("User %v is leaving.", client.UUID)

		case message := <-s.Broadcast: //received protoBuffer message -> it need to be decoded
			logx.Infof("A new message need to be broadcast : %v ", message)

		}
	}
}

func (s *SocketServer) Add() {}

func (s *SocketServer) Remove() {}

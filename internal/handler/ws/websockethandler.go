package ws

import (
	"github.com/ryantokmanmok/chat-app-server/internal/server"
	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	"net/http"
)

var socket *server.SocketServer

func WebSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	if socket == nil {
		socket = server.NewSocketServer()
		go socket.Start()
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.ServeWS(svcCtx, w, r, socket)
	}
}

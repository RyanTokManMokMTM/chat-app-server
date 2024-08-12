package ws

import (
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/server/socketServer"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/serverTypes"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/logic/ws"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/svc"
)

var AppSocket serverTypes.ISocketServer

func ConnectWsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	if AppSocket == nil {
		AppSocket = socketServer.NewSocketServer(svcCtx.Config.IceServer.Urls, svcCtx.RedisClient)
		go AppSocket.Start()
	}

	return func(w http.ResponseWriter, r *http.Request) {
		l := ws.NewConnectWsLogic(r.Context(), svcCtx)
		l.ConnectWs(w, r, AppSocket)
	}
}

package ws

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/server/socketServer"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/serverTypes"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/socket-proto"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/variable"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/logic/ws"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/svc"
)

var Socket serverTypes.ISocketServer

func ConnectWsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	if Socket == nil {
		Socket = socketServer.NewSocketServer(svcCtx.Config.IceServer.Urls)
		go Socket.Start()
	}

	return func(w http.ResponseWriter, r *http.Request) {
		l := ws.NewConnectWsLogic(r.Context(), svcCtx)
		l.ConnectWs(w, r, Socket)
	}
	//func(w http.ResponseWriter, r *http.Request) {
	//	l := ws.NewConnectWsLogic(r.Context(), svcCtx)
	//	err := l.ConnectWs()
	//	httpx.Ok(w)
	//}
}

func SendGroupSystemNotification(FromUUID, groupUUID, content string) {
	logx.Info("Sending notification")
	msg := &socket_message.Message{
		MessageID:   uuid.New().String(),
		FromUUID:    FromUUID,
		ToUUID:      groupUUID,
		ContentType: variable.SYS,
		Content:     content,
		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
		EventType:   variable.MESSAGE,
	}
	messageBytes, err := json.MarshalIndent(msg, "", "\t")
	if err != nil {
		logx.Error(err)
		return
	}

	logx.Info(msg)
	Socket.MulticastMessage(messageBytes)
}

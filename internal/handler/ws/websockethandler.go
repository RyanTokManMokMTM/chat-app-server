package ws

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

var Socket *server.SocketServer

func WebSocketHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	if Socket == nil {
		Socket = server.NewSocketServer()
		go Socket.Start()
	}
	return func(w http.ResponseWriter, r *http.Request) {
		server.ServeWS(svcCtx, w, r, Socket)
	}
}

func SendGroupSystemNotification(FromUUID, groupUUID, content string) {
	msg := socket_message.Message{
		MessageID:   uuid.New().String(),
		FromUUID:    FromUUID,
		ToUUID:      groupUUID,
		ContentType: variable.SYS,
		Content:     content,
		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
		Type:        variable.MESSAGE,
	}
	logx.Info(msg)
	messageBytes, err := json.MarshalIndent(msg, "", "\t")
	if err != nil {
		logx.Error(err)
		return
	}

	logx.Info(messageBytes)
	Socket.Broadcast <- messageBytes
}

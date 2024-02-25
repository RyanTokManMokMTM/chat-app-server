package server

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/socketClient"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/server/socketServer"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/serverTypes"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"time"
)

func ServeWS(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request, wsServer serverTypes.SocketServerIf) {
	//TODO: Upgrade http to websocket
	conn, err := socketServer.Upgrader.Upgrade(w, r, nil)
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

	websocketServer := wsServer.(*socketServer.SocketServer)
	client := socketClient.NewSocketClient(u.Uuid, u.NickName, conn, websocketServer, svcCtx)
	wsServer.RegisterClient(client)

	go client.ReadLoop()
	go client.WriteLoop()

	go func() {
		ctx := context.Background()

		//we need to create a connection for each user?
		len, err := variable.RedisConnection.LLen(ctx, u.Uuid).Result()
		if err != nil {
			logx.Error("getting Redis length err ", err)
			return
		}

		messages, err := svcCtx.RedisClient.LRange(ctx, u.Uuid, 0, len).Result()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logx.Errorf("get offline messages error %s ", err.Error())
			return
		}

		for _, msg := range messages {
			client.SendMessage(websocket.BinaryMessage, []byte(msg))
			time.Sleep(time.Second / 50)
		}

		_, err = svcCtx.RedisClient.LTrim(ctx, u.Uuid, 100, -1).Result()
		if err != nil {
			logx.Error(err)
		}
	}()

}

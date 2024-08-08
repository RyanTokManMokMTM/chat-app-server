package ws

import (
	"context"
	"github.com/gorilla/websocket"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/server/socketClient"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/server/socketServer"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/serverTypes"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/variable"
	"net/http"
	"time"

	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ConnectWsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewConnectWsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ConnectWsLogic {
	return &ConnectWsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ConnectWsLogic) ConnectWs(w http.ResponseWriter, r *http.Request, wsServer serverTypes.ISocketServer) {
	// todo: add your logic here and delete this line
	conn, err := socketServer.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Websocket upgrade error"))
		return
	}
	//TODO : Get UserID from Context
	userID := ctxtool.GetUserIDFromCTX(r.Context())
	//TODO : Find User Info from DB
	u, err := l.svcCtx.DAO.FindOneUser(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errx.NewCustomErrCode(errx.USER_NOT_EXIST).GetMessage()))
		return
	}

	websocketServer := wsServer.(*socketServer.SocketServer)
	client := socketClient.NewSocketClient(u.Uuid, u.NickName, conn, websocketServer, l.svcCtx)
	wsServer.RegisterClient(client)

	go client.ReadLoop()
	go client.WriteLoop()
	//
	go func() {
		ctx := context.Background()

		//we need to create a connection for each user?
		len, err := variable.RedisConnection.LLen(ctx, u.Uuid).Result()
		if err != nil {
			logx.Error("getting Redis length err ", err)
			return
		}

		messages, err := l.svcCtx.RedisClient.LRange(ctx, u.Uuid, 0, len).Result()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			logx.Errorf("get offline messages error %s ", err.Error())
			return
		}

		for _, msg := range messages {
			client.SendMessage(websocket.BinaryMessage, []byte(msg))
			time.Sleep(time.Second / 50)
		}

		_, err = l.svcCtx.RedisClient.LTrim(ctx, u.Uuid, 100, -1).Result()
		if err != nil {
			logx.Error(err)
		}
	}()
}

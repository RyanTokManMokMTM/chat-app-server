package router

import (
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler/ws"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/redisClient"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ConfigRouter(c config.Config) *rest.Server {
	server := rest.MustNewServer(c.RestConf) // new a router
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		var e *errx.CustomError
		switch {

		case errors.As(err, &e):
			return http.StatusOK, e.ToJSON()
		default:
			return http.StatusInternalServerError, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error()).ToJSON()
		}
	})
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/ws",
		Handler: ws.WebSocketHandler(ctx),
	}, rest.WithJwt(c.Auth.AccessSecret))

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/resources/:file",
		Handler: http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))).ServeHTTP,
	})

	if c.Mode != service.TestMode {
		client, err := redisClient.ConnectToClient(c.Redis.Addr, c.Redis.Password) //connect to redis for using in websocket
		if err != nil {
			panic("failed to connect to redis")
		}
		variable.RedisConnection = client
	}

	return server
}

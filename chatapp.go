package main

import (
	"flag"
	"fmt"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler/ws"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/redisClient"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chatapp.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	//fmt.Println(c.MaxFileSize)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {

		case *errx.CustomError:
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

	client, err := redisClient.ConnectToClient(c.Redis.Addr, c.Redis.Password)
	if err != nil {
		panic("failed to connect to redis")
	}

	variable.RedisConnection = client

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}

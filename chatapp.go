package main

import (
	"flag"
	"fmt"
	"github.com/ryantokmanmok/chat-app-server/internal/handler/ws"
	"net/http"

	"github.com/ryantokmanmok/chat-app-server/internal/config"
	"github.com/ryantokmanmok/chat-app-server/internal/handler"
	"github.com/ryantokmanmok/chat-app-server/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/chatapp.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/ws",
		Handler: ws.WebSocketHandler(ctx),
	}, rest.WithJwt(c.Auth.AccessSecret))

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()

}

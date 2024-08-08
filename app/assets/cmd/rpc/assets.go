package main

import (
	"flag"
	"fmt"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/server"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/assets.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		assets_api.RegisterAssetRPCServer(grpcServer, server.NewAssetRPCServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

package main

import (
	"flag"
	"fmt"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/rpc"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/config"
	friendserviceServer "github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/server/friendservice"
	groupserviceServer "github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/server/groupservice"
	storyserviceServer "github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/server/storyservice"
	userserviceServer "github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/server/userservice"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/core.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		core.RegisterUserServiceServer(grpcServer, userserviceServer.NewUserServiceServer(ctx))
		core.RegisterStoryServiceServer(grpcServer, storyserviceServer.NewStoryServiceServer(ctx))
		core.RegisterGroupServiceServer(grpcServer, groupserviceServer.NewGroupServiceServer(ctx))
		core.RegisterFriendServiceServer(grpcServer, friendserviceServer.NewFriendServiceServer(ctx))
		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	s.AddUnaryInterceptors(rpc.LoggerInterceptor)
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

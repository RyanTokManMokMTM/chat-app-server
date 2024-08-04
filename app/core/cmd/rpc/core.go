package main

import (
	"api/app/common/rpc"
	"flag"
	"fmt"

	"api/app/core/cmd/rpc/internal/config"
	friendserviceServer "api/app/core/cmd/rpc/internal/server/friendservice"
	groupserviceServer "api/app/core/cmd/rpc/internal/server/groupservice"
	storyserviceServer "api/app/core/cmd/rpc/internal/server/storyservice"
	userserviceServer "api/app/core/cmd/rpc/internal/server/userservice"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

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

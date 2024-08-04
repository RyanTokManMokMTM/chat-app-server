package svc

import (
	"api/app/core/cmd/api/internal/config"
	"api/app/core/cmd/rpc/client/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:      c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.CoreRPC)),
	}
}

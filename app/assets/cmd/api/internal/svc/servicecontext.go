package svc

import (
	"api/app/assets/cmd/api/internal/config"
	"api/app/assets/cmd/rpc/assetrpc"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	AssetRPC assetrpc.AssetRPC
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		AssetRPC: assetrpc.NewAssetRPC(zrpc.MustNewClient(c.AssetsRPC)),
	}
}

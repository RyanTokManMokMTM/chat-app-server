package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/assetrpc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/redisx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/dao"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/models"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	DAO       dao.Store
	AssetsRPC assetrpc.AssetRPC
	RedisCli  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := models.NewEngine(c.MySQL.DataSource, c.MySQL.MaxIdleConns, c.MySQL.MaxOpenConns)
	newDAO := dao.NewDao(db)
	client, _ := redisx.ConnectToClient(c.RedisPubSub.Addr, c.RedisPubSub.Password)
	return &ServiceContext{
		Config:    c,
		AssetsRPC: assetrpc.NewAssetRPC(zrpc.MustNewClient(c.AssetsRPC)),
		DAO:       newDAO,
		RedisCli:  client,
	}
}

package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/redisClient"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/dao"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/models"
)

type ServiceContext struct {
	Config      config.Config
	DAO         dao.Store
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := models.NewEngine(c.MySQL.DataSource, c.MySQL.MaxIdleConns, c.MySQL.MaxOpenConns)
	newDAO := dao.NewDao(db)
	client, _ := redisClient.ConnectToClient(c.Redis.Addr, c.Redis.Password)
	return &ServiceContext{
		Config:      c,
		DAO:         newDAO,
		RedisClient: client,
	}
}

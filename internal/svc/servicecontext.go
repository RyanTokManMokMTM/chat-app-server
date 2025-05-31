package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/redisClient"
	unitofwork "github.com/ryantokmanmokmtm/chat-app-server/internal/uow"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Client
	Uow         *unitofwork.UOW
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, _ := redisClient.ConnectToClient(c.Redis.Addr, c.Redis.Password)
	engine := models.NewEngine(&c)
	return &ServiceContext{
		Config:      c,
		RedisClient: client,
		Uow:         unitofwork.NewUOW(engine),
	}
}

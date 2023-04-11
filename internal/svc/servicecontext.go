package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/ryantokmanmok/chat-app-server/internal/config"
	"github.com/ryantokmanmok/chat-app-server/internal/dao"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
	"github.com/ryantokmanmok/chat-app-server/internal/redisClient"
)

type ServiceContext struct {
	Config      config.Config
	DAO         dao.DAOInterface
	RedisClient *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, _ := redisClient.ConnectToClient(c.Redis.Addr, c.Redis.Password)
	return &ServiceContext{
		Config:      c,
		DAO:         dao.NewDao(models.NewEngine(&c)),
		RedisClient: client,
	}
}

package svc

import (
	"github.com/ryantokmanmok/chat-app-server/internal/config"
	"github.com/ryantokmanmok/chat-app-server/internal/dao"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

type ServiceContext struct {
	Config config.Config
	DAO    dao.DAOInterface
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DAO:    dao.NewDao(models.NewEngine(&c)),
	}
}

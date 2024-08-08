package svc

import (
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/dao"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/models"
)

type ServiceContext struct {
	Config config.Config
	DAO    dao.Store
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := models.NewEngine(c.MySQL.DataSource, c.MySQL.MaxIdleConns, c.MySQL.MaxOpenConns)
	newDAO := dao.NewDao(db)
	return &ServiceContext{
		Config: c,
		DAO:    newDAO,
	}
}

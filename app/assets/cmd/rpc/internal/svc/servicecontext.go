package svc

import (
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/dao"
)

type ServiceContext struct {
	Config config.Config
	DAO    dao.Store
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}

package svc

import (
	"api/app/assets/cmd/rpc/internal/config"
	"api/app/internal/dao"
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

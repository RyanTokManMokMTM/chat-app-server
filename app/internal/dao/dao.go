package dao

import (
	"gorm.io/gorm"
)

type DAO struct {
	engine *gorm.DB
}

// CHECK DAO IMPLEMENTED DAO_INTERFACE
var _ Store = (*DAO)(nil)

func NewDao(engine *gorm.DB) Store {
	return &DAO{
		engine: engine,
	}
}

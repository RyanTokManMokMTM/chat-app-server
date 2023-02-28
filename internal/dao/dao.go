package dao

import (
	"gorm.io/gorm"
)

type DAO struct {
	engine *gorm.DB
}

// CHECK DAO IMPLEMENTED DAO_INTERFACE
var _ DAOInterface = (*DAO)(nil)

func NewDao(engine *gorm.DB) DAOInterface {
	return &DAO{
		engine: engine,
	}
}

package dao

import (
	"gorm.io/gorm"
)

type DAO struct {
	engine *gorm.DB
}

func NewDao(engine *gorm.DB) DAOInterface {
	return &DAO{
		engine: engine,
	}
}

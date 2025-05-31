package db

import (
	"gorm.io/gorm"
)

type DatabaseEngine[T any] interface {
	WithEngine() *T
}

var _ DatabaseEngine[gorm.DB] = (*GormEngine)(nil)

type GormEngine struct {
	engine *gorm.DB
}

func NewGormEngine(engine *gorm.DB) DatabaseEngine[gorm.DB] {
	return &GormEngine{
		engine: engine,
	}
}

func (ge *GormEngine) WithEngine() *gorm.DB {
	return ge.engine
}

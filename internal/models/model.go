package models

import (
	"github.com/ryantokmanmok/chat-app-server/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type CommonField struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewEngine(c *config.Config) *gorm.DB {

	sql, err := gorm.Open(mysql.Open(c.MySQL.DataSource), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	db, err := sql.DB()

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	db.SetMaxOpenConns(c.MySQL.MaxOpenConns)
	migration(sql)
	return sql
}

func migration(db *gorm.DB) {
	err := db.AutoMigrate(UserModel{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(UserFriend{})
	if err != nil {
		panic(err)
	}
}

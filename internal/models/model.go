package models

import (
	"github.com/ryantokmanmokmtm/chat-app-server/internal/config"
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

	err := db.AutoMigrate(&User{}, &Group{}, &UserGroup{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&UserGroup{})
	if err != nil {
		panic(err)
	}

	err = db.SetupJoinTable(&User{}, "Groups", &UserGroup{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(UserFriend{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(Message{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(StoryModel{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(UserStorySeen{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(UserStoryLikes{})
	if err != nil {
		panic(err)
	}

}

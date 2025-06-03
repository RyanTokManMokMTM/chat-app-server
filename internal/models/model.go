package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Model interface {
	TableName() string
}

type Base struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;not null"`
	UUID      string `gorm:"primaryKey;autoIncrement;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) error {
	b.UUID = uuid.New().String()
	return nil
}

func (b *Base) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
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

	err := db.AutoMigrate(&User{}, &Group{}, &UserGroup{}, Sticker{})
	if err != nil {
		panic(err)
	}
	//
	//err = db.AutoMigrate(&UserGroup{})
	//if err != nil {
	//	panic(err)
	//}

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

	err = db.AutoMigrate(&StickerResource{})
	if err != nil {
		panic(err)
	}

}

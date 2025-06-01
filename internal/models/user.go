package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateUserDTO struct {
	NickName      string
	Email         string
	Password      string
	Avatar        string
	Cover         string
	StatusMessage string
}

type UpdateUserDTO struct {
	NickName      string
	Avatar        string
	Cover         string
	StatusMessage string
}

type User struct {
	Base
	Id            uint   `gorm:"primaryKey;autoIncrement;not null"`
	Uuid          string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	NickName      string `gorm:"types:varchar(32)"`
	Email         string `gorm:"types:varchar(64)"`
	Password      string `gorm:"types:varchar(64)"`
	Avatar        string `gorm:"types:varchar(64);null;comment:'user avatar'"`
	Cover         string `gorm:"types:varchar(64);null;comment:'user cover'"`
	StatusMessage string `gorm:"types:varchar(64);null;comment:'user status message'"`

	Stories       []StoryModel `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Groups        []Group      `gorm:"many2many:users_groups;foreignKey:Id;joinForeignKey:UserId"`
	StickerGroups []Sticker    `gorm:"many2many:users_stickers;foreignKey:Id;joinForeignKey:UserId;References:Uuid;joinReferences:StickerId"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Uuid = uuid.New().String()
	u.Avatar = "/default.jpg"
	u.Cover = "/cover.jpg"
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (u *User) TableName() string {
	return "users_info"
}

package models

import (
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

const userTableName = "users"

type User struct {
	Base
	NickName      string `gorm:"types:varchar(32)"`
	Email         string `gorm:"types:varchar(64)"`
	Password      string `gorm:"types:varchar(64)"`
	Avatar        string `gorm:"types:varchar(64);null;comment:'user avatar'"`
	Cover         string `gorm:"types:varchar(64);null;comment:'user cover'"`
	StatusMessage string `gorm:"types:varchar(64);null;comment:'user status message'"`

	Stories       []StoryModel `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Groups        []Group      `gorm:"many2many:users_groups;foreignKey:ID;joinForeignKey:UserID"`
	StickerGroups []Sticker    `gorm:"many2many:users_stickers;foreignKey:ID;joinForeignKey:UserID;References:UUID;joinReferences:StickerID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Avatar = "/default.jpg"
	u.Cover = "/cover.jpg"
	return nil
}

func (u *User) TableName() string {
	return userTableName
}

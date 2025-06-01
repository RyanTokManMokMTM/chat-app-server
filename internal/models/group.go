package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateGroupDTO struct {
	GroupName   string
	GroupAvatar string
	GroupDesc   string
	GroupLead   uint
}

type UpdateGroupDTO struct {
	GroupName   string
	GroupAvatar string
	GroupDesc   string
}

type Group struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Uuid        string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	GroupName   string `gorm:"types:varchar(64);not null"`
	GroupAvatar string
	GroupDesc   string
	GroupLead   uint `gorm:"not null;index"`

	LeadInfo  User   `gorm:"foreignKey:GroupLead;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UsersInfo []User `gorm:"many2many:users_groups;foreignKey:Id;joinForeignKey:GroupId"`
	Base
}

func (g *Group) BeforeCreate(tx *gorm.DB) error {
	g.Uuid = uuid.New().String()
	//g.GroupAvatar = "/defaultGroup.jpg"
	return nil
}

func (g *Group) TableName() string {
	return "groups_info"
}

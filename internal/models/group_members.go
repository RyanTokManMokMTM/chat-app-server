package models

import (
	"context"
	"gorm.io/gorm"
)

type GroupMember struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	GroupID uint `gorm:"not null;index;comment:'belong to which group ID'"`
	UserID  uint `gorm:"not null;index;comment:'who belong to this group'"`

	MemberInfo UserModel `gorm:"foreignKey:UserID"`
	GroupInfo  Group     `gorm:"foreignKey:GroupID"`
	CommonField
}

func (gm *GroupMember) TableName() string {
	return "groups_members"
}

func (gm *GroupMember) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&gm).Error
}
func (gm *GroupMember) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().First(&gm).Error
}
func (gm *GroupMember) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Delete(&gm).Error
}

func (gm *GroupMember) GetGroupMemberList(ctx context.Context, db *gorm.DB) ([]*GroupMember, error) {
	var members []*GroupMember
	if err := db.WithContext(ctx).Debug().Where("group_id = ?", gm.GroupID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

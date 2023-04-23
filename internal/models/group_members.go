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
	return db.WithContext(ctx).Debug().Where("group_id = ? AND user_id = ?", gm.GroupID, gm.UserID).First(&gm).Error
}

//
//func (gm *GroupMember) FindAll(ctx context.Context, db *gorm.DB) ([]*GroupMember, error) {
//	var members []*GroupMember
//	if err := db.WithContext(ctx).Debug().Where("group_id = ?", gm.GroupID).Preload("MemberInfo").Find(&members).Error; err != nil {
//		return nil, err
//	}
//	return members, nil
//}

func (gm *GroupMember) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("group_id = ? AND user_id = ?", gm.GroupID, gm.UserID).Delete(&gm).Error
}

func (gm *GroupMember) DeleteAll(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("group_id = ?", gm.GroupID).Delete(&gm).Error
}

func (gm *GroupMember) GetGroupMemberList(ctx context.Context, db *gorm.DB) ([]*GroupMember, error) {
	var members []*GroupMember
	if err := db.WithContext(ctx).Debug().Where("group_id = ?", gm.GroupID).Preload("MemberInfo").Preload("GroupInfo").Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (gm *GroupMember) FindUserGroup(ctx context.Context, db *gorm.DB) ([]*GroupMember, error) {
	var groups []*GroupMember
	if err := db.WithContext(ctx).Debug().Preload("GroupInfo").Where("user_id = ?", gm.UserID).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (gm *GroupMember) CountGroupMembers(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.WithContext(ctx).Debug().Model(gm).Where("group_id = ?", gm.GroupID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

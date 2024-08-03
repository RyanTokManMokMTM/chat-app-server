package models

import (
	"context"
	"gorm.io/gorm"
)

type UserGroup struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	GroupId uint `gorm:"not uniqueIndex;index;comment:'belong to which group Id'"`
	UserId  uint `gorm:"not uniqueIndex;index;comment:'who belong to this group'"`

	MemberInfo User  `gorm:"foreignKey:UserId"`
	GroupInfo  Group `gorm:"foreignKey:GroupId"`
	CommonField
}

func (gm *UserGroup) TableName() string {
	return "users_groups"
}

func (gm *UserGroup) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&gm).Error
}
func (gm *UserGroup) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("group_id = ? AND user_id = ?", gm.GroupId, gm.UserId).First(&gm).Error
}

//
//func (gm *UserGroup) FindAll(ctx context.Context, db *gorm.DB) ([]*UserGroup, error) {
//	var members []*UserGroup
//	if err := db.WithContext(ctx).Debug().Where("group_id = ?", gm.GroupId).Preload("MemberInfo").Find(&members).Error; err != nil {
//		return nil, err
//	}
//	return members, nil
//}

func (gm *UserGroup) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("group_id = ? AND user_id = ?", gm.GroupId, gm.UserId).Delete(&gm).Error
}

func (gm *UserGroup) DeleteAll(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("group_id = ?", gm.GroupId).Delete(&gm).Error
}

func (gm *UserGroup) GetGroupMemberList(ctx context.Context, db *gorm.DB) ([]*UserGroup, error) {
	var members []*UserGroup
	if err := db.WithContext(ctx).Debug().
		Where("group_id = ?", gm.GroupId).
		Preload("MemberInfo").Preload("GroupInfo").
		Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (gm *UserGroup) GetGroupMemberListByPage(ctx context.Context, db *gorm.DB, pageOffset, pageLimit int) ([]*UserGroup, error) {
	var members []*UserGroup
	if err := db.WithContext(ctx).Debug().
		Where("group_id = ?", gm.GroupId).
		Preload("MemberInfo").Preload("GroupInfo").
		Offset(pageOffset).
		Limit(pageLimit).
		Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (gm *UserGroup) FindUserGroup(ctx context.Context, db *gorm.DB, pageOffset, pageSize int) ([]*UserGroup, error) {
	var groups []*UserGroup
	if err := db.WithContext(ctx).Debug().Preload("GroupInfo").Where("user_id = ?", gm.UserId).Offset(pageOffset).Limit(pageSize).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (gm *UserGroup) CountGroupMembers(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.WithContext(ctx).Debug().Model(gm).Where("group_id = ?", gm.GroupId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

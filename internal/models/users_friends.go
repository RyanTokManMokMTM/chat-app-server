package models

import (
	"context"
	"gorm.io/gorm"
)

type UserFriend struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"not null;index"`
	Friend uint `gorm:"not null;index"`
	CommonField
}

func (uf *UserFriend) TableName() string {
	return "users_friends"
}

func (uf *UserFriend) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&uf).Error
}
func (uf *UserFriend) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().First(&uf).Error
}
func (uf *UserFriend) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("user_id = ? AND friend = ?", uf.UserID, uf.Friend).Delete(&uf).Error
}
func (uf *UserFriend) GetFriendList(ctx context.Context, db *gorm.DB) ([]*UserFriend, error) {
	var list []*UserFriend
	if err := db.WithContext(ctx).Debug().Where("user_id = ?", uf.UserID).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

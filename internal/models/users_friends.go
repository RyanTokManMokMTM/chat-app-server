package models

import (
	"context"
	"gorm.io/gorm"
)

type UserFriend struct {
	//user A added user B ,but it doesn't mean user B has added userA ?
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserID   uint `gorm:"not null;index"`
	FriendID uint `gorm:"not null;index"`

	//UserInfo   UserModel `gorm:"foreignKey:UserID"`
	FriendInfo UserModel `gorm:"foreignKey:FriendID"`
	CommonField
}

func (uf *UserFriend) TableName() string {
	return "users_friends"
}

func (uf *UserFriend) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&uf).Error
}
func (uf *UserFriend) FindOne(ctx context.Context, db *gorm.DB, userID, friendID uint) error {
	return db.WithContext(ctx).Debug().Where("user_id = ? AND friend_id = ? ", userID, friendID).First(&uf).Error
}
func (uf *UserFriend) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("user_id = ? AND friend_id = ?", uf.UserID, uf.FriendID).Delete(&uf).Error
}
func (uf *UserFriend) GetFriendList(ctx context.Context, db *gorm.DB) ([]*UserFriend, error) {
	var list []*UserFriend
	
	if err := db.WithContext(ctx).Debug().Model(&uf).Where("user_id = ?", uf.UserID).Preload("FriendInfo").Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

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

	//UserInfo   User `gorm:"foreignKey:UserId"`
	FriendInfo User `gorm:"foreignKey:FriendID"`

	//All stories
	CommonField
}

func (uf *UserFriend) TableName() string {
	return "users_friends"
}

func (uf *UserFriend) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&uf).Error
}
func (uf *UserFriend) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&uf).Where("user_id = ? AND friend_id = ? ", uf.UserID, uf.FriendID).First(&uf).Error
}
func (uf *UserFriend) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("user_id = ? AND friend_id = ?", uf.UserID, uf.FriendID).Delete(&uf).Error
}

func (uf *UserFriend) GetFriendList(ctx context.Context, db *gorm.DB, pageOffset, pageSize int) ([]*UserFriend, error) {
	var list []*UserFriend
	if err := db.WithContext(ctx).Debug().Model(&uf).
		Preload("FriendInfo").
		Where("user_id = ?", uf.UserID, uf.ID).
		Offset(pageOffset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (uf *UserFriend) CountUserFriends(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64 = 0
	if err := db.Debug().WithContext(ctx).Model(uf).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

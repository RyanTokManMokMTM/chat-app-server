package models

import (
	"context"
	"gorm.io/gorm"
)

type UserStorySeen struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserId   uint `gorm:"comment:'belong to which user Id'"`
	FriendId uint `gorm:"comment:'belong to which friend Id'"`
	StoryId  uint `gorm:"comment:'belong to which story Id'"`

	StoryInfo StoryModel `gorm:"foreignKey:StoryId"`
	UserInfo  User       `gorm:"foreignKey:UserId"`
	CommonField
}

func (uss *UserStorySeen) TableName() string {
	return "user_story_seen"
}

func (uss *UserStorySeen) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(uss).Error
}
func (uss *UserStorySeen) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Preload("StoryInfo").Where("user_id = ? AND friend_id = ? AND story_id = ?", uss.UserId, uss.FriendId, uss.StoryId).First(uss).Error
}

func (uss *UserStorySeen) FindLatestOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Preload("StoryInfo").Where("user_id = ? AND friend_id = ?", uss.UserId, uss.FriendId).Last(uss).Error
}

func (uss *UserStorySeen) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Delete(uss).Error
}
func (uss *UserStorySeen) UpdateOne(ctx context.Context, db *gorm.DB, storyId uint) error {
	return db.WithContext(ctx).Debug().Model(&uss).Update("StoryId", storyId).Error
}

func (uss *UserStorySeen) GetStoryLikedUserSeen(ctx context.Context, db *gorm.DB, limit int) ([]*UserStorySeen, error) {
	var userList []*UserStorySeen
	if err := db.WithContext(ctx).Debug().Model(uss).Preload("UserInfo").Where("story_id= ?", uss.StoryId).Order("created_at DESC").Limit(limit).Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func (uss *UserStorySeen) CountOneStorySeen(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64
	if err := db.WithContext(ctx).Debug().Model(uss).Where("story_id= ?", uss.StoryId).Count(&count).Error; err != nil {
		return 0, nil
	}
	return count, nil
}

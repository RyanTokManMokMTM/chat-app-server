package models

import (
	"context"
	"gorm.io/gorm"
)

type UserStorySeen struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	UserId       uint `gorm:"comment:'belong to which user Id'"`
	FriendId     uint `gorm:"comment:'belong to which friend Id'"`
	StoryModelID uint `gorm:"comment:'belong to which story Id'"`

	StoryInfo StoryModel `gorm:"foreignKey:StoryModelID"`
	CommonField
}

func (uss *UserStorySeen) TableName() string {
	return "user_story_seen"
}

func (uss *UserStorySeen) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(uss).Error
}
func (uss *UserStorySeen) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Preload("StoryInfo").Where("user_id = ? AND friend_id = ?", uss.UserId, uss.FriendId).First(uss).Error
}
func (uss *UserStorySeen) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Delete(uss).Error
}
func (uss *UserStorySeen) UpdateOne(ctx context.Context, db *gorm.DB, storyId uint) error {
	return db.WithContext(ctx).Debug().Model(&uss).Update("StoryModelID", storyId).Error
}

package models

import (
	"context"
	"gorm.io/gorm"
)

type UserStoryLikes struct {
	ID      uint
	UserId  uint
	StoryId uint

	UserInfo User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommonField
}

func (usl *UserStoryLikes) TableName() string {
	return "user_story_likes"
}

func (usl *UserStoryLikes) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(usl).Error
}

func (usl *UserStoryLikes) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("user_id = ? AND story_id = ?", usl.UserId, usl.StoryId).First(usl).Error
}

func (usl *UserStoryLikes) UpdateOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&usl).Updates(usl).Error
}

func (usl *UserStoryLikes) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&usl).Delete(usl).Error
}

func (usl *UserStoryLikes) CountStoryLikes(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64 = 0
	if err := db.WithContext(ctx).Debug().Model(usl).Where("story_id = ?", usl.StoryId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

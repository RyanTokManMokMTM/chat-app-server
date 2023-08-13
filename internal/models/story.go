package models

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type StoryModel struct {
	Id             uint `gorm:"primaryKey;autoIncrement;type:int"`
	UserId         uint `gorm:"not null;index;comment:'belong to which group Id'"`
	StoryMediaPath string

	UserInfo User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommonField
}

type (
	StoriesWithIds struct {
		StoryModel
		Ids string
	}
)

func (s *StoryModel) TableName() string {
	return "stories"
}

func (s *StoryModel) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&s).Error
}

func (s *StoryModel) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("id = ?", s.Id).First(s).Error
}

func (s *StoryModel) FindOneUserStory(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("id = ? AND user_id = ?", s.Id, s.UserId).First(s).Error
}

func (s *StoryModel) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("id = ?", s.Id).Delete(s).Error
}

func (s *StoryModel) GetFriendStoryList(ctx context.Context, db *gorm.DB, userId uint, pageOffset, pageLimit int) ([]*StoriesWithIds, error) {
	var stories []*StoriesWithIds
	if err := db.WithContext(ctx).Debug().Model(s).Select("*", "group_concat(id) as ids").Preload("UserInfo").Where("user_id IN (?)",
		db.Model(UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY").Group("user_id").
		Offset(pageOffset).Limit(pageLimit).Find(&stories).Error; err != nil {
		return nil, err
	}

	return stories, nil
}

func (s *StoryModel) CountFriendActiveStory(ctx context.Context, db *gorm.DB, userId uint) (int64, error) {
	var count int64
	if err := db.WithContext(ctx).Debug().Model(s).Where("user_id IN (?)",
		db.Model(UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY").Group("user_id").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *StoryModel) FindAllUserStories(ctx context.Context, db *gorm.DB) ([]uint, error) {
	var ids []uint
	now := time.Now().Unix()
	logx.Info(now)
	beforeOneDay := now - (86400) //24 hour
	err := db.WithContext(ctx).Debug().Model(&s).Select("Id").Where("user_id = ? AND created_at BETWEEN FROM_UNIXTIME(?) AND FROM_UNIXTIME(?)", s.UserId, beforeOneDay, now).Find(&ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

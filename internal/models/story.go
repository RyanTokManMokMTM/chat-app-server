package models

import (
	"context"
	"gorm.io/gorm"
)

type StoryModel struct {
	Id             uint `gorm:"primaryKey;autoIncrement;type:int"`
	UserId         uint `gorm:"not null;index;comment:'belong to which group Id'"`
	StoryMediaPath string

	UserInfo User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommonField
}

//type (
//	StoriesWithTime struct {
//		StoryModel
//		Ids string
//	}
//)

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

func (s *StoryModel) GetActiveStoryList(ctx context.Context, db *gorm.DB, userId uint, pageOffset, pageLimit int) ([]*StoryModel, error) {
	var stories []*StoryModel
	if err := db.WithContext(ctx).Debug().Model(s).Select("*", "max(created_at) as t").Preload("UserInfo").Where("user_id IN (?)",
		db.Model(UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY").Group("user_id").
		Offset(pageOffset).
		Limit(pageLimit).Order("t desc").
		Find(&stories).Error; err != nil {
		return nil, err
	}

	return stories, nil
}

func (s *StoryModel) GetActiveStoryListByTime(ctx context.Context, db *gorm.DB, userId uint, pageOffset, pageLimit int, timeStamp int64) ([]*StoryModel, error) {
	var stories []*StoryModel
	if err := db.WithContext(ctx).Debug().Model(s).Select("*", "max(created_at) as t").Preload("UserInfo").Where("user_id IN (?)",
		db.Model(UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY AND created_at  <= FROM_UNIXTIME(?)", timeStamp).Group("user_id").
		Offset(pageOffset).
		Limit(pageLimit).Order("t desc").
		Find(&stories).Error; err != nil {
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

func (s *StoryModel) CountFriendActiveStoryByTime(ctx context.Context, db *gorm.DB, userId uint, timeStamp int64) (int64, error) {
	var count int64
	if err := db.WithContext(ctx).Debug().Model(s).Where("user_id IN (?)",
		db.Model(UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY  AND created_at <= FROM_UNIXTIME(?)", timeStamp).Group("user_id").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (s *StoryModel) FindAllUserStories(ctx context.Context, db *gorm.DB) ([]uint, error) {
	var ids []uint
	err := db.WithContext(ctx).Debug().
		Model(&s).
		Select("Id").
		Where("user_id = ? AND created_at >= NOW() - INTERVAL 1 DAY", s.UserId).Find(&ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (s *StoryModel) FindAllUserStoriesByTimeStamp(ctx context.Context, db *gorm.DB, timeStamp int64) ([]uint, error) {
	var ids []uint
	err := db.WithContext(ctx).Debug().
		Model(&s).
		Select("Id").
		Where("user_id = ? AND created_at >= NOW() - INTERVAL 1 DAY AND  created_at <= FROM_UNIXTIME(?)", s.UserId, timeStamp).Find(&ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

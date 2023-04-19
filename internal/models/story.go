package models

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type StoryModel struct {
	Id             uint `gorm:"primaryKey;autoIncrement;type:int"`
	UserId         uint `gorm:"not null;index;comment:'belong to which group ID'"`
	StoryMediaPath string

	UserInfo UserModel `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommonField
}

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

package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateStoryDTO struct {
	UserId         uint
	StoryMediaPath string
}

type UpdateStoryDTO struct {
	StoryMediaPath string
}

type StoryModel struct {
	Base
	Id             uint   `gorm:"primaryKey;autoIncrement;types:int"`
	Uuid           string `gorm:"index"`
	UserId         uint   `gorm:"not null;index;comment:'belong to which group Id'"`
	StoryMediaPath string

	UserInfo User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s *StoryModel) BeforeCreate(tx *gorm.DB) error {
	s.Uuid = uuid.New().String()
	return nil
}

type (
	StoriesWithLatestStoryTime struct {
		StoryModel
		LatestTime time.Time
	}
)

func (s *StoryModel) TableName() string {
	return "stories"
}

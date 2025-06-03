package models

import (
	"time"
)

type CreateStoryDTO struct {
	UserID         uint
	StoryMediaPath string
}

type UpdateStoryDTO struct {
	StoryMediaPath string
}

type (
	StoriesWithLatestStoryTime struct {
		StoryModel
		LatestTime time.Time
	}
)

const storyTableName = "stories"

type StoryModel struct {
	Base
	UserID         uint `gorm:"not null;index;comment:'belong to which group ID'"`
	StoryMediaPath string

	UserInfo User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (s *StoryModel) TableName() string {
	return storyTableName
}

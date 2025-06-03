package models

type CreateUserStorySeenDTO struct {
	UserID   uint
	FriendID uint
	StoryID  uint
}

type UpdateUserStorySeenDTO struct {
	StoryID uint
}

const userStorySeenTableName = "user_story_seen"

type UserStorySeen struct {
	Base
	UserID   uint `gorm:"comment:'belong to which user ID'"`
	FriendID uint `gorm:"comment:'belong to which friend ID'"`
	StoryID  uint `gorm:"comment:'belong to which story ID'"`

	StoryInfo StoryModel `gorm:"foreignKey:StoryID"`
	UserInfo  User       `gorm:"foreignKey:UserID"`
}

func (uss *UserStorySeen) TableName() string {
	return userStorySeenTableName
}

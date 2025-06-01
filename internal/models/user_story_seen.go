package models

type CreateUserStorySeenDTO struct {
	UserId   uint
	FriendId uint
	StoryId  uint
}

type UpdateUserStorySeenDTO struct {
	StoryId uint
}

type UserStorySeen struct {
	Base
	ID       uint `gorm:"primaryKey;autoIncrement"`
	UserId   uint `gorm:"comment:'belong to which user Id'"`
	FriendId uint `gorm:"comment:'belong to which friend Id'"`
	StoryId  uint `gorm:"comment:'belong to which story Id'"`

	StoryInfo StoryModel `gorm:"foreignKey:StoryId"`
	UserInfo  User       `gorm:"foreignKey:UserId"`
}

func (uss *UserStorySeen) TableName() string {
	return "user_story_seen"
}

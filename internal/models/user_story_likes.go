package models

type CreateUserStoryLikesDTO struct {
	UserId  uint
	StoryId uint
}

type UpdateUserStoryLikesDTO struct {
	UserId  uint
	StoryId uint
}

type UserStoryLikes struct {
	Base
	ID      uint
	UserId  uint
	StoryId uint

	UserInfo User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (usl *UserStoryLikes) TableName() string {
	return "user_story_likes"
}

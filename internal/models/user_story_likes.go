package models

type CreateUserStoryLikesDTO struct {
	UserID  uint
	StoryID uint
}

type UpdateUserStoryLikesDTO struct {
	UserID  uint
	StoryID uint
}

const userStoryLikeTableName = "user_story_likes"

type UserStoryLikes struct {
	Base
	UserID  uint
	StoryID uint

	UserInfo User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (usl *UserStoryLikes) TableName() string {
	return userStoryLikeTableName
}

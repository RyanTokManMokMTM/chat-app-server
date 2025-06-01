package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserStorySeenRepo[T any] interface {
	CreateOne(ctx context.Context, storySeen T) (*T, error)
	FindOneByID(ctx context.Context, id uint) (*T, error)
	FindOneByUserIdAndFriendIdAndStroyId(ctx context.Context, UserId, friendId, storyId uint) (*models.UserStorySeen, error)
	UpdateOne(ctx context.Context, storySeen T) error
	DeleteOne(ctx context.Context, id uint) error

	FindLatestOne(ctx context.Context, UserId uint, friendId uint) (*models.UserStorySeen, error)
	GetStoryLikedUserSeen(ctx context.Context, storyId uint, limit int) ([]*models.UserStorySeen, error)
	CountOneStorySeen(ctx context.Context, storyId uint) (int64, error)
}

var _ IUserStorySeenRepo[models.UserStorySeen] = (*UserStorySeenRepo)(nil)

type UserStorySeenRepo struct {
	engine *gorm.DB
	IRepository[*models.UserStorySeen, models.UserStorySeen]
}

func NewUserStorySeenRepo(engine *gorm.DB) *UserStorySeenRepo {
	return &UserStorySeenRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserStorySeen, models.UserStorySeen](engine),
	}
}

func (userStorySeenRepo *UserStorySeenRepo) CreateOne(ctx context.Context, storySeen models.UserStorySeen) (*models.UserStorySeen, error) {
	if err := userStorySeenRepo.Create(ctx, &storySeen); err != nil {
		return nil, err
	}
	return &storySeen, nil
}

func (userStorySeenRepo *UserStorySeenRepo) FindOneByID(ctx context.Context, id uint) (*models.UserStorySeen, error) {
	result, err := userStorySeenRepo.Find(ctx, &models.UserStorySeen{ID: id})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (userStorySeenRepo *UserStorySeenRepo) UpdateOne(ctx context.Context, storySeen models.UserStorySeen) error {
	return userStorySeenRepo.Update(ctx, &storySeen)
}

func (userStorySeenRepo *UserStorySeenRepo) DeleteOne(ctx context.Context, id uint) error {
	return userStorySeenRepo.Delete(ctx, &models.UserStorySeen{ID: id})
}

func (userStorySeenRepo *UserStorySeenRepo) FindLatestOne(
	ctx context.Context, UserId uint, friendId uint) (*models.UserStorySeen, error) {
	var result models.UserStorySeen
	if err := userStorySeenRepo.engine.
		WithContext(ctx).
		Debug().
		Preload("StoryInfo").
		Where("user_id = ? AND friend_id = ?", UserId, friendId).Last(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (userStorySeenRepo *UserStorySeenRepo) GetStoryLikedUserSeen(
	ctx context.Context,
	storyId uint,
	limit int) ([]*models.UserStorySeen, error) {
	var userList []*models.UserStorySeen
	if err := userStorySeenRepo.engine.
		WithContext(ctx).
		Debug().
		Model(models.UserStorySeen{}).
		Preload("UserInfo").
		Where("story_id= ?", storyId).Order("created_at DESC").Limit(limit).Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func (userStorySeenRepo *UserStorySeenRepo) CountOneStorySeen(ctx context.Context, storyId uint) (int64, error) {
	var count int64
	if err := userStorySeenRepo.engine.WithContext(ctx).Debug().Model(models.UserStorySeen{}).Where("story_id= ?", storyId).Count(&count).Error; err != nil {
		return 0, nil
	}
	return count, nil
}

func (userStorySeenRepo *UserStorySeenRepo) FindOneByUserIdAndFriendIdAndStroyId(ctx context.Context, UserId, friendId, storyId uint) (*models.UserStorySeen, error) {
	var result models.UserStorySeen
	if err := userStorySeenRepo.engine.
		WithContext(ctx).
		Debug().
		Preload("StoryInfo").
		Where("user_id = ? AND friend_id = ? AND story_id = ?", UserId, friendId, storyId).
		First(&result).Error; err != nil {
		return nil, err
	}

	return &result, nil
}

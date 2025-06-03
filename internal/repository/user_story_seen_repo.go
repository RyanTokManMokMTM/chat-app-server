package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserStorySeenRepo[T any] interface {
	CreateOne(ctx context.Context, storySeen T) (T, error)
	FindOneByID(ctx context.Context, ID uint) (T, error)
	FindOneByUserIDAndFriendIDAndStroyID(ctx context.Context, UserID, friendID, storyID uint) (T, error)
	UpdateOne(ctx context.Context, storySeen T) error
	DeleteOne(ctx context.Context, ID uint) error

	FindLatestOne(ctx context.Context, UserID uint, friendID uint) (T, error)
	GetStoryLikedUserSeen(ctx context.Context, storyID uint, limit int) ([]T, error)
	CountOneStorySeen(ctx context.Context, storyID uint) (int64, error)
}

var _ IUserStorySeenRepo[*models.UserStorySeen] = (*userStorySeenRepo)(nil)

type userStorySeenRepo struct {
	engine *gorm.DB
	IRepository[*models.UserStorySeen, models.UserStorySeen]
}

func NewUserStorySeenRepo(engine *gorm.DB) IUserStorySeenRepo[*models.UserStorySeen] {
	return &userStorySeenRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserStorySeen, models.UserStorySeen](engine),
	}
}

func (userStorySeenRepo *userStorySeenRepo) CreateOne(ctx context.Context, storySeen *models.UserStorySeen) (*models.UserStorySeen, error) {
	return userStorySeenRepo.Create(ctx, storySeen)
}

func (userStorySeenRepo *userStorySeenRepo) FindOneByID(ctx context.Context, ID uint) (*models.UserStorySeen, error) {
	return userStorySeenRepo.Find(ctx, &models.UserStorySeen{Base: models.Base{ID: ID}})
}

func (userStorySeenRepo *userStorySeenRepo) UpdateOne(ctx context.Context, storySeen *models.UserStorySeen) error {
	return userStorySeenRepo.Update(ctx, storySeen)
}

func (userStorySeenRepo *userStorySeenRepo) DeleteOne(ctx context.Context, ID uint) error {
	return userStorySeenRepo.Delete(ctx, &models.UserStorySeen{Base: models.Base{ID: ID}})
}

func (userStorySeenRepo *userStorySeenRepo) FindLatestOne(
	ctx context.Context, UserID uint, friendID uint) (*models.UserStorySeen, error) {
	var result models.UserStorySeen
	if err := userStorySeenRepo.engine.
		WithContext(ctx).
		Debug().
		Preload("StoryInfo").
		Where("user_ID = ? AND friend_ID = ?", UserID, friendID).Last(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (userStorySeenRepo *userStorySeenRepo) GetStoryLikedUserSeen(
	ctx context.Context,
	storyID uint,
	limit int) ([]*models.UserStorySeen, error) {
	var userList []*models.UserStorySeen
	if err := userStorySeenRepo.engine.
		WithContext(ctx).
		Debug().
		Model(models.UserStorySeen{}).
		Preload("UserInfo").
		Where("story_ID= ?", storyID).Order("created_at DESC").Limit(limit).Find(&userList).Error; err != nil {
		return nil, err
	}
	return userList, nil
}

func (userStorySeenRepo *userStorySeenRepo) CountOneStorySeen(ctx context.Context, storyID uint) (int64, error) {
	var count int64
	if err := userStorySeenRepo.engine.WithContext(ctx).Debug().Model(models.UserStorySeen{}).Where("story_ID= ?", storyID).Count(&count).Error; err != nil {
		return 0, nil
	}
	return count, nil
}

func (userStorySeenRepo *userStorySeenRepo) FindOneByUserIDAndFriendIDAndStroyID(ctx context.Context, UserID, friendID, storyID uint) (*models.UserStorySeen, error) {
	var result *models.UserStorySeen
	if err := userStorySeenRepo.engine.
		WithContext(ctx).
		Debug().
		Preload("StoryInfo").
		Where("user_ID = ? AND friend_ID = ? AND story_ID = ?", UserID, friendID, storyID).
		First(result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

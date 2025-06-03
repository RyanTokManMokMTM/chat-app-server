package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserStoryLikesRepo[T any] interface {
	CreateOne(ctx context.Context, storyLikes T) (T, error)
	FindOneByID(ctx context.Context, ID uint) (T, error)
	FindOneByUserIDAndStoryID(ctx context.Context, userID, storyID uint) (T, error)
	UpdateOne(ctx context.Context, storyLikes T) error
	DeleteOne(ctx context.Context, ID uint) error

	CountStoryLikes(ctx context.Context, storyID uint) (int64, error)
}

var _ IUserStoryLikesRepo[*models.UserStoryLikes] = (*userStoryLikesRepo)(nil)

type userStoryLikesRepo struct {
	engine *gorm.DB
	IRepository[*models.UserStoryLikes, models.UserStoryLikes]
}

func NewUserStoryLikesRepo(engine *gorm.DB) IUserStoryLikesRepo[*models.UserStoryLikes] {
	return &userStoryLikesRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserStoryLikes, models.UserStoryLikes](engine),
	}
}

func (userRepo *userStoryLikesRepo) CreateOne(ctx context.Context, storyLikes *models.UserStoryLikes) (*models.UserStoryLikes, error) {
	return userRepo.Create(ctx, storyLikes)
}

func (userRepo *userStoryLikesRepo) FindOneByID(ctx context.Context, ID uint) (*models.UserStoryLikes, error) {
	result, err := userRepo.Find(ctx, &models.UserStoryLikes{Base: models.Base{ID: ID}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (userRepo *userStoryLikesRepo) FindOneByUserIDAndStoryID(ctx context.Context, UserID, StoryID uint) (*models.UserStoryLikes, error) {
	result, err := userRepo.Find(ctx, &models.UserStoryLikes{
		UserID:  UserID,
		StoryID: StoryID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (userRepo *userStoryLikesRepo) UpdateOne(ctx context.Context, storyLikes *models.UserStoryLikes) error {
	return userRepo.Update(ctx, storyLikes)
}

func (userRepo *userStoryLikesRepo) DeleteOne(ctx context.Context, ID uint) error {
	return userRepo.Delete(ctx, &models.UserStoryLikes{Base: models.Base{ID: ID}})
}

func (userStoryLikesRepo *userStoryLikesRepo) CountStoryLikes(
	ctx context.Context, storyID uint) (int64, error) {
	var count int64 = 0
	if err := userStoryLikesRepo.engine.WithContext(ctx).Debug().Model(models.UserStoryLikes{}).Where("story_ID = ?", storyID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

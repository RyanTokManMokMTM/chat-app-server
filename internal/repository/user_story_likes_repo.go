package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserStoryLikesRepo[T any] interface {
	CreateOne(ctx context.Context, user *T) error
	FindOneByID(ctx context.Context, id uint) (T, error)
	UpdateOne(ctx context.Context, data *T) error
	DeleteOne(ctx context.Context, id uint) error

	CountStoryLikes(ctx context.Context, storyId uint) (int64, error)
}

var _ IUserStoryLikesRepo[models.UserStoryLikes] = (*UserStoryLikesRepo)(nil)

type UserStoryLikesRepo struct {
	engine *gorm.DB
	IRepository[*models.UserStoryLikes, models.UserStoryLikes]
}

func NewUserStoryLikesRepo(engine *gorm.DB) *UserStoryLikesRepo {
	return &UserStoryLikesRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserStoryLikes, models.UserStoryLikes](engine),
	}
}

func (userRepo *UserStoryLikesRepo) CreateOne(ctx context.Context, data *models.UserStoryLikes) error {
	return userRepo.Create(ctx, data)
}

func (userRepo *UserStoryLikesRepo) FindOneByID(ctx context.Context, id uint) (models.UserStoryLikes, error) {
	return userRepo.Find(ctx, &models.UserStoryLikes{ID: id})
}

func (userRepo *UserStoryLikesRepo) UpdateOne(ctx context.Context, data *models.UserStoryLikes) error {
	return userRepo.Update(ctx, data)
}

func (userRepo *UserStoryLikesRepo) DeleteOne(ctx context.Context, id uint) error {
	return userRepo.Delete(ctx, &models.UserStoryLikes{ID: id})
}

func (userStoryLikesRepo *UserStoryLikesRepo) CountStoryLikes(
	ctx context.Context, storyId uint) (int64, error) {
	var count int64 = 0
	if err := userStoryLikesRepo.engine.WithContext(ctx).Debug().Model(models.UserStoryLikes{}).Where("story_id = ?", storyId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

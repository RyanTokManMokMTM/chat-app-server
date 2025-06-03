package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserStickerRepo[T any] interface {
	CreateOne(ctx context.Context, userStikcer T) (T, error)
}

var _ IUserStickerRepo[*models.UserSticker] = (*userStickerRepo)(nil)

type userStickerRepo struct {
	engine *gorm.DB
	IRepository[*models.UserSticker, models.UserSticker]
}

func NewUserStickerRepo(engine *gorm.DB) IUserStickerRepo[*models.UserSticker] {
	return &userStickerRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserSticker, models.UserSticker](engine),
	}
}

func (userStickerRepo *userStickerRepo) CreateOne(ctx context.Context, userStikcer *models.UserSticker) (*models.UserSticker, error) {
	return userStickerRepo.Create(ctx, userStikcer)
}

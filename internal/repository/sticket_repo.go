package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IStickerRepo[T any] interface {
	CreateOne(ctx context.Context, user *T) error
	FindOneByID(ctx context.Context, uuid string) (T, error)
	UpdateOne(ctx context.Context, data *T) error
	DeleteOne(ctx context.Context, uuid string) error
	FindAll(ctx context.Context) ([]T, error)
}

var _ IStickerRepo[models.Sticker] = (*StickerRepo)(nil)

type StickerRepo struct {
	engine *gorm.DB
	IRepository[*models.Sticker, models.Sticker]
}

func NewStickerRepo(engine *gorm.DB) IStickerRepo[models.Sticker] {
	return &StickerRepo{
		engine:      engine,
		IRepository: NewRepository[*models.Sticker, models.Sticker](engine),
	}
}

func (stickerRepo *StickerRepo) CreateOne(ctx context.Context, data *models.Sticker) error {
	return stickerRepo.Create(ctx, data)
}

func (stickerRepo *StickerRepo) FindOneByID(ctx context.Context, uuid string) (models.Sticker, error) {
	return stickerRepo.Find(ctx, &models.Sticker{Uuid: uuid})
}

func (stickerRepo *StickerRepo) UpdateOne(ctx context.Context, data *models.Sticker) error {
	return stickerRepo.Update(ctx, data)
}

func (stickerRepo *StickerRepo) DeleteOne(ctx context.Context, uuid string) error {
	return stickerRepo.Delete(ctx, &models.Sticker{Uuid: uuid})
}

func (stickerRepo *StickerRepo) FindAll(
	ctx context.Context) ([]models.Sticker, error) {
	resp := make([]models.Sticker, 0)
	if err := stickerRepo.engine.WithContext(ctx).Debug().Model(&resp).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

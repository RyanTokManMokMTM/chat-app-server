package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IStickerRepo[T any] interface {
	CreateOne(ctx context.Context, sticker T) (*T, error)
	FindOneByID(ctx context.Context, id uint) (*T, error)
	FindOneByUUID(ctx context.Context, uuid string) (*T, error)
	UpdateOne(ctx context.Context, sticker T) error
	DeleteOne(ctx context.Context, uuid string) error
	FindAll(ctx context.Context) ([]T, error)
	InsertResources(ctx context.Context, sticker *models.Sticker, paths []string) error
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

func (stickerRepo *StickerRepo) CreateOne(ctx context.Context, sticker models.Sticker) (*models.Sticker, error) {
	if err := stickerRepo.Create(ctx, &sticker); err != nil {
		return nil, err
	}
	return &sticker, nil
}

func (stickerRepo *StickerRepo) FindOneByID(ctx context.Context, id uint) (*models.Sticker, error) {
	result, err := stickerRepo.Find(ctx, &models.Sticker{Id: id})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (stickerRepo *StickerRepo) FindOneByUUID(ctx context.Context, uuid string) (*models.Sticker, error) {
	result, err := stickerRepo.Find(ctx, &models.Sticker{Uuid: uuid})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (stickerRepo *StickerRepo) UpdateOne(ctx context.Context, sticker models.Sticker) error {
	return stickerRepo.Update(ctx, &sticker)
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

func (stickerRepo *StickerRepo) InsertResources(ctx context.Context, sticker *models.Sticker, paths []string) error {
	return stickerRepo.engine.Transaction(func(tx *gorm.DB) error {
		resources := make([]*models.StickerResource, 0)
		for _, path := range paths {
			resource := &models.StickerResource{
				Path:    path,
				Sticker: *sticker,
			}
			//Create the resources
			if err := tx.WithContext(ctx).Create(resource).Error; err != nil {
				return err
			}

			resources = append(resources, resource)
		}

		return tx.WithContext(ctx).Debug().Model(sticker).Association("Resources").Append(resources)
	})
}

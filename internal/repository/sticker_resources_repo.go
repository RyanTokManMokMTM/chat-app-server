package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IStickerResourcesRepo[T any] interface {
	CreateOne(ctx context.Context, user *T) error
	FindOneByID(ctx context.Context, uuid string) (T, error)
	FindOneByUuid(ctx context.Context, uuid string) (T, error)
	UpdateOne(ctx context.Context, data *T) error
	DeleteOne(ctx context.Context, uuid string) error
	CreateMany(ctx context.Context, paths []string) ([]T, error)
}

var _ IStickerResourcesRepo[models.StickerResource] = (*StickerResourcesRepo)(nil)

type StickerResourcesRepo struct {
	engine *gorm.DB
	IRepository[*models.StickerResource, models.StickerResource]
}

func NewStickerResourcesRepo(engine *gorm.DB) *StickerResourcesRepo {
	return &StickerResourcesRepo{
		engine:      engine,
		IRepository: NewRepository[*models.StickerResource, models.StickerResource](engine),
	}
}

func (stickerResourcesRepo *StickerResourcesRepo) CreateOne(ctx context.Context, data *models.StickerResource) error {
	return stickerResourcesRepo.Create(ctx, data)
}

func (stickerResourcesRepo *StickerResourcesRepo) CreateMany(
	ctx context.Context,
	paths []string) ([]models.StickerResource, error) {
	resources := make([]models.StickerResource, 0)
	for _, path := range paths {
		resources = append(resources, models.StickerResource{
			Path: path,
		})
	}

	if err := stickerResourcesRepo.engine.WithContext(ctx).Create(resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (stickerResourcesRepo *StickerResourcesRepo) FindOneByID(ctx context.Context, uuid string) (models.StickerResource, error) {
	return stickerResourcesRepo.Find(ctx, &models.StickerResource{Uuid: uuid})
}

func (stickerResourcesRepo *StickerResourcesRepo) UpdateOne(ctx context.Context, data *models.StickerResource) error {
	return stickerResourcesRepo.Update(ctx, data)
}

func (stickerResourcesRepo *StickerResourcesRepo) DeleteOne(ctx context.Context, uuid string) error {
	return stickerResourcesRepo.Delete(ctx, &models.StickerResource{Uuid: uuid})
}

func (stickerResourcesRepo *StickerResourcesRepo) FindOneByUuid(ctx context.Context, uuid string) (models.StickerResource, error) {
	return stickerResourcesRepo.Find(ctx, &models.StickerResource{Uuid: uuid})
}

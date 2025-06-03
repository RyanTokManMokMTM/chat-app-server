package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IStickerResourcesRepo[T any] interface {
	CreateOne(ctx context.Context, stickerResource T) (T, error)
	FindOneByID(ctx context.Context, UUID string) (T, error)
	FindOneByUuID(ctx context.Context, UUID string) (T, error)

	UpdateOne(ctx context.Context, stickerResource T) error
	DeleteOne(ctx context.Context, UUID string) error
	CreateMany(ctx context.Context, paths []string) ([]T, error)
}

var _ IStickerResourcesRepo[*models.StickerResource] = (*stickerResourcesRepo)(nil)

type stickerResourcesRepo struct {
	engine *gorm.DB
	IRepository[*models.StickerResource, models.StickerResource]
}

func NewStickerResourcesRepo(engine *gorm.DB) IStickerResourcesRepo[*models.StickerResource] {
	return &stickerResourcesRepo{
		engine:      engine,
		IRepository: NewRepository[*models.StickerResource, models.StickerResource](engine),
	}
}

func (stickerResourcesRepo *stickerResourcesRepo) CreateOne(ctx context.Context, stickerResource *models.StickerResource) (*models.StickerResource, error) {
	return stickerResourcesRepo.Create(ctx, stickerResource)
}

func (stickerResourcesRepo *stickerResourcesRepo) CreateMany(
	ctx context.Context,
	paths []string) ([]*models.StickerResource, error) {
	resources := make([]*models.StickerResource, 0)
	for _, path := range paths {
		resources = append(resources, &models.StickerResource{
			Path: path,
		})
	}

	if err := stickerResourcesRepo.engine.WithContext(ctx).Create(resources).Error; err != nil {
		return nil, err
	}

	return resources, nil
}

func (stickerResourcesRepo *stickerResourcesRepo) FindOneByID(ctx context.Context, UUID string) (*models.StickerResource, error) {
	return stickerResourcesRepo.Find(ctx, &models.StickerResource{UUID: UUID})
}

func (stickerResourcesRepo *stickerResourcesRepo) UpdateOne(ctx context.Context, stickerResource *models.StickerResource) error {
	return stickerResourcesRepo.Update(ctx, stickerResource)
}

func (stickerResourcesRepo *stickerResourcesRepo) DeleteOne(ctx context.Context, UUID string) error {
	return stickerResourcesRepo.Delete(ctx, &models.StickerResource{UUID: UUID})
}

func (stickerResourcesRepo *stickerResourcesRepo) FindOneByUuID(ctx context.Context, UUID string) (*models.StickerResource, error) {
	return stickerResourcesRepo.Find(ctx, &models.StickerResource{UUID: UUID})
}

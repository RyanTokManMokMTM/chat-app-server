package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

// for any basic CRUD operations
type IRepository[T models.Model, R any] interface {
	// For creating a new record for T
	Create(context.Context, T) (*R, error)

	// For creating many records for T
	CreateMany(context.Context, []T) ([]*R, error)

	// For finding a record by ID for T
	Find(context.Context, T) (*R, error)

	// For finding many records for T
	FindMany(context.Context, T) ([]*R, error)

	// For updating a record for T
	Update(context.Context, T) error

	// For updating many records for T
	UpdateMany(context.Context, []T) error

	// For deleting a record for T
	Delete(context.Context, T) error

	// For deleting many records for T
	DeleteMany(context.Context, []T) error
}

type repository[T models.Model, R any] struct {
	engine *gorm.DB
}

func NewRepository[T models.Model, R any](engine *gorm.DB) IRepository[T, R] {
	return &repository[T, R]{
		engine: engine,
	}
}

func (repo *repository[T, R]) Create(ctx context.Context, data T) (*R, error) {
	if err := repo.engine.WithContext(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	result := *any(data).(*R)
	return &result, nil
}

func (repo *repository[T, R]) CreateMany(ctx context.Context, data []T) ([]*R, error) {
	if err := repo.engine.WithContext(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	// conver []T into []*R
	results := make([]*R, len(data))
	for i, item := range data {
		result := *any(item).(*R)
		results[i] = &result
	}
	return results, nil
}

func (repo *repository[T, R]) Find(ctx context.Context, data T) (*R, error) {
	var result R
	err := repo.engine.WithContext(ctx).First(&result, data).Error
	return &result, err
}

func (repo *repository[T, R]) FindMany(ctx context.Context, data T) ([]*R, error) {
	var results []R
	err := repo.engine.WithContext(ctx).Find(&results, data).Error
	if err != nil {
		return nil, err
	}

	// 將 []R 轉換為 []*R
	ptrResults := make([]*R, len(results))
	for i := range results {
		ptrResults[i] = &results[i]
	}
	return ptrResults, nil
}

func (repo *repository[T, R]) Update(ctx context.Context, data T) error {
	return repo.engine.WithContext(ctx).Updates(data).Error
}

func (repo *repository[T, R]) UpdateMany(ctx context.Context, t []T) error {
	return repo.engine.WithContext(ctx).Updates(t).Error
}

func (repo *repository[T, R]) Delete(ctx context.Context, data T) error {
	return repo.engine.WithContext(ctx).Delete(data).Error
}

func (repo *repository[T, R]) DeleteMany(ctx context.Context, t []T) error {
	return repo.engine.WithContext(ctx).Delete(t).Error
}

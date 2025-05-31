package repository

import (
	"context"

	"gorm.io/gorm"
)

// for any basic CRUD operations
type IRepository[T any, R any] interface {
	// For creating a new record for T
	Create(context.Context, T) error

	// For creating many records for T
	CreateMany(context.Context, []T) error

	// For finding a record by ID for T
	Find(context.Context, T) (R, error)

	// For finding many records for T
	FindMany(context.Context, T) ([]R, error)

	// For updating a record for T
	Update(context.Context, T) error

	// For updating many records for T
	UpdateMany(context.Context, []T) error

	// For deleting a record for T
	Delete(context.Context, T) error

	// For deleting many records for T
	DeleteMany(context.Context, []T) error
}

type Repository[T any, R any] struct {
	engine *gorm.DB
}

func NewRepository[T any, R any](engine *gorm.DB) IRepository[T, R] {
	return &Repository[T, R]{
		engine: engine,
	}
}

func (repo *Repository[T, R]) Create(ctx context.Context, data T) error {
	return repo.engine.WithContext(ctx).Create(data).Error
}

func (repo *Repository[T, R]) CreateMany(ctx context.Context, data []T) error {
	return repo.engine.WithContext(ctx).Create(data).Error
}

func (repo *Repository[T, R]) Find(ctx context.Context, data T) (R, error) {
	var result R
	err := repo.engine.WithContext(ctx).First(&result, data).Error
	return result, err
}

func (repo *Repository[T, R]) FindMany(ctx context.Context, data T) ([]R, error) {
	var results []R
	err := repo.engine.WithContext(ctx).Find(&results, data).Error
	return results, err
}

func (repo *Repository[T, R]) Update(ctx context.Context, data T) error {
	return repo.engine.WithContext(ctx).Updates(data).Error
}

func (repo *Repository[T, R]) UpdateMany(ctx context.Context, t []T) error {
	return repo.engine.WithContext(ctx).Updates(t).Error
}

func (repo *Repository[T, R]) Delete(ctx context.Context, data T) error {
	return repo.engine.WithContext(ctx).Delete(data).Error
}

func (repo *Repository[T, R]) DeleteMany(ctx context.Context, t []T) error {
	return repo.engine.WithContext(ctx).Delete(t).Error
}

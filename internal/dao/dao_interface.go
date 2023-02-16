package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

type DAOInterface interface {
	InsertOneUser(ctx context.Context, name, email, password string) error
	FindOneUser(ctx context.Context, id uint) (*models.UserModel, error)
}

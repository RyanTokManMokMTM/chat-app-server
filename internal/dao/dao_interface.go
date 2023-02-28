package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

type DAOInterface interface {
	InsertOneUser(ctx context.Context, name, email, password string) (*models.UserModel, error)
	FindOneUser(ctx context.Context, id uint) (*models.UserModel, error)
	FindOneUserByEmail(ctx context.Context, email string) (*models.UserModel, error)
	UpdateUserProfile(ctx context.Context, id uint, name string) error
	UpdateUserAvatar(ctx context.Context, id uint, avatarName string) error

	InsertOneFriend(ctx context.Context, friendID uint) error
	FindOneFriend(ctx context.Context, friendID uint) (bool, error)
	DeleteOneFriend(ctx context.Context, friendID uint) error
	GetUserFriendList(ctx context.Context) ([]*models.UserModel, error)
}

package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserFriendsRepo[T any] interface {
	CreateOne(ctx context.Context, userFriend T) (T, error)
	FindOneByID(ctx context.Context, ID uint) (T, error)
	FindOneByUserIDAndFriendID(ctx context.Context, userID, friendID uint) (T, error)
	UpdateOne(ctx context.Context, userFriend T) error
	DeleteOneByUserIDAndFriendID(ctx context.Context, userID, friendID uint) error

	CountUserFriends(ctx context.Context, UserID uint) (int64, error)
	GetFriendList(
		ctx context.Context,
		UserID uint,
		pageOffset,
		pageSize int) ([]T, error)
}

var _ IUserFriendsRepo[*models.UserFriend] = (*userFriendsRepo)(nil)

type userFriendsRepo struct {
	engine *gorm.DB
	IRepository[*models.UserFriend, models.UserFriend]
}

func NewUserFriendsRepo(engine *gorm.DB) IUserFriendsRepo[*models.UserFriend] {
	return &userFriendsRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserFriend, models.UserFriend](engine),
	}
}

func (userFriendsRepo *userFriendsRepo) CreateOne(ctx context.Context, userFriend *models.UserFriend) (*models.UserFriend, error) {
	return userFriendsRepo.Create(ctx, userFriend)
}

func (userFriendsRepo *userFriendsRepo) FindOneByID(ctx context.Context, ID uint) (*models.UserFriend, error) {
	return userFriendsRepo.Find(ctx, &models.UserFriend{Base: models.Base{ID: ID}})
}

func (userFriendsRepo *userFriendsRepo) FindOneByUserIDAndFriendID(ctx context.Context, userID, friendID uint) (*models.UserFriend, error) {
	result, err := userFriendsRepo.Find(ctx, &models.UserFriend{
		UserID:   userID,
		FriendID: friendID,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (userFriendsRepo *userFriendsRepo) UpdateOne(ctx context.Context, userFriend *models.UserFriend) error {
	return userFriendsRepo.Update(ctx, userFriend)
}

func (userFriendsRepo *userFriendsRepo) DeleteOne(ctx context.Context, ID uint) error {
	return userFriendsRepo.Delete(ctx, &models.UserFriend{Base: models.Base{
		ID: ID,
	}})
}

func (userFriendsRepo *userFriendsRepo) DeleteOneByUserIDAndFriendID(ctx context.Context, userID, friendID uint) error {
	return userFriendsRepo.Delete(ctx, &models.UserFriend{
		UserID:   userID,
		FriendID: friendID,
	})
}

func (userFriendsRepo *userFriendsRepo) GetFriendList(
	ctx context.Context,
	UserID uint,
	pageOffset,
	pageSize int) ([]*models.UserFriend, error) {
	var list []*models.UserFriend
	if err := userFriendsRepo.engine.WithContext(ctx).Debug().Model(&models.UserFriend{}).
		Preload("FriendInfo").
		Where("user_ID = ?", UserID).
		Offset(pageOffset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
func (userFriendsRepo *userFriendsRepo) CountUserFriends(ctx context.Context, UserID uint) (int64, error) {
	var count int64 = 0
	if err := userFriendsRepo.engine.WithContext(ctx).Debug().Model(&models.UserFriend{}).
		Where("user_ID = ?", UserID).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

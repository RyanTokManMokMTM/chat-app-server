package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserFriendsRepo[T any] interface {
	CreateOne(ctx context.Context, user *T) error
	FindOneByID(ctx context.Context, id uint) (T, error)
	FindOneByUserIdAndFriendById(ctx context.Context, userId, frinedId uint) (T, error)
	UpdateOne(ctx context.Context, data *T) error
	DeleteOne(ctx context.Context, id uint) error
}

var _ IUserFriendsRepo[models.UserFriend] = (*UserFriendsRepo)(nil)

type UserFriendsRepo struct {
	engine *gorm.DB
	IRepository[*models.UserFriend, models.UserFriend]
}

func NewUserFriendsRepo(engine *gorm.DB) *UserFriendsRepo {
	return &UserFriendsRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserFriend, models.UserFriend](engine),
	}
}

func (userFriendsRepo *UserFriendsRepo) CreateOne(ctx context.Context, data *models.UserFriend) error {
	return userFriendsRepo.Create(ctx, data)
}

func (userFriendsRepo *UserFriendsRepo) FindOneByID(ctx context.Context, id uint) (models.UserFriend, error) {
	return userFriendsRepo.Find(ctx, &models.UserFriend{ID: id})
}

func (userFriendsRepo *UserFriendsRepo) FindOneByUserIdAndFriendById(ctx context.Context, userId, frinedId uint) (models.UserFriend, error) {
	return userFriendsRepo.Find(ctx, &models.UserFriend{
		UserID:   userId,
		FriendID: frinedId,
	})
}

func (userFriendsRepo *UserFriendsRepo) UpdateOne(ctx context.Context, data *models.UserFriend) error {
	return userFriendsRepo.Update(ctx, data)
}

func (userFriendsRepo *UserFriendsRepo) DeleteOne(ctx context.Context, id uint) error {
	return userFriendsRepo.Delete(ctx, &models.UserFriend{ID: id})
}

func (userFriendsRepo *UserFriendsRepo) GetFriendList(
	ctx context.Context,
	userId uint,
	pageOffset,
	pageSize int) ([]*models.UserFriend, error) {
	var list []*models.UserFriend
	if err := userFriendsRepo.engine.WithContext(ctx).Debug().Model(&models.UserFriend{}).
		Preload("FriendInfo").
		Where("user_id = ?", userId).
		Offset(pageOffset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}

func (userFriendsRepo *UserFriendsRepo) CountUserFriends(ctx context.Context, userId uint) (int64, error) {
	var count int64 = 0
	if err := userFriendsRepo.engine.WithContext(ctx).Debug().Model(&models.UserFriend{}).
		Where("user_id = ?", userId).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

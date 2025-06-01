package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserFriendsRepo[T any] interface {
	CreateOne(ctx context.Context, userFriend models.UserFriend) (*models.UserFriend, error)
	FindOneByID(ctx context.Context, id uint) (*models.UserFriend, error)
	FindOneByUserIdAndFriendId(ctx context.Context, userId, friendId uint) (*models.UserFriend, error)
	UpdateOne(ctx context.Context, userFriend models.UserFriend) error
	DeleteOneByUserIdAndFriendID(ctx context.Context, userId, friendId uint) error

	CountUserFriends(ctx context.Context, UserId uint) (int64, error)
	GetFriendList(
		ctx context.Context,
		UserId uint,
		pageOffset,
		pageSize int) ([]*models.UserFriend, error)
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

func (userFriendsRepo *UserFriendsRepo) CreateOne(ctx context.Context, userFriend models.UserFriend) (*models.UserFriend, error) {
	if err := userFriendsRepo.Create(ctx, &userFriend); err != nil {
		return nil, err
	}
	return &userFriend, nil
}

func (userFriendsRepo *UserFriendsRepo) FindOneByID(ctx context.Context, id uint) (*models.UserFriend, error) {
	result, err := userFriendsRepo.Find(ctx, &models.UserFriend{ID: id})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (userFriendsRepo *UserFriendsRepo) FindOneByUserIdAndFriendId(ctx context.Context, userId, friendId uint) (*models.UserFriend, error) {
	result, err := userFriendsRepo.Find(ctx, &models.UserFriend{
		UserId:   userId,
		FriendID: friendId,
	})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (userFriendsRepo *UserFriendsRepo) UpdateOne(ctx context.Context, userFriend models.UserFriend) error {
	return userFriendsRepo.Update(ctx, &userFriend)
}

func (userFriendsRepo *UserFriendsRepo) DeleteOne(ctx context.Context, id uint) error {
	return userFriendsRepo.Delete(ctx, &models.UserFriend{ID: id})
}

func (userFriendsRepo *UserFriendsRepo) DeleteOneByUserIdAndFriendID(ctx context.Context, userId, friendId uint) error {
	return userFriendsRepo.Delete(ctx, &models.UserFriend{
		UserId:   userId,
		FriendID: friendId,
	})
}

func (userFriendsRepo *UserFriendsRepo) GetFriendList(
	ctx context.Context,
	UserId uint,
	pageOffset,
	pageSize int) ([]*models.UserFriend, error) {
	var list []*models.UserFriend
	if err := userFriendsRepo.engine.WithContext(ctx).Debug().Model(&models.UserFriend{}).
		Preload("FriendInfo").
		Where("user_id = ?", UserId).
		Offset(pageOffset).
		Limit(pageSize).
		Find(&list).Error; err != nil {
		return nil, err
	}

	return list, nil
}
func (userFriendsRepo *UserFriendsRepo) CountUserFriends(ctx context.Context, UserId uint) (int64, error) {
	var count int64 = 0
	if err := userFriendsRepo.engine.WithContext(ctx).Debug().Model(&models.UserFriend{}).
		Where("user_id = ?", UserId).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

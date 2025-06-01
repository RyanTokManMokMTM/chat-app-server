package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserGroupRepo[T any] interface {
	CreateOne(ctx context.Context, userGroup T) (*T, error)
	FindOneById(ctx context.Context, id uint) (*T, error)
	FindOneByGroupIdAndUserId(ctx context.Context, groupId, userId uint) (*T, error)
	UpdateOne(ctx context.Context, userGroup T) error
	DeleteOne(ctx context.Context, id uint) error
	DeleteOneByGroupIdAndUserId(ctx context.Context, groupId, userId uint) error
	DeleteAll(ctx context.Context, groupId uint) error

	GetGroupMemberList(ctx context.Context, groupId uint) ([]models.UserGroup, error)
	GetGroupMemberListByPage(ctx context.Context, groupId uint, pageOffset, pageLimit int) ([]models.UserGroup, error)
	FindUserGroup(ctx context.Context, userId uint, pageOffset, pageSize int) ([]models.UserGroup, error)
	CountGroupMembers(ctx context.Context, groupId uint) (int64, error)
}

var _ IUserGroupRepo[models.UserGroup] = (*UserGroupRepo)(nil)

type UserGroupRepo struct {
	engine *gorm.DB
	IRepository[*models.UserGroup, models.UserGroup]
}

func NewUserGroupRepo(engine *gorm.DB) *UserGroupRepo {
	return &UserGroupRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserGroup, models.UserGroup](engine),
	}
}

func (userGroupRepo *UserGroupRepo) CreateOne(ctx context.Context, userGroup models.UserGroup) (*models.UserGroup, error) {
	if err := userGroupRepo.Create(ctx, &userGroup); err != nil {
		return nil, err
	}
	return &userGroup, nil
}

func (userGroupRepo *UserGroupRepo) FindOneById(ctx context.Context, id uint) (*models.UserGroup, error) {
	result, err := userGroupRepo.Find(ctx, &models.UserGroup{ID: id})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (userGroupRepo *UserGroupRepo) UpdateOne(ctx context.Context, userGroup models.UserGroup) error {
	return userGroupRepo.Update(ctx, &userGroup)
}

func (userGroupRepo *UserGroupRepo) DeleteOne(ctx context.Context, id uint) error {
	return userGroupRepo.Delete(ctx, &models.UserGroup{ID: id})
}

func (userGroupRepo *UserGroupRepo) DeleteAll(ctx context.Context, groupId uint) error {
	return userGroupRepo.engine.WithContext(ctx).Debug().Where("group_id = ?", groupId).Delete(&models.UserGroup{}).Error
}

func (userGroupRepo *UserGroupRepo) GetGroupMemberList(ctx context.Context, groupId uint) ([]models.UserGroup, error) {
	var members []models.UserGroup
	if err := userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_id = ?", groupId).
		Preload("MemberInfo").Preload("GroupInfo").
		Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (userGroupRepo *UserGroupRepo) GetGroupMemberListByPage(ctx context.Context, groupId uint, pageOffset, pageLimit int) ([]models.UserGroup, error) {
	var members []models.UserGroup
	if err := userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_id = ?", groupId).
		Preload("MemberInfo").Preload("GroupInfo").
		Offset(pageOffset).
		Limit(pageLimit).
		Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (userGroupRepo *UserGroupRepo) FindUserGroup(ctx context.Context, userId uint, pageOffset, pageSize int) ([]models.UserGroup, error) {
	var groups []models.UserGroup
	if err := userGroupRepo.engine.
		WithContext(ctx).
		Debug().
		Preload("GroupInfo").
		Where("user_id = ?", userId).Offset(pageOffset).Limit(pageSize).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (userGroupRepo *UserGroupRepo) CountGroupMembers(ctx context.Context, groupId uint) (int64, error) {
	var count int64
	if err := userGroupRepo.engine.
		WithContext(ctx).
		Debug().
		Model(models.UserGroup{}).Where("group_id = ?", groupId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (userGroupRepo *UserGroupRepo) FindOneByGroupIdAndUserId(ctx context.Context, groupId, userId uint) (*models.UserGroup, error) {
	var result models.UserGroup
	if err := userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_id = ? AND user_id = ?", groupId, userId).
		First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (userGroupRepo *UserGroupRepo) DeleteOneByGroupIdAndUserId(ctx context.Context, groupId, userId uint) error {
	return userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_id = ? AND user_id = ?", groupId, userId).
		Delete(&models.UserGroup{}).Error
}

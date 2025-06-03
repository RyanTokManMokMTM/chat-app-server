package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IUserGroupRepo[T any] interface {
	CreateOne(ctx context.Context, userGroup T) (T, error)
	FindOneByID(ctx context.Context, ID uint) (T, error)
	FindOneByGroupIDAndUserID(ctx context.Context, groupID, userID uint) (T, error)
	UpdateOne(ctx context.Context, userGroup T) error
	DeleteOne(ctx context.Context, ID uint) error
	DeleteOneByGroupIDAndUserID(ctx context.Context, groupID, userID uint) error
	DeleteAll(ctx context.Context, groupID uint) error

	GetGroupMemberList(ctx context.Context, groupID uint) ([]T, error)
	GetGroupMemberListByPage(ctx context.Context, groupID uint, pageOffset, pageLimit int) ([]T, error)
	FindUserGroup(ctx context.Context, userID uint, pageOffset, pageSize int) ([]T, error)
	CountGroupMembers(ctx context.Context, groupID uint) (int64, error)
}

var _ IUserGroupRepo[*models.UserGroup] = (*userGroupRepo)(nil)

type userGroupRepo struct {
	engine *gorm.DB
	IRepository[*models.UserGroup, models.UserGroup]
}

func NewUserGroupRepo(engine *gorm.DB) IUserGroupRepo[*models.UserGroup] {
	return &userGroupRepo{
		engine:      engine,
		IRepository: NewRepository[*models.UserGroup, models.UserGroup](engine),
	}
}

func (userGroupRepo *userGroupRepo) CreateOne(ctx context.Context, userGroup *models.UserGroup) (*models.UserGroup, error) {
	return userGroupRepo.Create(ctx, userGroup)
}

func (userGroupRepo *userGroupRepo) FindOneByID(ctx context.Context, ID uint) (*models.UserGroup, error) {
	result, err := userGroupRepo.Find(ctx, &models.UserGroup{Base: models.Base{ID: ID}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (userGroupRepo *userGroupRepo) UpdateOne(ctx context.Context, userGroup *models.UserGroup) error {
	return userGroupRepo.Update(ctx, userGroup)
}

func (userGroupRepo *userGroupRepo) DeleteOne(ctx context.Context, ID uint) error {
	return userGroupRepo.Delete(ctx, &models.UserGroup{Base: models.Base{ID: ID}})
}

func (userGroupRepo *userGroupRepo) DeleteAll(ctx context.Context, groupID uint) error {
	return userGroupRepo.engine.WithContext(ctx).Debug().Where("group_ID = ?", groupID).Delete(&models.UserGroup{}).Error
}

func (userGroupRepo *userGroupRepo) GetGroupMemberList(ctx context.Context, groupID uint) ([]*models.UserGroup, error) {
	var members []*models.UserGroup
	if err := userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_ID = ?", groupID).
		Preload("MemberInfo").Preload("GroupInfo").
		Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (userGroupRepo *userGroupRepo) GetGroupMemberListByPage(ctx context.Context, groupID uint, pageOffset, pageLimit int) ([]*models.UserGroup, error) {
	var members []*models.UserGroup
	if err := userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_ID = ?", groupID).
		Preload("MemberInfo").Preload("GroupInfo").
		Offset(pageOffset).
		Limit(pageLimit).
		Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (userGroupRepo *userGroupRepo) FindUserGroup(ctx context.Context, userID uint, pageOffset, pageSize int) ([]*models.UserGroup, error) {
	var groups []*models.UserGroup
	if err := userGroupRepo.engine.
		WithContext(ctx).
		Debug().
		Preload("GroupInfo").
		Where("user_ID = ?", userID).Offset(pageOffset).Limit(pageSize).Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (userGroupRepo *userGroupRepo) CountGroupMembers(ctx context.Context, groupID uint) (int64, error) {
	var count int64
	if err := userGroupRepo.engine.
		WithContext(ctx).
		Debug().
		Model(models.UserGroup{}).Where("group_ID = ?", groupID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (userGroupRepo *userGroupRepo) FindOneByGroupIDAndUserID(ctx context.Context, groupID, userID uint) (*models.UserGroup, error) {
	var result models.UserGroup
	if err := userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_ID = ? AND user_ID = ?", groupID, userID).
		First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (userGroupRepo *userGroupRepo) DeleteOneByGroupIDAndUserID(ctx context.Context, groupID, userID uint) error {
	return userGroupRepo.engine.WithContext(ctx).Debug().
		Where("group_ID = ? AND user_ID = ?", groupID, userID).
		Delete(&models.UserGroup{}).Error
}

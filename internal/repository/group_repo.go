package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IGroupRepo[T any] interface {
	CreateOne(ctx context.Context, group T) (*T, error)
	FindOneById(ctx context.Context, id uint) (*T, error)
	FindOneByUUID(ctx context.Context, uuid string) (*T, error)
	FindOneByGroupIdAndUserId(ctx context.Context, groupId, userId uint) (*T, error)
	UpdateOne(ctx context.Context, group T) error
	DeleteOne(ctx context.Context, uuid string) error
	DeleteOneById(ctx context.Context, id uint) error

	UpdateOneByNameAndDesc(ctx context.Context, id uint, name, desc string) error
	UpdateOneAvatar(ctx context.Context, id uint, avatarPath string) error
	FindOneByQuery(ctx context.Context, query string) ([]*models.Group, error)
}

var _ IGroupRepo[models.Group] = (*GroupRepo)(nil)

type GroupRepo struct {
	engine *gorm.DB
	IRepository[*models.Group, models.Group]
}

func NewGroupRepo(engine *gorm.DB) IGroupRepo[models.Group] {
	return &GroupRepo{
		IRepository: NewRepository[*models.Group, models.Group](engine), // repo for group
		engine:      engine,
	}
}

func (groupRepo *GroupRepo) CreateOne(ctx context.Context, group models.Group) (*models.Group, error) {
	if err := groupRepo.Create(ctx, &group); err != nil {
		return nil, err
	}
	return &group, nil
}

func (groupRepo *GroupRepo) FindOneById(ctx context.Context, id uint) (*models.Group, error) {
	result, err := groupRepo.Find(ctx, &models.Group{Id: id})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (groupRepo *GroupRepo) FindOneByUUID(ctx context.Context, uuid string) (*models.Group, error) {
	result, err := groupRepo.Find(ctx, &models.Group{Uuid: uuid})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (groupRepo *GroupRepo) UpdateOne(ctx context.Context, group models.Group) error {
	return groupRepo.Update(ctx, &group)
}

func (groupRepo *GroupRepo) DeleteOne(ctx context.Context, uuid string) error {
	return groupRepo.Delete(ctx, &models.Group{Uuid: uuid})
}

func (groupRepo *GroupRepo) DeleteOneById(ctx context.Context, id uint) error {
	return groupRepo.Delete(ctx, &models.Group{Id: id})
}

// Other

func (groupRepo *GroupRepo) FindOneByUUID(ctx context.Context, uuid string) (*models.Group, error) {
	result, err := groupRepo.Find(ctx, &models.Group{Uuid: uuid})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (groupRepo *GroupRepo) UpdateOneByNameAndDesc(ctx context.Context, id uint, name, desc string) error {
	group := models.Group{
		Id:        id,
		GroupName: name,
		GroupDesc: desc,
	}
	return groupRepo.engine.WithContext(ctx).Debug().Model(models.Group{}).Where("id = ?", group.Id).UpdateColumns(map[string]any{
		"GroupName": group.GroupName,
		"GroupDesc": group.GroupDesc,
	}).Error
}

func (groupRepo *GroupRepo) UpdateOneAvatar(ctx context.Context, id uint, avatarPath string) error {
	group := models.Group{
		Id:          id,
		GroupAvatar: avatarPath,
	}
	return groupRepo.engine.WithContext(ctx).Debug().Model(group).Where("id = ?", group.Id).Update("GroupAvatar", group.GroupAvatar).Error
}

func (groupRepo *GroupRepo) FindOneByQuery(ctx context.Context, query string) ([]*models.Group, error) {
	var groups []*models.Group
	if err := groupRepo.
		engine.
		WithContext(ctx).
		Debug().
		Model(models.Group{}).
		Where("group_name like ?", "%"+query+"%").Preload("LeadInfo").Find(&groups).Error; err != nil {
		return nil, err
	}
	return groups, nil
}

func (groupRepo *GroupRepo) FindOneByGroupIdAndUserId(ctx context.Context, groupId, userId uint) (*models.Group, error) {
	var group models.Group
	err := groupRepo.engine.WithContext(ctx).Debug().Where("group_id = ? AND user_id = ?", groupId, userId).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

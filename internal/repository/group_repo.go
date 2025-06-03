package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IGroupRepo[T models.Model] interface {
	CreateOne(ctx context.Context, group T) (T, error)
	FindOneByID(ctx context.Context, ID uint) (T, error)
	FindOneByUUID(ctx context.Context, UUID string) (T, error)
	FindOneByGroupIDAndUserID(ctx context.Context, groupID, userID uint) (T, error)
	UpdateOne(ctx context.Context, group T) error
	DeleteOne(ctx context.Context, UUID string) error
	DeleteOneByID(ctx context.Context, ID uint) error

	UpdateOneByNameAndDesc(ctx context.Context, ID uint, name, desc string) error
	UpdateOneAvatar(ctx context.Context, ID uint, avatarPath string) error
	FindOneByQuery(ctx context.Context, query string) ([]T, error)
}

var _ IGroupRepo[*models.Group] = (*groupRepo)(nil)

type groupRepo struct {
	engine *gorm.DB
	IRepository[*models.Group, models.Group]
}

func NewGroupRepo(engine *gorm.DB) IGroupRepo[*models.Group] {
	return &groupRepo{
		IRepository: NewRepository[*models.Group, models.Group](engine),
		engine:      engine,
	}
}

func (groupRepo *groupRepo) CreateOne(ctx context.Context, group *models.Group) (*models.Group, error) {
	return groupRepo.Create(ctx, group)
}

func (groupRepo *groupRepo) FindOneByID(ctx context.Context, ID uint) (*models.Group, error) {
	return groupRepo.Find(ctx, &models.Group{Base: models.Base{
		ID: ID,
	}})
}

func (groupRepo *groupRepo) FindOneByUUID(ctx context.Context, UUID string) (*models.Group, error) {
	return groupRepo.Find(ctx, &models.Group{
		Base: models.Base{
			UUID: UUID,
		}})
}

func (groupRepo *groupRepo) UpdateOne(ctx context.Context, group *models.Group) error {
	return groupRepo.Update(ctx, group)
}

func (groupRepo *groupRepo) DeleteOne(ctx context.Context, UUID string) error {
	return groupRepo.Delete(ctx, &models.Group{Base: models.Base{UUID: UUID}})
}

func (groupRepo *groupRepo) DeleteOneByID(ctx context.Context, ID uint) error {
	return groupRepo.Delete(ctx, &models.Group{Base: models.Base{ID: ID}})
}

// Other

func (groupRepo *groupRepo) UpdateOneByNameAndDesc(ctx context.Context, ID uint, name, desc string) error {
	group := models.Group{
		Base: models.Base{
			ID: ID,
		},
		GroupName: name,
		GroupDesc: desc,
	}
	return groupRepo.engine.WithContext(ctx).Debug().Model(models.Group{}).Where("ID = ?", group.ID).UpdateColumns(map[string]any{
		"GroupName": group.GroupName,
		"GroupDesc": group.GroupDesc,
	}).Error
}

func (groupRepo *groupRepo) UpdateOneAvatar(ctx context.Context, ID uint, avatarPath string) error {
	group := models.Group{
		Base: models.Base{
			ID: ID,
		},
		GroupAvatar: avatarPath,
	}
	return groupRepo.engine.WithContext(ctx).Debug().Model(group).Where("ID = ?", group.ID).Update("GroupAvatar", group.GroupAvatar).Error
}

func (groupRepo *groupRepo) FindOneByQuery(ctx context.Context, query string) ([]*models.Group, error) {
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

func (groupRepo *groupRepo) FindOneByGroupIDAndUserID(ctx context.Context, groupID, userID uint) (*models.Group, error) {
	var group models.Group
	err := groupRepo.engine.WithContext(ctx).Debug().Where("group_ID = ? AND user_ID = ?", groupID, userID).First(&group).Error
	if err != nil {
		return nil, err
	}
	return &group, nil
}

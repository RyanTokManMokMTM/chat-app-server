package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
)

func (d *DAO) InsertOneGroup(ctx context.Context, groupName, avatar string, userID uint) (*models.Group, error) {
	g := &models.Group{
		GroupName:   groupName,
		GroupLead:   userID,
		GroupAvatar: avatar, // leave it empty now
	}

	if err := g.InsertOne(ctx, d.engine); err != nil {
		return nil, err
	}

	return g, nil
}

func (d *DAO) FindOneGroup(ctx context.Context, groupID uint) (*models.Group, error) {
	g := &models.Group{
		ID: groupID,
	}

	if err := g.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}

	return g, nil
}

func (d *DAO) FindOneGroupByUUID(ctx context.Context, groupUUID string) (*models.Group, error) {
	g := &models.Group{
		Uuid: groupUUID,
	}

	if err := g.FindOneByUUID(ctx, d.engine); err != nil {
		return nil, err
	}

	return g, nil
}

func (d *DAO) DeleteOneGroup(ctx context.Context, groupID uint) error {
	g := &models.Group{
		ID: groupID,
	}

	return g.DeleteOne(ctx, d.engine)
}

func (d *DAO) UpdateOneGroup(ctx context.Context, groupID uint, groupName string) error {
	group := &models.Group{
		ID:        groupID,
		GroupName: groupName,
	}
	return group.UpdateOne(ctx, d.engine)
}

func (d *DAO) UpdateOneGroupAvatar(ctx context.Context, groupID uint, avatarName string) error {
	g := &models.Group{
		ID:          groupID,
		GroupAvatar: avatarName,
	}

	return g.UpdateOneAvatar(ctx, d.engine)
}

func (d *DAO) SearchGroup(ctx context.Context, query string) ([]*models.Group, error) {
	g := &models.Group{}
	return g.SearchGroup(ctx, d.engine, query)
}

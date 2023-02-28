package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

func (d *DAO) InsertOneGroup(ctx context.Context, groupName string, userID uint) (*models.Group, error) {
	g := &models.Group{
		GroupName:   groupName,
		GroupLead:   userID,
		GroupAvatar: "", // leave it empty now
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
func (d *DAO) DeleteOneGroup(ctx context.Context, groupID uint) error {
	g := &models.Group{
		ID: groupID,
	}

	return g.DeleteOne(ctx, d.engine)
}

func (d *DAO) UpdateOneGroup(ctx context.Context, g *models.Group) error {
	return g.UpdateOne(ctx, d.engine)
}

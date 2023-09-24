package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
)

func (d *DAO) InsertOneUserStorySeen(ctx context.Context, userID, friendId, storyId uint) error {
	model := &models.UserStorySeen{
		UserId:       userID,
		FriendId:     friendId,
		StoryModelID: storyId,
	}
	return model.InsertOne(ctx, d.engine)
}

func (d *DAO) FindOneUserStorySeen(ctx context.Context, userID, friendId uint) (*models.UserStorySeen, error) {
	model := &models.UserStorySeen{
		UserId:   userID,
		FriendId: friendId,
	}
	if err := model.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DAO) UpdateOneUserStorySeen(ctx context.Context, Id, storyId uint) error {
	model := &models.UserStorySeen{
		ID: Id,
	}
	return model.UpdateOne(ctx, d.engine, storyId)
}

func (d *DAO) DeleteOneUserStorySeen(ctx context.Context, ID uint) error {
	model := &models.UserStorySeen{
		ID: ID,
	}
	return model.DeleteOne(ctx, d.engine)
}

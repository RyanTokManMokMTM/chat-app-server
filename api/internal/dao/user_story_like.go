package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
)

func (d *DAO) InsertOneUserStoryLike(ctx context.Context, userID, storyId uint) error {
	model := &models.UserStoryLikes{
		UserId:  userID,
		StoryId: storyId,
	}

	return model.InsertOne(ctx, d.engine)
}

func (d *DAO) FindOneUserStoryLike(ctx context.Context, userID, storyId uint) (*models.UserStoryLikes, error) {
	model := &models.UserStoryLikes{
		UserId:  userID,
		StoryId: storyId,
	}

	if err := model.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DAO) DeleteOneUserStoryLike(ctx context.Context, ID uint) error {
	model := &models.UserStoryLikes{
		ID: ID,
	}
	return model.DeleteOne(ctx, d.engine)
}

func (d *DAO) CountStoryLikes(ctx context.Context, storyId uint) (int64, error) {
	model := &models.UserStoryLikes{
		StoryId: storyId,
	}

	return model.CountStoryLikes(ctx, d.engine)
}

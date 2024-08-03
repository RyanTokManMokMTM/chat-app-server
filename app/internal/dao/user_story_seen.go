package dao

import (
	"api/app/internal/models"
	"context"
)

func (d *DAO) InsertOneUserStorySeen(ctx context.Context, userID, friendId, storyId uint) error {
	model := &models.UserStorySeen{
		UserId:   userID,
		FriendId: friendId,
		StoryId:  storyId,
	}
	return model.InsertOne(ctx, d.engine)
}

func (d *DAO) FindOneUserStorySeen(ctx context.Context, userID, friendId, storyId uint) (*models.UserStorySeen, error) {
	model := &models.UserStorySeen{
		UserId:   userID,
		FriendId: friendId,
		StoryId:  storyId,
	}
	if err := model.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return model, nil
}

func (d *DAO) FindOneLatestUserStorySeen(ctx context.Context, userID, friendId uint) (*models.UserStorySeen, error) {
	model := &models.UserStorySeen{
		UserId:   userID,
		FriendId: friendId,
	}
	if err := model.FindLatestOne(ctx, d.engine); err != nil {
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

func (d *DAO) GetStorySeenUserList(ctx context.Context, storyId uint, limit int) ([]*models.UserStorySeen, error) {
	model := &models.UserStorySeen{
		StoryId: storyId,
	}

	return model.GetStoryLikedUserSeen(ctx, d.engine, limit)
}

func (d *DAO) CountOneStorySeen(ctx context.Context, storyId uint) (int64, error) {
	model := &models.UserStorySeen{
		StoryId: storyId,
	}
	return model.CountOneStorySeen(ctx, d.engine)
}

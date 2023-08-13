package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
)

func (d *DAO) InsertOneStory(ctx context.Context, userID uint, mediaPath string) (uint, error) {
	s := &models.StoryModel{
		UserId:         userID,
		StoryMediaPath: mediaPath,
	}

	if err := s.InsertOne(ctx, d.engine); err != nil {
		return 0, err
	}
	return s.Id, nil
}

func (d *DAO) FindOneStory(ctx context.Context, storyID uint) (*models.StoryModel, error) {
	s := &models.StoryModel{
		Id: storyID,
	}
	if err := s.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return s, nil
}

func (d *DAO) FindOneUserStory(ctx context.Context, storyID, userID uint) (*models.StoryModel, error) {
	s := &models.StoryModel{
		Id:     storyID,
		UserId: userID,
	}

	if err := s.FindOneUserStory(ctx, d.engine); err != nil {
		return nil, err
	}
	return s, nil
}

func (d *DAO) GetUserStories(ctx context.Context, userID uint) ([]uint, error) {
	s := &models.StoryModel{
		UserId: userID,
	}

	return s.FindAllUserStories(ctx, d.engine)
}

func (d *DAO) GetFriendActiveStories(ctx context.Context, userID uint, pageOffset, pageLimit int) ([]*models.StoriesWithIds, error) {
	s := &models.StoryModel{}
	return s.GetFriendStoryList(ctx, d.engine, userID, pageOffset, pageLimit)
}

func (d *DAO) CountActiveStory(ctx context.Context, userId uint) (int64, error) {
	s := &models.StoryModel{}
	return s.CountFriendActiveStory(ctx, d.engine, userId)
}

func (d *DAO) DeleteStories(ctx context.Context, storyID uint) error {
	s := &models.StoryModel{
		Id: storyID,
	}
	return s.DeleteOne(ctx, d.engine)
}

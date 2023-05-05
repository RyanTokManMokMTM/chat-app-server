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

func (d *DAO) GetActiveUsers(ctx context.Context, userID uint) ([]*models.UserModel, error) {
	//find all friend

	return nil, nil
}
func (d *DAO) DeleteStories(ctx context.Context, storyID uint) error {
	s := &models.StoryModel{
		Id: storyID,
	}
	return s.DeleteOne(ctx, d.engine)
}

package dao

import (
	"api/app/internal/models"
	"context"
)

func (d *DAO) InsertOneStory(ctx context.Context, userID uint, mediaPath string) (*models.StoryModel, error) {
	s := &models.StoryModel{
		UserId:         userID,
		StoryMediaPath: mediaPath,
	}

	if err := s.InsertOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return s, nil
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

func (d *DAO) GetUserStoriesByTimeStamp(ctx context.Context, userID uint, timeStamp int64) ([]*models.StoryModel, error) {
	s := &models.StoryModel{
		UserId: userID,
	}

	return s.FindAllUserStoriesByTimeStamp(ctx, d.engine, timeStamp)
}

func (d *DAO) GetFriendActiveStories(ctx context.Context, userID uint, pageOffset, pageLimit int) ([]*models.StoriesWithLatestStoryTime, error) {
	s := &models.StoryModel{}
	return s.GetActiveStoryList(ctx, d.engine, userID, pageOffset, pageLimit)
}

func (d *DAO) GetFriendActiveStoriesByTimeStamp(ctx context.Context, userID uint, pageOffset, pageLimit int, timeStamp int64) ([]*models.StoriesWithLatestStoryTime, error) {
	s := &models.StoryModel{}
	return s.GetActiveStoryListByTime(ctx, d.engine, userID, pageOffset, pageLimit, timeStamp)
}

func (d *DAO) CountActiveStoryByTimeStamp(ctx context.Context, userId uint, timeStamp int64) (int64, error) {
	s := &models.StoryModel{}
	return s.CountFriendActiveStoryByTime(ctx, d.engine, userId, timeStamp)
}

func (d *DAO) DeleteStories(ctx context.Context, storyID uint) error {
	s := &models.StoryModel{
		Id: storyID,
	}
	return s.DeleteOne(ctx, d.engine)
}

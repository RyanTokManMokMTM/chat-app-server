package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IStoryRepo[T any] interface {
	CreateOne(ctx context.Context, story T) (T, error)
	FindOneByUUID(ctx context.Context, UUID string) (T, error)
	FindOneByID(ctx context.Context, ID uint) (T, error)
	FindOneUserStory(ctx context.Context, ID uint, userID uint) (T, error)
	UpdateOne(ctx context.Context, story T) (T, error)
	DeleteOne(ctx context.Context, UUID string) error
	DeleteOneByID(ctx context.Context, ID uint) error

	GetActiveStoryList(
		ctx context.Context,
		userID uint,
		pageOffset,
		pageLimit int) ([]models.StoriesWithLatestStoryTime, error)

	GetActiveStoryListByTime(
		ctx context.Context,
		userID uint,
		pageOffset,
		pageLimit int,
		timeStamp int64) ([]models.StoriesWithLatestStoryTime, error)

	CountFriendActiveStory(
		ctx context.Context,
		userID uint) (int64, error)

	CountFriendActiveStoryByTime(
		ctx context.Context,
		userID uint,
		timeStamp int64) (int64, error)

	FindAllUserStories(ctx context.Context, userID uint) ([]uint, error)
	FindAllUserStoriesByTimeStamp(ctx context.Context, userID uint, timeStamp int64) ([]T, error)
}

var _ IStoryRepo[*models.StoryModel] = (*StoryRepo)(nil)

type StoryRepo struct {
	engine *gorm.DB
	IRepository[*models.StoryModel, models.StoryModel]
}

func NewStoryRepo(engine *gorm.DB) IStoryRepo[*models.StoryModel] {
	return &StoryRepo{
		engine:      engine,
		IRepository: NewRepository[*models.StoryModel, models.StoryModel](engine),
	}
}

func (storyRepo *StoryRepo) CreateOne(ctx context.Context, story *models.StoryModel) (*models.StoryModel, error) {
	return storyRepo.Create(ctx, story)
}

func (storyRepo *StoryRepo) FindOneByUUID(ctx context.Context, UUID string) (*models.StoryModel, error) {
	return storyRepo.Find(ctx, &models.StoryModel{Base: models.Base{UUID: UUID}})
}

func (storyRepo *StoryRepo) FindOneByID(ctx context.Context, ID uint) (*models.StoryModel, error) {
	return storyRepo.Find(ctx, &models.StoryModel{Base: models.Base{ID: ID}})
}

func (storyRepo *StoryRepo) UpdateOne(ctx context.Context, story *models.StoryModel) (*models.StoryModel, error) {
	if err := storyRepo.Update(ctx, story); err != nil {
		return nil, err
	}
	return story, nil
}

func (storyRepo *StoryRepo) DeleteOne(ctx context.Context, UUID string) error {
	return storyRepo.Delete(ctx, &models.StoryModel{Base: models.Base{UUID: UUID}})
}

func (storyRepo *StoryRepo) DeleteOneByID(ctx context.Context, ID uint) error {
	return storyRepo.Delete(ctx, &models.StoryModel{Base: models.Base{ID: ID}})
}

func (storyRepo *StoryRepo) FindOneUserStory(ctx context.Context, ID uint, userID uint) (*models.StoryModel, error) {
	result := &models.StoryModel{}
	if err := storyRepo.engine.WithContext(ctx).Debug().Where("ID = ? AND user_ID = ?", ID, userID).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (storyRepo *StoryRepo) GetActiveStoryList(
	ctx context.Context,
	userID uint,
	pageOffset,
	pageLimit int) ([]models.StoriesWithLatestStoryTime, error) {
	var stories []models.StoriesWithLatestStoryTime
	if err := storyRepo.engine.
		WithContext(ctx).Debug().Model(models.StoryModel{}).Select("*", "max(created_at) as latest_time").
		Preload("UserInfo").Where("user_ID IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_ID").Where("user_ID = ?", userID)).
		Where("created_at >= NOW() - INTERVAL 1 DAY").Group("user_ID").
		Offset(pageOffset).
		Limit(pageLimit).Order("latest_time desc").
		Find(&stories).Error; err != nil {
		return nil, err
	}

	return stories, nil
}

func (storyRepo *StoryRepo) GetActiveStoryListByTime(
	ctx context.Context,
	userID uint,
	pageOffset,
	pageLimit int,
	timeStamp int64) ([]models.StoriesWithLatestStoryTime, error) {
	var stories []models.StoriesWithLatestStoryTime
	if err := storyRepo.engine.WithContext(ctx).Debug().Model(models.StoryModel{}).Select("*", "max(created_at) as latest_time").Preload("UserInfo").Where("user_ID IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_ID").Where("user_ID = ?", userID)).
		Where("created_at >= NOW() - INTERVAL 1 DAY AND created_at  <= FROM_UNIXTIME(?)", timeStamp).Group("user_ID").
		Offset(pageOffset).
		Limit(pageLimit).Order("latest_time desc").
		Find(&stories).Error; err != nil {
		return nil, err
	}

	return stories, nil
}

func (storyRepo *StoryRepo) CountFriendActiveStory(ctx context.Context, userID uint) (int64, error) {
	var count int64
	if err := storyRepo.engine.WithContext(ctx).Debug().Model(models.StoryModel{}).Where("user_ID IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_ID").Where("user_ID = ?", userID)).
		Where("created_at >= NOW() - INTERVAL 1 DAY").Group("user_ID").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (storyRepo *StoryRepo) CountFriendActiveStoryByTime(ctx context.Context, userID uint, timeStamp int64) (int64, error) {
	var count int64
	if err := storyRepo.engine.WithContext(ctx).Debug().Model(models.StoryModel{}).Where("user_ID IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_ID").Where("user_ID = ?", userID)).
		Where("created_at >= NOW() - INTERVAL 1 DAY  AND created_at <= FROM_UNIXTIME(?)", timeStamp).Group("user_ID").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (storyRepo *StoryRepo) FindAllUserStories(ctx context.Context, userID uint) ([]uint, error) {
	var IDs []uint
	err := storyRepo.engine.WithContext(ctx).Debug().
		Model(&models.StoryModel{}).
		Select("ID").
		Where("user_ID = ? AND created_at >= NOW() - INTERVAL 1 DAY", userID).Find(&IDs).Error
	if err != nil {
		return nil, err
	}
	return IDs, nil
}

func (storyRepo *StoryRepo) FindAllUserStoriesByTimeStamp(
	ctx context.Context,
	userID uint,
	timeStamp int64) ([]*models.StoryModel, error) {
	var story []*models.StoryModel
	err := storyRepo.engine.WithContext(ctx).Debug().
		Model(&models.StoryModel{}).
		Where("user_ID = ? AND created_at >= NOW() - INTERVAL 1 DAY AND  created_at <= FROM_UNIXTIME(?)", userID, timeStamp).Find(&story).Error
	if err != nil {
		return nil, err
	}
	return story, nil
}

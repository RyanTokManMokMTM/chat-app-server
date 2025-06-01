package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IStoryRepo[T any] interface {
	CreateOne(ctx context.Context, story T) (*T, error)
	FindOneByUUID(ctx context.Context, uuid string) (*T, error)
	FindOneByID(ctx context.Context, id uint) (*T, error)
	FindOneUserStory(ctx context.Context, id uint, userId uint) (*T, error)
	UpdateOne(ctx context.Context, story T) (*T, error)
	DeleteOne(ctx context.Context, uuid string) error
	DeleteOneById(ctx context.Context, id uint) error

	GetActiveStoryList(
		ctx context.Context,
		userId uint,
		pageOffset,
		pageLimit int) ([]models.StoriesWithLatestStoryTime, error)

	GetActiveStoryListByTime(
		ctx context.Context,
		userId uint,
		pageOffset,
		pageLimit int,
		timeStamp int64) ([]models.StoriesWithLatestStoryTime, error)

	CountFriendActiveStory(
		ctx context.Context,
		userId uint) (int64, error)

	CountFriendActiveStoryByTime(
		ctx context.Context,
		userId uint,
		timeStamp int64) (int64, error)

	FindAllUserStories(ctx context.Context, userId uint) ([]uint, error)
	FindAllUserStoriesByTimeStamp(ctx context.Context, userId uint, timeStamp int64) ([]*models.StoryModel, error)
}

var _ IStoryRepo[models.StoryModel] = (*StoryRepo)(nil)

type StoryRepo struct {
	engine *gorm.DB
	IRepository[*models.StoryModel, models.StoryModel]
}

func NewStoryRepo(engine *gorm.DB) IStoryRepo[models.StoryModel] {
	return &StoryRepo{
		engine:      engine,
		IRepository: NewRepository[*models.StoryModel, models.StoryModel](engine),
	}
}

func (storyRepo *StoryRepo) CreateOne(ctx context.Context, story models.StoryModel) (*models.StoryModel, error) {
	if err := storyRepo.Create(ctx, &story); err != nil {
		return nil, err
	}
	return &story, nil
}

func (storyRepo *StoryRepo) FindOneByUUID(ctx context.Context, uuid string) (*models.StoryModel, error) {
	result, err := storyRepo.Find(ctx, &models.StoryModel{Uuid: uuid})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (storyRepo *StoryRepo) FindOneByID(ctx context.Context, id uint) (*models.StoryModel, error) {
	result, err := storyRepo.Find(ctx, &models.StoryModel{Id: id})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (storyRepo *StoryRepo) UpdateOne(ctx context.Context, story models.StoryModel) (*models.StoryModel, error) {
	if err := storyRepo.Update(ctx, &story); err != nil {
		return nil, err
	}
	return &story, nil
}

func (storyRepo *StoryRepo) DeleteOne(ctx context.Context, uuid string) error {
	return storyRepo.Delete(ctx, &models.StoryModel{Uuid: uuid})
}

func (storyRepo *StoryRepo) DeleteOneById(ctx context.Context, id uint) error {
	return storyRepo.Delete(ctx, &models.StoryModel{Id: id})
}

func (storyRepo *StoryRepo) FindOneUserStory(ctx context.Context, id uint, userId uint) (*models.StoryModel, error) {
	result := &models.StoryModel{}
	if err := storyRepo.engine.WithContext(ctx).Debug().Where("id = ? AND user_id = ?", id, userId).First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (storyRepo *StoryRepo) GetActiveStoryList(
	ctx context.Context,
	userId uint,
	pageOffset,
	pageLimit int) ([]models.StoriesWithLatestStoryTime, error) {
	var stories []models.StoriesWithLatestStoryTime
	if err := storyRepo.engine.
		WithContext(ctx).Debug().Model(models.StoryModel{}).Select("*", "max(created_at) as latest_time").
		Preload("UserInfo").Where("user_id IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY").Group("user_id").
		Offset(pageOffset).
		Limit(pageLimit).Order("latest_time desc").
		Find(&stories).Error; err != nil {
		return nil, err
	}

	return stories, nil
}

func (storyRepo *StoryRepo) GetActiveStoryListByTime(
	ctx context.Context,
	userId uint,
	pageOffset,
	pageLimit int,
	timeStamp int64) ([]models.StoriesWithLatestStoryTime, error) {
	var stories []models.StoriesWithLatestStoryTime
	if err := storyRepo.engine.WithContext(ctx).Debug().Model(models.StoryModel{}).Select("*", "max(created_at) as latest_time").Preload("UserInfo").Where("user_id IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY AND created_at  <= FROM_UNIXTIME(?)", timeStamp).Group("user_id").
		Offset(pageOffset).
		Limit(pageLimit).Order("latest_time desc").
		Find(&stories).Error; err != nil {
		return nil, err
	}

	return stories, nil
}

func (storyRepo *StoryRepo) CountFriendActiveStory(ctx context.Context, userId uint) (int64, error) {
	var count int64
	if err := storyRepo.engine.WithContext(ctx).Debug().Model(models.StoryModel{}).Where("user_id IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY").Group("user_id").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (storyRepo *StoryRepo) CountFriendActiveStoryByTime(ctx context.Context, userId uint, timeStamp int64) (int64, error) {
	var count int64
	if err := storyRepo.engine.WithContext(ctx).Debug().Model(models.StoryModel{}).Where("user_id IN (?)",
		storyRepo.engine.Model(models.UserFriend{}).Select("friend_id").Where("user_id = ?", userId)).
		Where("created_at >= NOW() - INTERVAL 1 DAY  AND created_at <= FROM_UNIXTIME(?)", timeStamp).Group("user_id").
		Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

func (storyRepo *StoryRepo) FindAllUserStories(ctx context.Context, userId uint) ([]uint, error) {
	var ids []uint
	err := storyRepo.engine.WithContext(ctx).Debug().
		Model(&models.StoryModel{}).
		Select("Id").
		Where("user_id = ? AND created_at >= NOW() - INTERVAL 1 DAY", userId).Find(&ids).Error
	if err != nil {
		return nil, err
	}
	return ids, nil
}

func (storyRepo *StoryRepo) FindAllUserStoriesByTimeStamp(
	ctx context.Context,
	userId uint,
	timeStamp int64) ([]*models.StoryModel, error) {
	var story []*models.StoryModel
	err := storyRepo.engine.WithContext(ctx).Debug().
		Model(&models.StoryModel{}).
		Where("user_id = ? AND created_at >= NOW() - INTERVAL 1 DAY AND  created_at <= FROM_UNIXTIME(?)", userId, timeStamp).Find(&story).Error
	if err != nil {
		return nil, err
	}
	return story, nil
}

package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IMessageRepo[T any] interface {
	CreateOne(ctx context.Context, user *T) error
	FindOneByID(ctx context.Context, uuid string) (T, error)
	UpdateOne(ctx context.Context, data *T) error
	DeleteOne(ctx context.Context, uuid string) error

	CountMessage(
		ctx context.Context,
		messageType models.MessageType,
		from, to uint,
	) (int64, error)

	GetMessages(
		ctx context.Context,
		messageType models.MessageType,
		from, to uint,
		pageLimit int,
	) ([]*models.Message, error)

	GetMessagesByLatestId(
		ctx context.Context,
		messageType models.MessageType,
		from, to uint,
		latestId uint,
		pageLimit int,
	) ([]*models.Message, error)
}

var _ IMessageRepo[models.Message] = (*MessageRepo)(nil)

type MessageRepo struct {
	engine *gorm.DB
	IRepository[*models.Message, models.Message]
}

func NewMessageRepo(engine *gorm.DB) IMessageRepo[models.Message] {
	return &MessageRepo{
		engine:      engine,
		IRepository: NewRepository[*models.Message, models.Message](engine),
	}
}

func (messageRepo *MessageRepo) CreateOne(ctx context.Context, data *models.Message) error {
	return messageRepo.Create(ctx, data)
}

func (messageRepo *MessageRepo) FindOneByID(ctx context.Context, uuid string) (models.Message, error) {
	return messageRepo.Find(ctx, &models.Message{Uuid: uuid})
}

func (messageRepo *MessageRepo) UpdateOne(ctx context.Context, data *models.Message) error {
	return messageRepo.Update(ctx, data)
}

func (messageRepo *MessageRepo) DeleteOne(ctx context.Context, uuid string) error {
	return messageRepo.Delete(ctx, &models.Message{Uuid: uuid})
}

func (messageRepo *MessageRepo) CountMessage(
	ctx context.Context,
	messageType models.MessageType,
	from, to uint,
) (int64, error) {
	var count int64 = 0
	if err := messageRepo.engine.WithContext(ctx).Model(&models.Message{}).
		Where("message_type = ? AND (from_user_id in (?,?) or to_user_id in (?,?))", messageType, from, to, from, to).
		Debug().Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (messageRepo *MessageRepo) GetMessages(
	ctx context.Context,
	messageType models.MessageType,
	from, to uint,
	pageLimit int,
) ([]*models.Message, error) {
	var message = make([]*models.Message, 0)
	if err := messageRepo.engine.WithContext(ctx).Debug().
		Where("message_type = ? and (from_user_id in (?,?) or to_user_id in (?,?)) ", messageType, from, to, to, from).
		Limit(pageLimit).Order("created_at DESC").
		Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (messageRepo *MessageRepo) GetMessagesByLatestId(
	ctx context.Context,
	messageType models.MessageType,
	from, to uint,
	latestId uint,
	pageLimit int,
) ([]*models.Message, error) {
	var message = make([]*models.Message, 0)
	if err := messageRepo.engine.WithContext(ctx).Debug().
		Where("message_type = ? and (from_user_id in (?,?) or to_user_id in (?,?)) AND id < ?", messageType, from, to, to, from, latestId).
		Limit(pageLimit).Order("created_at DESC").
		Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

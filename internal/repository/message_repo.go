package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IMessageRepo[T any] interface {
	CreateOne(ctx context.Context, msg T) (*T, error)
	FindOneByID(ctx context.Context, id uint) (*T, error)
	FindOneByUUID(ctx context.Context, uuid string) (*T, error)
	UpdateOne(ctx context.Context, msg T) error
	DeleteOneByUUID(ctx context.Context, uuid string) error
	DeleteOneByID(ctx context.Context, id uint) error

	CountMessage(
		ctx context.Context,
		messageType models.MessageType,
		from, to uint,
	) (int64, error)

	FindMessages(
		ctx context.Context,
		from, to uint,
		messageType models.MessageType,
		pageLimit int,
		latestId uint) ([]*models.Message, error)

	// GetMessages(
	// 	ctx context.Context,
	// 	messageType models.MessageType,
	// 	from, to uint,
	// 	pageLimit int,
	// ) ([]*models.Message, error)

	// GetMessagesByLatestId(
	// 	ctx context.Context,
	// 	messageType models.MessageType,
	// 	from, to uint,
	// 	latestId uint,
	// 	pageLimit int,
	// ) ([]*models.Message, error)
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

func (messageRepo *MessageRepo) CreateOne(ctx context.Context, msg models.Message) (*models.Message, error) {
	if err := messageRepo.Create(ctx, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}

func (messageRepo *MessageRepo) FindOneByID(ctx context.Context, id uint) (*models.Message, error) {
	result, err := messageRepo.Find(ctx, &models.Message{ID: id})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (messageRepo *MessageRepo) FindOneByUUID(ctx context.Context, uuid string) (*models.Message, error) {
	result, err := messageRepo.Find(ctx, &models.Message{Uuid: uuid})
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (messageRepo *MessageRepo) UpdateOne(ctx context.Context, msg models.Message) error {
	return messageRepo.Update(ctx, &msg)
}

func (messageRepo *MessageRepo) DeleteOneByUUID(ctx context.Context, uuid string) error {
	return messageRepo.Delete(ctx, &models.Message{Uuid: uuid})
}

func (messageRepo *MessageRepo) DeleteOneByID(ctx context.Context, id uint) error {
	return messageRepo.Delete(ctx, &models.Message{ID: id})
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

func (messageRepo *MessageRepo) getMessages(
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

func (messageRepo *MessageRepo) getMessagesByLatestId(
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

func (messageRepo *MessageRepo) FindMessages(
	ctx context.Context,
	from, to uint,
	messageType models.MessageType,
	pageLimit int,
	latestId uint) ([]*models.Message, error) {

	if latestId <= 0 {
		return messageRepo.getMessages(ctx, messageType, from, to, pageLimit)
	}
	return messageRepo.getMessagesByLatestId(ctx, messageType, from, to, latestId, pageLimit)
}

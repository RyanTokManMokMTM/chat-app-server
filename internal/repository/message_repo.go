package repository

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type IMessageRepo[T any] interface {
	CreateOne(ctx context.Context, msg T) (T, error)
	FindOneByID(ctx context.Context, ID uint) (T, error)
	FindOneByUUID(ctx context.Context, UUID string) (T, error)
	UpdateOne(ctx context.Context, msg T) error
	DeleteOneByUUID(ctx context.Context, UUID string) error
	DeleteOneByID(ctx context.Context, ID uint) error

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
		latestID uint) ([]T, error)

	// GetMessages(
	// 	ctx context.Context,
	// 	messageType models.MessageType,
	// 	from, to uint,
	// 	pageLimit int,
	// ) ([]*models.Message, error)

	// GetMessagesByLatestID(
	// 	ctx context.Context,
	// 	messageType models.MessageType,
	// 	from, to uint,
	// 	latestID uint,
	// 	pageLimit int,
	// ) ([]*models.Message, error)
}

var _ IMessageRepo[*models.Message] = (*messageRepo)(nil)

type messageRepo struct {
	engine *gorm.DB
	IRepository[*models.Message, models.Message]
}

func NewMessageRepo(engine *gorm.DB) IMessageRepo[*models.Message] {
	return &messageRepo{
		engine:      engine,
		IRepository: NewRepository[*models.Message, models.Message](engine),
	}
}

func (messageRepo *messageRepo) CreateOne(ctx context.Context, msg *models.Message) (*models.Message, error) {
	msg, err := messageRepo.Create(ctx, msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (messageRepo *messageRepo) FindOneByID(ctx context.Context, ID uint) (*models.Message, error) {
	result, err := messageRepo.Find(ctx, &models.Message{Base: models.Base{ID: ID}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (messageRepo *messageRepo) FindOneByUUID(ctx context.Context, UUID string) (*models.Message, error) {
	result, err := messageRepo.Find(ctx, &models.Message{Base: models.Base{UUID: UUID}})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (messageRepo *messageRepo) UpdateOne(ctx context.Context, msg *models.Message) error {
	return messageRepo.Update(ctx, msg)
}

func (messageRepo *messageRepo) DeleteOneByUUID(ctx context.Context, UUID string) error {
	return messageRepo.Delete(ctx, &models.Message{Base: models.Base{UUID: UUID}})
}

func (messageRepo *messageRepo) DeleteOneByID(ctx context.Context, ID uint) error {
	return messageRepo.Delete(ctx, &models.Message{Base: models.Base{ID: ID}})
}

func (messageRepo *messageRepo) CountMessage(
	ctx context.Context,
	messageType models.MessageType,
	from, to uint,
) (int64, error) {
	var count int64 = 0
	if err := messageRepo.engine.WithContext(ctx).Model(&models.Message{}).
		Where("message_type = ? AND (from_user_ID in (?,?) or to_user_ID in (?,?))", messageType, from, to, from, to).
		Debug().Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (messageRepo *messageRepo) getMessages(
	ctx context.Context,
	messageType models.MessageType,
	from, to uint,
	pageLimit int,
) ([]*models.Message, error) {
	var message = make([]*models.Message, 0)
	if err := messageRepo.engine.WithContext(ctx).Debug().
		Where("message_type = ? and (from_user_ID in (?,?) or to_user_ID in (?,?)) ", messageType, from, to, to, from).
		Limit(pageLimit).Order("created_at DESC").
		Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (messageRepo *messageRepo) getMessagesByLatestID(
	ctx context.Context,
	messageType models.MessageType,
	from, to uint,
	latestID uint,
	pageLimit int,
) ([]*models.Message, error) {
	var message = make([]*models.Message, 0)
	if err := messageRepo.engine.WithContext(ctx).Debug().
		Where("message_type = ? and (from_user_ID in (?,?) or to_user_ID in (?,?)) AND ID < ?", messageType, from, to, to, from, latestID).
		Limit(pageLimit).Order("created_at DESC").
		Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (messageRepo *messageRepo) FindMessages(
	ctx context.Context,
	from, to uint,
	messageType models.MessageType,
	pageLimit int,
	latestID uint) ([]*models.Message, error) {

	if latestID <= 0 {
		return messageRepo.getMessages(ctx, messageType, from, to, pageLimit)
	}
	return messageRepo.getMessagesByLatestID(ctx, messageType, from, to, latestID, pageLimit)
}

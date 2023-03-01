package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

func (d *DAO) InsertOneMessage(ctx context.Context, content string, from, to, messageType, contentType uint) error {
	msg := &models.Message{
		FromUserID:  from,
		ToUserID:    to,
		Content:     content,
		MessageType: messageType,
		ContentType: contentType,
	}

	return msg.InsertOne(ctx, d.engine)
}

func (d *DAO) FindOneMessage(ctx context.Context, messageID uint) (*models.Message, error) {
	msg := &models.Message{
		ID: messageID,
	}

	if err := msg.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}

	return msg, nil
}
func (d *DAO) DeleteOneMessage(ctx context.Context, messageID uint) error {
	msg := &models.Message{
		ID: messageID,
	}
	return msg.DeleteOne(ctx, d.engine)
}

func (d *DAO) GetMessage(ctx context.Context, from, to, messageType uint) ([]*models.Message, error) {
	msg := &models.Message{
		FromUserID:  from,
		ToUserID:    to,
		MessageType: messageType,
	}

	return msg.GetMessages(ctx, d.engine)
}

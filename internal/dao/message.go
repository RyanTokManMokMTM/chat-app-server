package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

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

func (d *DAO) InsertOneMessage(ctx context.Context, message *socket_message.Message) {
	var msg *models.Message
	if message.MessageType == variable.MESSAGE_TYPE_USERCHAT {
		msg = insertUserMessage(ctx, message, d.engine)
	} else if message.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
		msg = insertOneGroupMessage(ctx, message, d.engine)
	}
	if msg == nil {
		logx.Error("message is null...")
		return
	}
	if err := msg.InsertOne(ctx, d.engine); err != nil {
		logx.Error("insert message error %v", err.Error())
	}
}

func insertUserMessage(ctx context.Context, message *socket_message.Message, engine *gorm.DB) *models.Message {
	fromUser := &models.UserModel{
		Uuid: message.FromUUID,
	}
	if err := fromUser.FindOneUserByUUID(engine, ctx); err != nil {
		logx.Errorf("find from user by UUID err : %v", err.Error())
		return nil
	}

	toUser := &models.UserModel{
		Uuid: message.ToUUID,
	}
	if err := toUser.FindOneUserByUUID(engine, ctx); err != nil {
		logx.Errorf("find user by UUID err : %v", err.Error())
		return nil
	}

	return &models.Message{
		Uuid:        message.MessageID,
		FromUserID:  fromUser.ID,
		ToUserID:    toUser.ID,
		Content:     message.Content,
		MessageType: uint(message.MessageType),
		ContentType: uint(message.ContentType),
		URL:         message.UrlPath,
	}

}

func insertOneGroupMessage(ctx context.Context, message *socket_message.Message, engine *gorm.DB) *models.Message {
	fromUser := models.UserModel{
		Uuid: message.FromUUID,
	}
	if err := fromUser.FindOneUserByUUID(engine, ctx); err != nil {
		logx.Errorf("find from user by UUID err : %v", err.Error())
		return nil
	}

	groupInfo := models.Group{
		Uuid: message.ToUUID,
	}
	if err := groupInfo.FindOneByUUID(ctx, engine); err != nil {
		logx.Error("find one group error : %v ", err.Error())
		return nil
	}

	return &models.Message{
		Uuid:        message.MessageID,
		FromUserID:  fromUser.ID,
		ToUserID:    groupInfo.ID,
		Content:     message.Content,
		ContentType: uint(message.ContentType),
		MessageType: uint(message.MessageType),
	}

}

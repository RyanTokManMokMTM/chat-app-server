package dao

import (
	"api/app/internal/models"
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
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

func (d *DAO) GetMessage(ctx context.Context, from, to, messageType uint, pageLimit int, latestId uint) ([]*models.Message, error) {
	msg := &models.Message{
		ID:          latestId,
		FromUserID:  from,
		ToUserID:    to,
		MessageType: messageType,
	}

	if latestId <= 0 {
		return msg.GetMessages(ctx, d.engine, pageLimit)
	}
	return msg.GetMessagesByLatestId(ctx, d.engine, pageLimit)

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
	fromUser := &models.User{
		Uuid: message.FromUUID,
	}
	if err := fromUser.FindOneUserByUUID(engine, ctx); err != nil {
		logx.Errorf("find from user by UUID err : %v", err.Error())
		return nil
	}

	toUser := &models.User{
		Uuid: message.ToUUID,
	}
	if err := toUser.FindOneUserByUUID(engine, ctx); err != nil {
		logx.Errorf("find user by UUID err : %v", err.Error())
		return nil
	}

	return &models.Message{
		Uuid:                 message.MessageID,
		FromUserID:           fromUser.Id,
		ToUserID:             toUser.Id,
		Content:              message.Content,
		MessageType:          uint(message.MessageType),
		ContentType:          message.ContentType,
		Url:                  message.UrlPath,
		FileName:             message.FileName,
		FileSize:             uint(message.FileSize),
		ContentAvailableTime: uint(message.ContentAvailableTime),
		ContentId:            message.ContentUUID,
		ContentUserName:      message.ContentUserName,
		ContentUserAvatar:    message.ContentUserAvatar,
		ContentUserUUID:      message.ContentUserUUID,
	}

}
func (d *DAO) CountMessage(ctx context.Context, messageType, id uint) (int64, error) {
	m := &models.Message{
		FromUserID:  id,
		ToUserID:    id,
		MessageType: messageType,
	}

	return m.CountMessage(ctx, d.engine)
}

func insertOneGroupMessage(ctx context.Context, message *socket_message.Message, engine *gorm.DB) *models.Message {
	fromUser := models.User{
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
		Uuid:                 message.MessageID,
		FromUserID:           fromUser.Id,
		ToUserID:             groupInfo.Id,
		Content:              message.Content,
		ContentType:          message.ContentType,
		MessageType:          uint(message.MessageType),
		Url:                  message.UrlPath,
		FileName:             message.FileName,
		FileSize:             uint(message.FileSize),
		ContentAvailableTime: uint(message.ContentAvailableTime),
		ContentId:            message.ContentUUID,
		ContentUserName:      message.ContentUserName,
		ContentUserAvatar:    message.ContentUserAvatar,
		ContentUserUUID:      message.ContentUserUUID,
	}

}

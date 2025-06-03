package helper

import (
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/socket-proto"
)

// FIXME: should be use UUId instead
func ConvertSocketMessageToMessage(fromId, toId uint, message *socket_message.Message) *models.Message {
	if message == nil {
		return nil
	}
	return &models.Message{
		UUId:                 message.MessageId,
		FromUserId:           fromId,
		ToUserId:             toId,
		Content:              message.Content,
		MessageType:          uint(message.MessageType),
		ContentType:          message.ContentType,
		Url:                  message.UrlPath,
		FileName:             message.FileName,
		FileSize:             uint(message.FileSize),
		ContentAvailableTime: uint(message.ContentAvailableTime),
		ContentId:            message.ContentUUId,
		ContentUserName:      message.ContentUserName,
		ContentUserAvatar:    message.ContentUserAvatar,
		ContentUserUUId:      message.ContentUserUUId,
	}
}

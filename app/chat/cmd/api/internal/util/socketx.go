package util

import (
	"github.com/google/uuid"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/variable"
	socket_message "github.com/ryantokmanmokmtm/chat-app-server/app/internal/socket-proto"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

func MarshalNotificationContent(FromUUID, groupUUID, content string) ([]byte, error) {
	msg := &socket_message.Message{
		MessageID:   uuid.New().String(),
		FromUUID:    FromUUID,
		ToUUID:      groupUUID,
		ContentType: variable.SYS,
		Content:     content,
		MessageType: variable.MESSAGE_TYPE_GROUPCHAT,
		EventType:   variable.MESSAGE,
	}

	messageBytes, err := jsonx.Marshal(msg)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	return messageBytes, nil
}
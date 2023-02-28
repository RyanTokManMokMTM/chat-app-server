package models

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/common/variable"
	"gorm.io/gorm"
)

type Message struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;type:int"`
	UUID string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`

	/*
		MessageType: 1 -> single
		From A User to Other User
		example:
			UUID: abc -> UUID:edf
			UUID: edf -> UUID:abc....
		MessageType: 2 -> group
		But it different to group chat
		ToUserID will always be the group UUID
		example:

			UUID - group member(abc) -> group's UUID(aaa)
			UUID - group member(efd) -> group's UUID(aaa)
			in group's with UUID aaa have 2 message which send from UUID abc and UUID efd
	*/
	FromUserID  uint   `gorm:"index;comment:'sender userID'"`
	ToUserID    uint   `gorm:"index;comment:'receiver userID'"`
	Content     string `gorm:"comment:'message content'"`
	MessageType uint   `gorm:"comment;'sent message type: 1:single ,2: group'"`
	ContentType uint   `gorm:"comment:'content type: 1: text, 2: file,3:audio,4:video'"`
	CommonField
}

func (m *Message) TableName() string {
	return "messages"
}

func (m *Message) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&m).Error
}

func (m *Message) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().First(&m).Error
}

func (m *Message) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Delete(&m).Error
}

func (m *Message) GetMessages(ctx context.Context, db *gorm.DB) ([]*Message, error) {
	var message []*Message = make([]*Message, 0)
	if m.ContentType == variable.MESSAGE_TYPE_USERCHAT {
		//TODO: User chat will exactly have 2 user in the chat, so from -> to || to -> form are mean the same chat room
		if err := db.WithContext(ctx).Debug().Where("message_type = ? and (from_user_id in (?,?) or to_user_id in (?,?))", m.MessageType, m.FromUserID, m.ToUserID, m.ToUserID, m.FromUserID).Find(&message).Error; err != nil {
			return nil, err
		}
	} else if m.ContentType == variable.MESSAGE_TYPE_GROUPCHAT {
		//TODO: for group chat message
		//TODO: GROUP ID is to user id
		if err := db.WithContext(ctx).Debug().Where("message_type = ? AND to_user_id = ?", m.MessageType, m.ToUserID).Find(&message).Error; err != nil {
			return nil, err
		}
	}

	return message, nil
}

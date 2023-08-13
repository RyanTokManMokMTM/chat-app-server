package models

import (
	"context"
	"gorm.io/gorm"
)

type Message struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;type:int"`
	Uuid string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`

	/*
		MessageType: 1 -> single
		From A User to Other User
		example:
			Uuid: abc -> Uuid:edf
			Uuid: edf -> Uuid:abc....
		MessageType: 2 -> group
		But it different to group chat
		ToUserID will always be the group Uuid
		example:

			Uuid - group member(abc) -> group's Uuid(aaa)
			Uuid - group member(efd) -> group's Uuid(aaa)
			in group's with Uuid aaa have 2 message which send from Uuid abc and Uuid efd
	*/
	FromUserID  uint   `gorm:"index;comment:'sender userID'"`
	ToUserID    uint   `gorm:"index;comment:'receiver userID'"`
	Content     string `gorm:"comment:'message content'"`
	MessageType uint   `gorm:"comment;'sent message type: 1:single ,2: group'"`
	ContentType uint   `gorm:"comment:'content type: 1: text, 2: file,3:audio,4:video'"`
	Url         string `gorm:"comment:'image url path'"`
	FileName    string `gorm:"comment:'file name'"`
	FileSize    uint   `gorm:"comment:'file size'"`
	StoryTime   uint   `gorm:"comment:'reply story available time'"`
	CommonField
}

//func (m *Message) BeforeCreate(tx *gorm.DB) error {
//	m.Uuid = uuid.New().String()
//	return nil
//}

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

func (m *Message) CountMessage(ctx context.Context, db *gorm.DB) (int64, error) {
	var count int64 = 0
	if err := db.WithContext(ctx).Model(&m).
		Where("message_type = ? AND (from_user_id in (?,?) or to_user_id in (?,?))", m.MessageType, m.FromUserID, m.ToUserID, m.FromUserID, m.ToUserID).
		Debug().Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (m *Message) GetMessages(ctx context.Context, db *gorm.DB, pageOffset, pageLimit int) ([]*Message, error) {
	var message = make([]*Message, 0)
	//if m.ContentType == variable.MESSAGE_TYPE_USERCHAT {
	//	//TODO: User chat will exactly have 2 user in the chat, so from -> to || to -> form are mean the same chat room
	//	if err := db.WithContext(ctx).Debug().
	//		Where("message_type = ? and (from_user_id in (?,?) or to_user_id in (?,?))", m.MessageType, m.FromUserID, m.ToUserID, m.ToUserID, m.FromUserID).
	//		Offset(pageOffset).
	//		Limit(pageLimit).
	//		Find(&message).Error; err != nil {
	//		return nil, err
	//	}
	//} else if m.ContentType == variable.MESSAGE_TYPE_GROUPCHAT {
	//	//TODO: for group chat message
	//	//TODO: GROUP Id is to user id
	//	if err := db.WithContext(ctx).Debug().
	//		Where("message_type = ? AND to_user_id = ?", m.MessageType, m.ToUserID).
	//		Offset(pageOffset).
	//		Limit(pageLimit).Find(&message).Error; err != nil {
	//		return nil, err
	//	}
	//}

	/*
		SINGLE CHAT:
			MESSAGE TYPE = 1
			TO USER = ID OR FROM USER =ID

		GROUP CHAT
			MESSAGE TYPE = 2
			TO USER = GROUP ID OR FROM USER ID = GROUP ID

	*/
	if err := db.WithContext(ctx).Debug().
		Where("message_type = ? and (from_user_id in (?,?) or to_user_id in (?,?))", m.MessageType, m.FromUserID, m.ToUserID, m.ToUserID, m.FromUserID).
		Offset(pageOffset).
		Limit(pageLimit).Order("created_at DESC").
		Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

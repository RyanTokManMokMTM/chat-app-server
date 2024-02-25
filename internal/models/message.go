package models

import (
	"context"
	"gorm.io/gorm"
)

type Message struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;types:int"`
	Uuid string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`

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
	FromUserID           uint   `gorm:"index;comment:'sender userID'"`
	ToUserID             uint   `gorm:"index;comment:'receiver userID'"`
	Content              string `gorm:"comment:'message content'"`
	MessageType          uint   `gorm:"comment;'sent message types: 1:single ,2: group'"`
	ContentType          uint   `gorm:"comment:'content types: 1: text, 2: file,3:audio,4:video'"`
	Url                  string `gorm:"comment:'image url path'"`
	FileName             string `gorm:"comment:'file name'"`
	FileSize             uint   `gorm:"comment:'file size'"`
	ContentAvailableTime uint   `gorm:"comment:'reply content available time'"`
	ContentId            string `gorm:"comment:'reply content id'"`
	ContentUserName      string `gorm:"comment:'reply content belong to which user name'"`
	ContentUserAvatar    string `gorm:"comment:'reply content belong to which user avatar'"`
	ContentUserUUID      string `gorm:"comment:'reply content belong to which user uuid'"`
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

func (m *Message) GetMessages(ctx context.Context, db *gorm.DB, pageLimit int) ([]*Message, error) {
	var message = make([]*Message, 0)
	if err := db.WithContext(ctx).Debug().
		Where("message_type = ? and (from_user_id in (?,?) or to_user_id in (?,?)) ", m.MessageType, m.FromUserID, m.ToUserID, m.ToUserID, m.FromUserID).
		Limit(pageLimit).Order("created_at DESC").
		Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

func (m *Message) GetMessagesByLatestId(ctx context.Context, db *gorm.DB, pageLimit int) ([]*Message, error) {
	var message = make([]*Message, 0)

	if err := db.WithContext(ctx).Debug().
		Where("message_type = ? and (from_user_id in (?,?) or to_user_id in (?,?)) AND id < ?", m.MessageType, m.FromUserID, m.ToUserID, m.ToUserID, m.FromUserID, m.ID).
		Limit(pageLimit).Order("created_at DESC").
		Find(&message).Error; err != nil {
		return nil, err
	}
	return message, nil
}

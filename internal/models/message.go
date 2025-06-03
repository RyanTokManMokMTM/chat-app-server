package models

type CreateMessageDTO struct {
	FromUserId           uint
	ToUserId             uint
	Content              string
	MessageType          uint
	ContentType          string
	Url                  string
	FileName             string
	FileSize             uint
	ContentAvailableTime uint
	ContentId            string
	ContentUserName      string
	ContentUserAvatar    string
	ContentUserUUId      string
}

type UpdateMessageDTO struct {
	Content              string
	ContentAvailableTime uint
}

const messageTableColName = "messages"

type MessageType uint

const (
	MessageTypeSingle MessageType = iota + 1
	MessageTypeGroup
)

type Message struct {
	/*
		MessageType: 1 -> single
		From A User to Other User
		example:
			UUID: abc -> UUID:edf
			UUID: edf -> UUID:abc....
		MessageType: 2 -> group
		But it different to group chat
		ToUserId will always be the group UUID
		example:

			UUID - group member(abc) -> group's UUID(aaa)
			UUID - group member(efd) -> group's UUID(aaa)
			in group's with UUID aaa have 2 message which send from UUID abc and UUID efd
	*/
	Base
	FromUserID           uint   `gorm:"index;comment:'sender UserId'"`
	ToUserID             uint   `gorm:"index;comment:'receiver UserId'"`
	Content              string `gorm:"comment:'message content'"`
	MessageType          uint   `gorm:"comment;'sent message types: 1:single ,2: group'"`
	ContentType          string `gorm:"comment:'content types : text,image,audio..."`
	Url                  string `gorm:"comment:'image url path'"`
	FileName             string `gorm:"comment:'file name'"`
	FileSize             uint   `gorm:"comment:'file size'"`
	ContentAvailableTime uint   `gorm:"comment:'reply content available time'"`
	ContentID            string `gorm:"comment:'reply content id'"`
	ContentUserName      string `gorm:"comment:'reply content belong to which user name'"`
	ContentUserAvatar    string `gorm:"comment:'reply content belong to which user avatar'"`
	ContentUserUUID      string `gorm:"comment:'reply content belong to which user uuid'"`
}

func (m *Message) TableName() string {
	return messageTableColName
}

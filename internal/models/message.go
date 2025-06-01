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
	ContentUserUUID      string
}

type UpdateMessageDTO struct {
	Content              string
	ContentAvailableTime uint
}

type MessageType uint

const (
	MessageTypeSingle MessageType = iota + 1
	MessageTypeGroup
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
		ToUserId will always be the group Uuid
		example:

			Uuid - group member(abc) -> group's Uuid(aaa)
			Uuid - group member(efd) -> group's Uuid(aaa)
			in group's with Uuid aaa have 2 message which send from Uuid abc and Uuid efd
	*/
	FromUserId           uint   `gorm:"index;comment:'sender UserId'"`
	ToUserId             uint   `gorm:"index;comment:'receiver UserId'"`
	Content              string `gorm:"comment:'message content'"`
	MessageType          uint   `gorm:"comment;'sent message types: 1:single ,2: group'"`
	ContentType          string `gorm:"comment:'content types : text,image,audio..."`
	Url                  string `gorm:"comment:'image url path'"`
	FileName             string `gorm:"comment:'file name'"`
	FileSize             uint   `gorm:"comment:'file size'"`
	ContentAvailableTime uint   `gorm:"comment:'reply content available time'"`
	ContentId            string `gorm:"comment:'reply content id'"`
	ContentUserName      string `gorm:"comment:'reply content belong to which user name'"`
	ContentUserAvatar    string `gorm:"comment:'reply content belong to which user avatar'"`
	ContentUserUUID      string `gorm:"comment:'reply content belong to which user uuid'"`
	Base
}

//func (m *Message) BeforeCreate(tx *gorm.DB) error {
//	m.Uuid = uuid.New().String()
//	return nil
//}

func (m *Message) TableName() string {
	return "messages"
}

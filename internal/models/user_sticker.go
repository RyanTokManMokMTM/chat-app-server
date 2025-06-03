package models

const userStickerTableName = "users_stickers"

type UserSticker struct {
	Base
	UserID    uint `gorm:"not uniqueIndex;index;comment:'belong to which user ID'"`
	StickerID uint `gorm:"not uniqueIndex;index;comment:'belong to which sticker ID'"`

	UserInfo    User    `gorm:"foreignKey:UserID"`
	StickerInfo Sticker `gorm:"foreignKey:StickerID"`
}

func (us *UserSticker) TableName() string {
	return userStickerTableName
}

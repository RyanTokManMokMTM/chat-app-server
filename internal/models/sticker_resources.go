package models

type CreateStickerResourceDTO struct {
	StickerId uint
	Path      string
}

type UpdateStickerResourceDTO struct {
	Path string
}

const stickerResourceTableName = "sticker_resources"

type StickerResource struct {
	Base
	StickerID uint   `gorm:"not null"`
	UUID      string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	Path      string `gorm:"not null"`

	Sticker Sticker `gorm:"foreignKey:StickerID;"`
}

func (sks *StickerResource) TableName() string {
	return stickerResourceTableName
}

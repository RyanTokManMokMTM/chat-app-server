package models

type CreateStickerDTO struct {
	StickerName string
	StickerThum string
}

type UpdateStickerDTO struct {
	StickerName string
	StickerThum string
}

const stickerTableName = "sticker"

type Sticker struct {
	Base
	StickerName string `gorm:"not null"`
	StickerThum string

	Resources []StickerResource `gorm:"foreignKey:StickerID;"`
}

func (sk *Sticker) TableName() string {
	return stickerTableName
}

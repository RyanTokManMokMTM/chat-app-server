package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateStickerDTO struct {
	StickerName string
	StickerThum string
}

type UpdateStickerDTO struct {
	StickerName string
	StickerThum string
}

type Sticker struct {
	Base
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Uuid        string `gorm:"types:varchar(64);not null;index:unique"`
	StickerName string `gorm:"not null"`
	StickerThum string

	Resources []StickerResource `gorm:"foreignKey:StickerId;"`
}

func (sk *Sticker) TableName() string {
	return "sticker"
}

func (sk *Sticker) BeforeCreate(tx *gorm.DB) error {
	sk.Uuid = uuid.New().String()
	return nil
}

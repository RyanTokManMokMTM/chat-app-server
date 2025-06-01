package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateStickerResourceDTO struct {
	StickerId uint
	Path      string
}

type UpdateStickerResourceDTO struct {
	Path string
}

type StickerResource struct {
	Base
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	StickerId uint   `gorm:"not null"`
	Uuid      string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	Path      string `gorm:"not null"`

	Sticker Sticker `gorm:"foreignKey:StickerId;"`
}

func (sks *StickerResource) TableName() string {
	return "sticker_resources"
}

func (sks *StickerResource) BeforeCreate(tx *gorm.DB) error {
	sks.Uuid = uuid.New().String()
	return nil
}

package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StickerResource struct {
	Id        uint   `gorm:"primaryKey;autoIncrement"`
	StickerId uint   `gorm:"not null"`
	Uuid      string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	Path      string `gorm:"not null"`

	Sticker Sticker `gorm:"foreignKey:StickerId;"`
	CommonField
}

func (sks *StickerResource) TableName() string {
	return "sticker_resources"
}

func (sks *StickerResource) BeforeCreate(tx *gorm.DB) error {
	sks.Uuid = uuid.New().String()
	return nil
}

func (sks *StickerResource) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(sks).Error
}

func (sks *StickerResource) FindOneById(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("id = ?", sks.Id).Find(sks).Error
}

func (sks *StickerResource) FindOneByUuid(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("uuid = ?", sks.Id).Find(sks).Error
}

func (sks *StickerResource) UpdateOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&sks).Updates(sks).Error
}

func (sks *StickerResource) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&sks).Delete(sks).Error
}

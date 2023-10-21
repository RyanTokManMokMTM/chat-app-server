package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type Sticker struct {
	Id         uint   `gorm:"primaryKey;autoIncrement"`
	Uuid       string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`
	SickerName string `gorm:"not null"`

	Resources []StickerResource `gorm:"foreignKey:StickerId;"`
	CommonField
}

func (sk *Sticker) TableName() string {
	return "sticker"
}

func (sk *Sticker) BeforeCreate(tx *gorm.DB) error {
	sk.Uuid = uuid.New().String()
	return nil
}

func (sk *Sticker) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&sk).Error
}

func (sk *Sticker) FindOneById(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Preload("Resources").Where("id = ?", sk.Id).Find(sk).Error
}

func (sk *Sticker) FindOneByUuid(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Preload("Resources").Where("uuid = ?", sk.Uuid).Find(sk).Error
}

func (sk *Sticker) UpdateOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&sk).Updates(sk).Error
}

func (sk *Sticker) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&sk).Delete(sk).Error
}

func (sk *Sticker) InsertResources(ctx context.Context, db *gorm.DB, paths []string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		for _, path := range paths {
			logx.Infof("Insert path %s into db", path)
			resources := &StickerResource{
				Path:    path,
				Sticker: *sk,
			}
			//Create the resources
			if err := db.WithContext(ctx).Create(resources).Error; err != nil {
				return err
			}

			if err := db.WithContext(ctx).Debug().Model(sk).Association("Resources").Append(resources); err != nil {
				return err
			}
		}
		return nil
	})

}

package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sticker struct {
	Id          uint   `gorm:"primaryKey;autoIncrement"`
	Uuid        string `gorm:"type:varchar(64);not null;index:unique"`
	StickerName string `gorm:"not null"`
	StickerThum string

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
	return db.WithContext(ctx).Debug().Preload("Resources").Where("id = ?", sk.Id).First(sk).Error
}

func (sk *Sticker) FindOneByUuid(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Preload("Resources").Where("uuid = ?", sk.Uuid).First(sk).Error
}

func (sk *Sticker) UpdateOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&sk).Updates(sk).Error
}

func (sk *Sticker) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Model(&sk).Delete(sk).Error
}

func (sk *Sticker) FindAll(ctx context.Context, db *gorm.DB) ([]*Sticker, error) {
	resp := make([]*Sticker, 0)
	if err := db.WithContext(ctx).Debug().Model(&resp).Find(&resp).Error; err != nil {
		return nil, err
	}
	return resp, nil
}

func (sk *Sticker) InsertResources(ctx context.Context, db *gorm.DB, paths []string) error {
	return db.Transaction(func(tx *gorm.DB) error {
		resources := make([]*StickerResource, 0)
		for _, path := range paths {
			resource := &StickerResource{
				Path:    path,
				Sticker: *sk,
			}
			//Create the resources
			if err := db.WithContext(ctx).Create(resource).Error; err != nil {
				return err
			}

			resources = append(resources, resource)
		}

		return db.WithContext(ctx).Debug().Model(sk).Association("Resources").Append(resources)
	})

}

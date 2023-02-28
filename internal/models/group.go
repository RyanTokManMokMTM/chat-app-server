package models

import (
	"context"
	"gorm.io/gorm"
)

type Group struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Uuid        string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`
	GroupName   string
	GroupAvatar string
	GroupLead   uint
	CommonField
}

func (g *Group) TableName() string {
	return "groups_info"
}

func (g *Group) InsertOne(ctx context.Context, db gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&g).Error
}
func (g *Group) FindOne(ctx context.Context, db gorm.DB) error {
	return db.WithContext(ctx).Debug().First(&g).Error
}
func (g *Group) DeleteOne(ctx context.Context, db gorm.DB) error {
	return db.WithContext(ctx).Debug().Delete(&g).Error
}
func (g *Group) UpdateOne(ctx context.Context, db gorm.DB) error {
	return db.WithContext(ctx).Debug().Updates(&g).Error
}

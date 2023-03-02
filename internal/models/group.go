package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Group struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	Uuid        string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`
	GroupName   string `gorm:"type:varchar(64);not null"`
	GroupAvatar string
	GroupLead   uint `gorm:"not null;index"`

	LeadInfo UserModel `gorm:"foreignKey:GroupLead;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CommonField
}

func (g *Group) BeforeCreate(tx *gorm.DB) error {
	g.Uuid = uuid.New().String()
	return nil
}

func (g *Group) TableName() string {
	return "groups_info"
}

func (g *Group) InsertOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Create(&g).Error
}

func (g *Group) FindOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().First(&g).Error
}

func (g *Group) DeleteOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Where("group_id = ?", g.ID).Delete(&g).Error
}

func (g *Group) UpdateOne(ctx context.Context, db *gorm.DB) error {
	return db.WithContext(ctx).Debug().Updates(&g).Error
}

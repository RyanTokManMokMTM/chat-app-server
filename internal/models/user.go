package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID       uint      `gorm:"primaryKey"`
	UUID     uuid.UUID `gorm:"index"`
	Name     string    `gorm:"type:varchar(32)"`
	Email    string    `gorm:"type:varchar(64)"`
	Password string    `gorm:"type:varchar(64)"`
	CommonField
}

// Hook
func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	uniqueID := uuid.New()
	tx.Statement.SetColumn("ID", uniqueID)
	return nil
}

func (u *UserModel) TableName() string {
	return "user_info"
}

func (u *UserModel) FindOneUser(db *gorm.DB, ctx context.Context) error {
	err := db.Debug().WithContext(ctx).Find(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserModel) InsertOneUser(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Create(u).Error
}

func (u *UserModel) DeleteOneUser() error {
	return nil
}

func (u *UserModel) UpdateOneUser() error {
	return nil
}

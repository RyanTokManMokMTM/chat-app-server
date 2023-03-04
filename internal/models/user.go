package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;not null"`
	Uuid     string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`
	NickName string `gorm:"type:varchar(32)"`
	Email    string `gorm:"type:varchar(64)"`
	Password string `gorm:"type:varchar(64)"`
	Avatar   string `gorm:"type:varchar(64);null;comment:'user avatar'"`
	CommonField
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	u.Uuid = uuid.New().String()
	u.Avatar = "/default.jpg"
	return nil
}

func (u *UserModel) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (u *UserModel) TableName() string {
	return "users_info"
}

func (u *UserModel) FindOneUserByID(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).First(&u).Error
}

func (u *UserModel) FindOneUserByUUID(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Where("uuid = ?", u.Uuid).First(&u).Error
}

func (u *UserModel) FindOneUserByEmail(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Where("email = ?", u.Email).Find(&u).Error
}

func (u *UserModel) InsertOneUser(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Create(u).Error
}

func (u *UserModel) DeleteOneUser() error {
	return nil
}

func (u *UserModel) UpdateOneUser(db *gorm.DB, ctx context.Context, name string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.ID).Update("NickName", name).Error
}

func (u *UserModel) UpdateOneUserAvatar(db *gorm.DB, ctx context.Context, avatarName string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.ID).Update("Avatar", avatarName).Error
}

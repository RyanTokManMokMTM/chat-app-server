package models

import (
	"context"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	ID            uint   `gorm:"primaryKey;autoIncrement;not null"`
	Uuid          string `gorm:"type:varchar(64);not null;unique_index:idx_uuid"`
	NickName      string `gorm:"type:varchar(32)"`
	Email         string `gorm:"type:varchar(64)"`
	Password      string `gorm:"type:varchar(64)"`
	Avatar        string `gorm:"type:varchar(64);null;comment:'user avatar'"`
	Cover         string `gorm:"type:varchar(64);null;comment:'user cover'"`
	StatusMessage string `gorm:"type:varchar(64);null;comment:'user status message'"`
	CommonField

	Stories []StoryModel `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) error {
	u.Uuid = uuid.New().String()
	u.Avatar = "/default.jpg"
	u.Cover = "/cover.jpg"
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
	return db.Debug().WithContext(ctx).Where("email = ?", u.Email).First(&u).Error
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

func (u *UserModel) UpdateOneUserStatus(db *gorm.DB, ctx context.Context, message string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.ID).Update("StatusMessage", message).Error
}

func (u *UserModel) UpdateOneUserAvatar(db *gorm.DB, ctx context.Context, avatarPath string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.ID).Update("Avatar", avatarPath).Error
}

func (u *UserModel) UpdateOneUserCover(db *gorm.DB, ctx context.Context, coverPath string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.ID).Update("Cover", coverPath).Error
}

func (u *UserModel) FindUsers(db *gorm.DB, ctx context.Context, query string) ([]*UserModel, error) {
	var results []*UserModel
	if err := db.WithContext(ctx).Debug().Model(&u).Where("nick_name like ?", "%"+query+"%").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (u *UserModel) CountUserStory(db *gorm.DB, ctx context.Context) (int64, error) {
	now := time.Now().Unix()
	availableTime := now - 86400
	if err := db.WithContext(ctx).Debug().Preload("Stories", "created_at BETWEEN FROM_UINXTIME(?) AND FROM_UNIXTIME(?)", availableTime, now).Where("id = ?", u.ID).First(&u).Error; err != nil {
		return 0, err
	}

	return int64(len(u.Stories)), nil
}

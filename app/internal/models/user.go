package models

import (
	"context"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id            uint   `gorm:"primaryKey;autoIncrement;not null"`
	Uuid          string `gorm:"types:varchar(64);not null;unique_index:idx_uuid"`
	NickName      string `gorm:"types:varchar(32)"`
	Email         string `gorm:"types:varchar(64)"`
	Password      string `gorm:"types:varchar(64)"`
	Avatar        string `gorm:"types:varchar(64);null;comment:'user avatar'"`
	Cover         string `gorm:"types:varchar(64);null;comment:'user cover'"`
	StatusMessage string `gorm:"types:varchar(64);null;comment:'user status message'"`
	CommonField

	Stories       []StoryModel `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Groups        []Group      `gorm:"many2many:users_groups;foreignKey:Id;joinForeignKey:UserId"`
	StickerGroups []Sticker    `gorm:"many2many:users_stickers;foreignKey:Id;joinForeignKey:UserId;References:Uuid;joinReferences:StickerId"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.Uuid = uuid.New().String()
	u.Avatar = "/default/default.jpg"
	u.Cover = "/default/cover.jpg"
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	tx.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}

func (u *User) TableName() string {
	return "users_info"
}

func (u *User) FindOneUserByID(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Preload("StickerGroups").First(&u).Error
}

func (u *User) FindOneUserByUUID(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Preload("StickerGroups").Where("uuid = ?", u.Uuid).First(&u).Error
}

func (u *User) FindOneUserByEmail(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Where("email = ?", u.Email).First(&u).Error
}

func (u *User) InsertOneUser(db *gorm.DB, ctx context.Context) error {
	return db.Debug().WithContext(ctx).Create(u).Error
}

func (u *User) DeleteOneUser() error {
	return nil
}

func (u *User) UpdateOneUser(db *gorm.DB, ctx context.Context, name string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.Id).Update("NickName", name).Error
}

func (u *User) UpdateOneUserStatus(db *gorm.DB, ctx context.Context, message string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.Id).Update("StatusMessage", message).Error
}

func (u *User) UpdateOneUserAvatar(db *gorm.DB, ctx context.Context, avatarPath string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.Id).Update("Avatar", avatarPath).Error
}

func (u *User) UpdateOneUserCover(db *gorm.DB, ctx context.Context, coverPath string) error {
	return db.Debug().WithContext(ctx).Model(u).Where("id = ?", u.Id).Update("Cover", coverPath).Error
}

func (u *User) FindUsers(db *gorm.DB, ctx context.Context, query string) ([]*User, error) {
	var results []*User
	if err := db.WithContext(ctx).Debug().Model(&u).Where("nick_name like ?", "%"+query+"%").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (u *User) CountUserStory(db *gorm.DB, ctx context.Context) (int64, error) {
	now := time.Now().Unix()
	availableTime := now - 86400
	if err := db.WithContext(ctx).Debug().Preload("Stories", "created_at BETWEEN FROM_UINXTIME(?) AND FROM_UNIXTIME(?)", availableTime, now).Where("id = ?", u.Id).First(&u).Error; err != nil {
		return 0, err
	}

	return int64(len(u.Stories)), nil
}

func (u *User) CountUserGroup(db *gorm.DB, ctx context.Context) int64 {
	count := db.WithContext(ctx).Model(&u).Association("Groups").Count()
	return count
}

func (u *User) JoinGroup(db *gorm.DB, ctx context.Context, group *Group) error {
	return db.WithContext(ctx).Model(&u).Association("GroupInfos").Append(group)
}

func (u *User) InsertOneSticker(db *gorm.DB, ctx context.Context, sticker *Sticker) error {
	return db.WithContext(ctx).Debug().Omit("StickerGroups.*").Model(&u).Association("StickerGroups").Append(sticker)
}

func (u *User) FindOneSticker(db *gorm.DB, ctx context.Context, stickerUUID string) (*Sticker, error) {
	var sticker Sticker
	if err := db.WithContext(ctx).Debug().Model(&u).Association("StickerGroups").Find(&sticker, "sticker_id = ?", stickerUUID); err != nil {
		return nil, err
	}
	return &sticker, nil
}

func (u *User) FindAllSticker(db *gorm.DB, ctx context.Context) ([]*Sticker, error) {
	var sticker []*Sticker
	logx.Info("Testing final all stickers")
	if err := db.WithContext(ctx).Debug().Model(&u).Association("StickerGroups").Find(&sticker); err != nil {
		return nil, err
	}
	return sticker, nil
}

func (u *User) DeleteOneSticker(db *gorm.DB, ctx context.Context, sticker *Sticker) error {
	return db.WithContext(ctx).Model(&u).Association("StickerGroups").Delete(sticker)
}

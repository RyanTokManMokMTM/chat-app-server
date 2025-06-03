package repository

import (
	"context"
	"time"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
)

type UpdateFieldType string

const (
	UpdateFieldStatus   UpdateFieldType = "StatusMessage"
	UpdateFieldAvatar   UpdateFieldType = "Avatar"
	UpdateFieldCover    UpdateFieldType = "Cover"
	UpdateFieldNickName UpdateFieldType = "NickName"
)

type IUserRepo[T any] interface {
	CreateOne(ctx context.Context, user T) (T, error)
	FindOne(ctx context.Context, user T) (T, error)
	UpdateOne(ctx context.Context, user T) (T, error)
	DeleteOne(ctx context.Context, UUID string) error

	// TODO: Add other methods
	FindOneUserByID(ctx context.Context, id uint) (T, error)
	FindOneUserByUUID(ctx context.Context, UUID string) (T, error)
	FindOneUserByEmail(ctx context.Context, email string) (T, error)
	UpdateOneUserStatusMessage(ctx context.Context, id uint, status string) error
	UpdateOneUserAvatar(ctx context.Context, id uint, path string) error
	UpdateOneUserCover(ctx context.Context, id uint, path string) error
	UpdateOneUserNickName(ctx context.Context, id uint, name string) error
	FindUsersByNickName(ctx context.Context, query string) ([]T, error)
	FindUsers(ctx context.Context, query string) ([]T, error)
	CountUserStorys(ctx context.Context, id uint) (int64, error)
	CountUserGroups(ctx context.Context, id uint) int64

	// Sticker related methods
	// FIXME: should put in sticker repo....
	InsertOneSticker(ctx context.Context, userID uint, stickerUUID string) error
	FindOneSticker(ctx context.Context, userID uint, stickerUUID string) (*models.Sticker, error)
	FindAllSticker(ctx context.Context, userID uint) ([]*models.Sticker, error)
	DeleteOneSticker(ctx context.Context, userID uint, sticker *models.Sticker) error
}

var _ IUserRepo[*models.User] = (*userRepo)(nil)

type userRepo struct {
	IRepository[*models.User, models.User]
	engine *gorm.DB
}

func NewUserRepo(engine *gorm.DB) IUserRepo[*models.User] {
	return &userRepo{
		IRepository: NewRepository[*models.User, models.User](engine),
		engine:      engine,
	}
}

func (userRepo *userRepo) CreateOne(ctx context.Context, user *models.User) (*models.User, error) {
	return userRepo.Create(ctx, user)
}

func (userRepo *userRepo) FindOne(ctx context.Context, user *models.User) (*models.User, error) {
	return userRepo.Find(ctx, user)
}

func (userRepo *userRepo) UpdateOne(ctx context.Context, user *models.User) (*models.User, error) {
	if err := userRepo.Update(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepo *userRepo) DeleteOne(ctx context.Context, UUID string) error {
	return userRepo.Delete(ctx, &models.User{Base: models.Base{UUID: UUID}})
}

func (userRepo *userRepo) FindOneUserByID(ctx context.Context, id uint) (*models.User, error) {
	u := &models.User{Base: models.Base{ID: id}}
	err := userRepo.engine.Debug().WithContext(ctx).Preload("StickerGroups").First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userRepo *userRepo) FindOneUserByUUID(ctx context.Context, UUID string) (*models.User, error) {
	u := &models.User{Base: models.Base{UUID: UUID}}
	err := userRepo.engine.Debug().WithContext(ctx).Preload("StickerGroups").First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userRepo *userRepo) FindOneUserByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{Email: email}
	err := userRepo.engine.Debug().WithContext(ctx).Where("email = ?", u.Email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userRepo *userRepo) FindUsersByNickName(ctx context.Context, query string) ([]*models.User, error) {
	var results []*models.User
	if err := userRepo.engine.Debug().WithContext(ctx).Model(&models.User{}).Where("nick_name like ?", "%"+query+"%").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// updateUserField updates a specific field of a user
func (userRepo *userRepo) updateUserField(ctx context.Context, userID uint, field UpdateFieldType, value interface{}) error {
	return userRepo.engine.WithContext(ctx).Debug().
		Model(&models.User{}).
		Where("id = ?", userID).
		Update(string(field), value).Error
}

// UpdateOneUserStatus updates user's status message
func (userRepo *userRepo) UpdateOneUserStatusMessage(ctx context.Context, id uint, status string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldStatus, status)
}

// UpdateOneUserAvatar updates user's avatar
func (userRepo *userRepo) UpdateOneUserAvatar(ctx context.Context, id uint, path string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldAvatar, path)
}

// UpdateOneUserCover updates user's cover image
func (userRepo *userRepo) UpdateOneUserCover(ctx context.Context, id uint, path string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldCover, path)
}

// UpdateOneUserNickName updates user's nickname
func (userRepo *userRepo) UpdateOneUserNickName(ctx context.Context, id uint, name string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldNickName, name)
}

func (userRepo *userRepo) FindUsers(ctx context.Context, query string) ([]*models.User, error) {
	var results []*models.User
	if err := userRepo.engine.Debug().WithContext(ctx).Model(&models.User{}).Where("nick_name like ?", "%"+query+"%").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (userRepo *userRepo) CountUserStorys(ctx context.Context, id uint) (int64, error) {
	now := time.Now().Unix()
	availableTime := now - 86400 // 24 hours ago
	var count int64

	if err := userRepo.engine.WithContext(ctx).Debug().
		Model(&models.StoryModel{}).
		Where("user_id = ? AND created_at BETWEEN FROM_UNIXTIME(?) AND FROM_UNIXTIME(?)",
			id, availableTime, now).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (userRepo *userRepo) CountUserGroups(ctx context.Context, id uint) int64 {
	var count int64
	userRepo.engine.WithContext(ctx).Model(&models.UserGroup{}).
		Where("user_id = ?", id).
		Count(&count)
	return count
}

func (userRepo *userRepo) InsertOneSticker(ctx context.Context, userID uint, stickerUUID string) error {
	user := &models.User{Base: models.Base{ID: userID}}
	return userRepo.engine.WithContext(ctx).Debug().
		Model(user).
		Association("StickerGroups").
		Append(&models.Sticker{
			Base: models.Base{UUID: stickerUUID},
		})
}

func (userRepo *userRepo) FindOneSticker(ctx context.Context, userID uint, stickerUUID string) (*models.Sticker, error) {
	var sticker models.Sticker
	user := &models.User{Base: models.Base{ID: userID}}

	if err := userRepo.engine.WithContext(ctx).Debug().
		Model(user).
		Association("StickerGroups").
		Find(&sticker, "UUID = ?", stickerUUID); err != nil {
		return nil, err
	}
	return &sticker, nil
}

func (userRepo *userRepo) FindAllSticker(ctx context.Context, userID uint) ([]*models.Sticker, error) {
	var stickers []*models.Sticker
	user := &models.User{Base: models.Base{ID: userID}}

	if err := userRepo.engine.WithContext(ctx).Debug().
		Model(user).
		Association("StickerGroups").
		Find(&stickers); err != nil {
		return nil, err
	}
	return stickers, nil
}

func (userRepo *userRepo) DeleteOneSticker(ctx context.Context, userID uint, sticker *models.Sticker) error {
	user := &models.User{Base: models.Base{ID: userID}}
	return userRepo.engine.WithContext(ctx).
		Model(user).
		Association("StickerGroups").
		Delete(sticker)
}

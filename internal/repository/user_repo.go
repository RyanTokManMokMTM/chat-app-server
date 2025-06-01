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
	CreateOne(ctx context.Context, user T) (*T, error)
	FindOne(ctx context.Context, user T) (*T, error)
	UpdateOne(ctx context.Context, user T) (*T, error)
	DeleteOne(ctx context.Context, uuid string) error

	// TODO: Add other methods
	FindOneUserByID(ctx context.Context, id uint) (*T, error)
	FindOneUserByUUID(ctx context.Context, uuid string) (*T, error)
	FindOneUserByEmail(ctx context.Context, email string) (*T, error)
	UpdateOneUserStatusMessage(ctx context.Context, id uint, status string) error
	UpdateOneUserAvatar(ctx context.Context, id uint, path string) error
	UpdateOneUserCover(ctx context.Context, id uint, path string) error
	UpdateOneUserNickName(ctx context.Context, id uint, name string) error
	FindUsersByNickName(ctx context.Context, query string) ([]T, error)
	FindUsers(ctx context.Context, query string) ([]models.User, error)
	CountUserStorys(ctx context.Context, id uint) (int64, error)
	CountUserGroups(ctx context.Context, id uint) int64
	JoinGroup(ctx context.Context, group models.Group) error

	// Sticker related methods
	InsertOneSticker(ctx context.Context, userId uint, sticker models.Sticker) error
	FindOneSticker(ctx context.Context, userId uint, stickerUUID string) (*models.Sticker, error)
	FindAllSticker(ctx context.Context, userId uint) ([]*models.Sticker, error)
	DeleteOneSticker(ctx context.Context, userId uint, sticker *models.Sticker) error
}

var _ IUserRepo[models.User] = (*UserRepo)(nil)

type UserRepo struct {
	IRepository[*models.User, models.User]
	engine *gorm.DB
}

func NewUserRepo(engine *gorm.DB) IUserRepo[models.User] {
	return &UserRepo{
		IRepository: NewRepository[*models.User, models.User](engine),
		engine:      engine,
	}
}

func (userRepo *UserRepo) CreateOne(ctx context.Context, user models.User) (*models.User, error) {
	if err := userRepo.Create(ctx, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *UserRepo) FindOne(ctx context.Context, user models.User) (*models.User, error) {
	result, err := userRepo.Find(ctx, &user)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (userRepo *UserRepo) UpdateOne(ctx context.Context, user models.User) (*models.User, error) {
	if err := userRepo.Update(ctx, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepo *UserRepo) DeleteOne(ctx context.Context, uuid string) error {
	return userRepo.Delete(ctx, &models.User{Uuid: uuid})
}

func (userRepo *UserRepo) FindOneUserByID(ctx context.Context, id uint) (*models.User, error) {
	u := &models.User{Id: id}
	err := userRepo.engine.Debug().WithContext(ctx).Preload("StickerGroups").First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userRepo *UserRepo) FindOneUserByUUID(ctx context.Context, uuid string) (*models.User, error) {
	u := &models.User{Uuid: uuid}
	err := userRepo.engine.Debug().WithContext(ctx).Preload("StickerGroups").First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userRepo *UserRepo) FindOneUserByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{Email: email}
	err := userRepo.engine.Debug().WithContext(ctx).Where("email = ?", u.Email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (userRepo *UserRepo) FindUsersByNickName(ctx context.Context, query string) ([]models.User, error) {
	var results []models.User
	if err := userRepo.engine.Debug().WithContext(ctx).Model(&models.User{}).Where("nick_name like ?", "%"+query+"%").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// updateUserField updates a specific field of a user
func (userRepo *UserRepo) updateUserField(ctx context.Context, userId uint, field UpdateFieldType, value interface{}) error {
	return userRepo.engine.WithContext(ctx).Debug().
		Model(&models.User{}).
		Where("id = ?", userId).
		Update(string(field), value).Error
}

// UpdateOneUserStatus updates user's status message
func (userRepo *UserRepo) UpdateOneUserStatusMessage(ctx context.Context, id uint, status string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldStatus, status)
}

// UpdateOneUserAvatar updates user's avatar
func (userRepo *UserRepo) UpdateOneUserAvatar(ctx context.Context, id uint, path string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldAvatar, path)
}

// UpdateOneUserCover updates user's cover image
func (userRepo *UserRepo) UpdateOneUserCover(ctx context.Context, id uint, path string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldCover, path)
}

// UpdateOneUserNickName updates user's nickname
func (userRepo *UserRepo) UpdateOneUserNickName(ctx context.Context, id uint, name string) error {
	return userRepo.updateUserField(ctx, id, UpdateFieldNickName, name)
}

func (userRepo *UserRepo) FindUsers(ctx context.Context, query string) ([]models.User, error) {
	var results []models.User
	if err := userRepo.engine.Debug().WithContext(ctx).Model(&models.User{}).Where("nick_name like ?", "%"+query+"%").Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (userRepo *UserRepo) CountUserStorys(ctx context.Context, id uint) (int64, error) {
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

func (userRepo *UserRepo) CountUserGroups(ctx context.Context, id uint) int64 {
	var count int64
	userRepo.engine.WithContext(ctx).Model(&models.UserGroup{}).
		Where("user_id = ?", id).
		Count(&count)
	return count
}

func (userRepo *UserRepo) JoinGroup(ctx context.Context, group models.Group) error {
	userGroup := &models.UserGroup{
		UserId:  group.GroupLead,
		GroupId: group.Id,
	}
	return userRepo.engine.WithContext(ctx).Debug().Create(userGroup).Error
}

func (userRepo *UserRepo) InsertOneSticker(ctx context.Context, userId uint, sticker models.Sticker) error {
	user := &models.User{Id: userId}
	return userRepo.engine.WithContext(ctx).Debug().
		Model(user).
		Association("StickerGroups").
		Append(&sticker)
}

func (userRepo *UserRepo) FindOneSticker(ctx context.Context, userId uint, stickerUUID string) (*models.Sticker, error) {
	var sticker models.Sticker
	user := &models.User{Id: userId}

	if err := userRepo.engine.WithContext(ctx).Debug().
		Model(user).
		Association("StickerGroups").
		Find(&sticker, "uuid = ?", stickerUUID); err != nil {
		return nil, err
	}
	return &sticker, nil
}

func (userRepo *UserRepo) FindAllSticker(ctx context.Context, userId uint) ([]*models.Sticker, error) {
	var stickers []*models.Sticker
	user := &models.User{Id: userId}

	if err := userRepo.engine.WithContext(ctx).Debug().
		Model(user).
		Association("StickerGroups").
		Find(&stickers); err != nil {
		return nil, err
	}
	return stickers, nil
}

func (userRepo *UserRepo) DeleteOneSticker(ctx context.Context, userId uint, sticker *models.Sticker) error {
	user := &models.User{Id: userId}
	return userRepo.engine.WithContext(ctx).
		Model(user).
		Association("StickerGroups").
		Delete(sticker)
}

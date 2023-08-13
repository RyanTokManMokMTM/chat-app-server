package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
)

func (d *DAO) InsertOneUser(ctx context.Context, name, email, password string) (*models.User, error) {
	u := &models.User{
		NickName: name,
		Email:    email,
		Password: password,
	}
	if err := u.InsertOneUser(d.engine, ctx); err != nil {
		return nil, err
	}
	return u, nil
}
func (d *DAO) FindOneUser(ctx context.Context, id uint) (*models.User, error) {
	u := &models.User{
		Id: id,
	}
	if err := u.FindOneUserByID(d.engine, ctx); err != nil {
		return nil, err
	}

	return u, nil
}

func (d *DAO) FindOneUserByEmail(ctx context.Context, email string) (*models.User, error) {
	u := &models.User{
		Email: email,
	}

	if err := u.FindOneUserByEmail(d.engine, ctx); err != nil {
		return nil, err
	}
	return u, nil
}

func (d *DAO) FindOneUserByUUID(ctx context.Context, uuid string) (*models.User, error) {
	u := &models.User{
		Uuid: uuid,
	}

	if err := u.FindOneUserByUUID(d.engine, ctx); err != nil {
		return nil, err
	}
	return u, nil
}

func (d *DAO) UpdateUserProfile(ctx context.Context, id uint, name string) error {
	u := &models.User{
		Id: id,
	}

	return u.UpdateOneUser(d.engine, ctx, name)
}

func (d *DAO) UpdateUserStatusMessage(ctx context.Context, id uint, message string) error {
	u := &models.User{
		Id: id,
	}

	return u.UpdateOneUserStatus(d.engine, ctx, message)
}

func (d *DAO) UpdateUserAvatar(ctx context.Context, id uint, avatarPath string) error {
	u := &models.User{
		Id: id,
	}

	return u.UpdateOneUserAvatar(d.engine, ctx, avatarPath)
}

func (d *DAO) UpdateUserCover(ctx context.Context, id uint, coverPath string) error {
	u := &models.User{
		Id: id,
	}

	return u.UpdateOneUserCover(d.engine, ctx, coverPath)
}

func (d *DAO) FindUsers(ctx context.Context, query string) ([]*models.User, error) {
	return (&models.User{}).FindUsers(d.engine, ctx, query)
}

func (d *DAO) CountUserAvailableStory(ctx context.Context, userID uint) (int64, error) {
	u := &models.User{
		Id: userID,
	}
	return u.CountUserStory(d.engine, ctx)
}

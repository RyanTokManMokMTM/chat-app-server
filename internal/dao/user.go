package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

func (d *DAO) InsertOneUser(ctx context.Context, name, email, password string) (*models.UserModel, error) {
	u := &models.UserModel{
		NickName: name,
		Email:    email,
		Password: password,
	}
	if err := u.InsertOneUser(d.engine, ctx); err != nil {
		return nil, err
	}
	return u, nil
}
func (d *DAO) FindOneUser(ctx context.Context, id uint) (*models.UserModel, error) {
	u := &models.UserModel{
		ID: id,
	}
	if err := u.FindOneUserByID(d.engine, ctx); err != nil {
		return nil, err
	}

	return u, nil
}

func (d *DAO) FindOneUserByEmail(ctx context.Context, email string) (*models.UserModel, error) {
	u := &models.UserModel{
		Email: email,
	}

	if err := u.FindOneUserByEmail(d.engine, ctx); err != nil {
		return nil, err
	}
	return u, nil
}

func (d *DAO) UpdateUserProfile(ctx context.Context, id uint, name string) error {
	u := &models.UserModel{
		ID: id,
	}

	return u.UpdateOneUser(d.engine, ctx, name)
}

func (d *DAO) UpdateUserAvatar(ctx context.Context, id uint, avatarName string) error {
	u := &models.UserModel{
		ID: id,
	}

	return u.UpdateOneUserAvatar(d.engine, ctx, avatarName)
}

func (d *DAO) FindUsers(ctx context.Context, query string) ([]*models.UserModel, error) {
	return (&models.UserModel{}).FindUsers(d.engine, ctx, query)
}

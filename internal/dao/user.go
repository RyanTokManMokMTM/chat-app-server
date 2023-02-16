package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

func (d *DAO) InsertOneUser(ctx context.Context, name, email, password string) error {
	u := &models.UserModel{
		Name:     name,
		Email:    email,
		Password: password,
	}
	return u.InsertOneUser(d.engine, ctx)
}
func (d *DAO) FindOneUser(ctx context.Context, id uint) (*models.UserModel, error) {
	u := &models.UserModel{
		ID: id,
	}
	if err := u.FindOneUser(d.engine, ctx); err != nil {
		return nil, err
	}

	return u, nil
}

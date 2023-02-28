package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

func (d *DAO) InsertOneFriend(ctx context.Context, friendID uint) error {
	return nil
}
func (d *DAO) FindOneFriend(ctx context.Context, friendID uint) (bool, error) {
	return false, nil
}
func (d *DAO) DeleteOneFriend(ctx context.Context, friendID uint) error {
	return nil
}
func (d *DAO) GetUserFriendList(ctx context.Context) ([]*models.UserModel, error) {
	return nil, nil
}

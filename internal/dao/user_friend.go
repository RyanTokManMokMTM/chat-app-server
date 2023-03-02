package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

func (d *DAO) InsertOneFriend(ctx context.Context, userID, friendID uint) error {
	uf := &models.UserFriend{
		UserID:   userID,
		FriendID: friendID,
	}
	return uf.InsertOne(ctx, d.engine)
}
func (d *DAO) FindOneFriend(ctx context.Context, userID, friendID uint) error {
	uf := &models.UserFriend{}
	err := uf.FindOne(ctx, d.engine, userID, friendID)

	return err
}
func (d *DAO) DeleteOneFriend(ctx context.Context, userID, friendID uint) error {
	uf := &models.UserFriend{
		UserID:   userID,
		FriendID: friendID,
	}
	return uf.DeleteOne(ctx, d.engine)
}
func (d *DAO) GetUserFriendList(ctx context.Context, userID uint) ([]*models.UserFriend, error) {
	uf := &models.UserFriend{
		UserID: userID,
	}

	//TODO: It must include user info
	return uf.GetFriendList(ctx, d.engine)
}

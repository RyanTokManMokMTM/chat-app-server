package dao

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
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

func (d *DAO) GetUserFriendListByPageSize(ctx context.Context, userID uint, pageOffset, PageLimit int) ([]*models.UserFriend, error) {
	uf := &models.UserFriend{
		UserID: userID,
	}

	return uf.GetFriendList(ctx, d.engine, pageOffset, PageLimit)
}

func (d *DAO) CountUserFriend(ctx context.Context, userID uint) (int64, error) {
	uf := &models.UserFriend{
		UserID: userID,
	}

	return uf.CountUserFriends(ctx, d.engine)
}

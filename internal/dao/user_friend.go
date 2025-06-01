package dao

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
)

func (d *DAO) InsertOneFriend(ctx context.Context, userId, friendID uint) error {
	uf := &models.UserFriend{
		userId:   userId,
		FriendID: friendID,
	}
	return uf.InsertOne(ctx, d.engine)
}
func (d *DAO) FindOneFriend(ctx context.Context, userId, friendID uint) (*models.User, error) {
	uf := &models.UserFriend{
		userId:   userId,
		FriendID: friendID,
	}
	if err := uf.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}

	return &uf.FriendInfo, nil
}
func (d *DAO) DeleteOneFriend(ctx context.Context, userId, friendID uint) error {
	uf := &models.UserFriend{
		userId:   userId,
		FriendID: friendID,
	}
	return uf.DeleteOne(ctx, d.engine)
}

func (d *DAO) GetUserFriendListByPageSize(ctx context.Context, userId uint, pageOffset, PageLimit int) ([]*models.UserFriend, error) {
	uf := &models.UserFriend{
		userId: userId,
	}

	return uf.GetFriendList(ctx, d.engine, pageOffset, PageLimit)
}

func (d *DAO) CountUserFriend(ctx context.Context, userId uint) (int64, error) {
	uf := &models.UserFriend{
		userId: userId,
	}

	return uf.CountUserFriends(ctx, d.engine)
}

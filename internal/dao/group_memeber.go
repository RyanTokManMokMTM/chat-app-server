package dao

import (
	"context"
	"github.com/ryantokmanmok/chat-app-server/internal/models"
)

func (d *DAO) InsertOneGroupMember(ctx context.Context, groupID, userID uint) error {
	gm := &models.GroupMember{
		GroupID: groupID,
		UserID:  userID,
	}

	return gm.InsertOne(ctx, d.engine)
}
func (d *DAO) FindOneGroupMember(ctx context.Context, groupID, userID uint) (*models.GroupMember, error) {
	gm := &models.GroupMember{
		GroupID: groupID,
		UserID:  userID,
	}

	if err := gm.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return gm, nil
}

func (d *DAO) FindOneGroupMembers(ctx context.Context, groupID uint) ([]*models.GroupMember, error) {
	gm := &models.GroupMember{
		GroupID: groupID,
	}

	return gm.FindAll(ctx, d.engine)
}

func (d *DAO) DeleteGroupMember(ctx context.Context, groupID, userID uint) error {
	gm := &models.GroupMember{
		GroupID: groupID,
		UserID:  userID,
	}

	return gm.DeleteOne(ctx, d.engine)
}

func (d *DAO) DeleteAllGroupMembers(ctx context.Context, groupID uint) error {
	gm := &models.GroupMember{
		GroupID: groupID,
	}

	return gm.DeleteAll(ctx, d.engine)
}
func (d *DAO) GetGroupMembers(ctx context.Context, groupID uint) ([]*models.GroupMember, error) {
	gm := &models.GroupMember{
		GroupID: groupID,
	}

	return gm.GetGroupMemberList(ctx, d.engine)

}

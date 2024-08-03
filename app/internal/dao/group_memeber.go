package dao

import (
	"api/app/internal/models"
	"context"
)

func (d *DAO) InsertOneGroupMember(ctx context.Context, groupID, userID uint) error {
	gm := &models.UserGroup{
		GroupId: groupID,
		UserId:  userID,
	}

	return gm.InsertOne(ctx, d.engine)
}
func (d *DAO) FindOneGroupMember(ctx context.Context, groupID, userID uint) (*models.UserGroup, error) {
	gm := &models.UserGroup{
		GroupId: groupID,
		UserId:  userID,
	}

	if err := gm.FindOne(ctx, d.engine); err != nil {
		return nil, err
	}
	return gm, nil
}

func (d *DAO) FindOneGroupMembers(ctx context.Context, groupID uint) ([]*models.UserGroup, error) {
	gm := &models.UserGroup{
		GroupId: groupID,
	}

	return gm.GetGroupMemberList(ctx, d.engine)
}

func (d *DAO) FindOneGroupMembersByPage(ctx context.Context, groupID uint, pageOffset, pageLimit int) ([]*models.UserGroup, error) {
	gm := &models.UserGroup{
		GroupId: groupID,
	}

	return gm.GetGroupMemberListByPage(ctx, d.engine, pageOffset, pageLimit)
}

func (d *DAO) DeleteGroupMember(ctx context.Context, groupID, userID uint) error {
	gm := &models.UserGroup{
		GroupId: groupID,
		UserId:  userID,
	}

	return gm.DeleteOne(ctx, d.engine)
}

func (d *DAO) DeleteAllGroupMembers(ctx context.Context, groupID uint) error {
	gm := &models.UserGroup{
		GroupId: groupID,
	}

	return gm.DeleteAll(ctx, d.engine)
}
func (d *DAO) GetGroupMembers(ctx context.Context, groupID uint, pageOffset, pageLimit int) ([]*models.UserGroup, error) {
	gm := &models.UserGroup{
		GroupId: groupID,
	}

	return gm.GetGroupMemberListByPage(ctx, d.engine, pageOffset, pageLimit)

}

func (d *DAO) GetUserGroups(ctx context.Context, userID uint, pageOffset, pageLimit int) ([]*models.UserGroup, error) {
	gm := &models.UserGroup{
		UserId: userID,
	}
	return gm.FindUserGroup(ctx, d.engine, pageOffset, pageLimit)
}

func (d *DAO) CountUserGroups(ctx context.Context, userID uint) int64 {
	u := &models.User{
		Id: userID,
	}

	return u.CountUserGroup(d.engine, ctx)
}

func (d *DAO) CountGroupMembers(ctx context.Context, groupID uint) (int64, error) {
	gm := &models.UserGroup{
		GroupId: groupID,
	}
	return gm.CountGroupMembers(ctx, d.engine)

}

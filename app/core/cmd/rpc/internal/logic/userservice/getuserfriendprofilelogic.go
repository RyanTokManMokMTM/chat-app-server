package userservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"api/app/internal/models"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFriendProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFriendProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFriendProfileLogic {
	return &GetUserFriendProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFriendProfileLogic) GetUserFriendProfile(in *core.GetUserFriendProfileReq) (*core.GetUserFriendProfileResp, error) {
	// todo: add your logic here and delete this line
	if in.FriendUuid == nil && in.FriendUserId == nil {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "request parameter invalid")
	}

	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error: %+v", err)
		}
		return nil, err
	}

	var u *models.User
	if in.FriendUserId != nil {
		friendUserId := uint(*in.FriendUserId)
		_, err := l.svcCtx.DAO.FindOneUser(l.ctx, friendUserId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
			}
			return nil, err
		}
	} else {
		_, err := l.svcCtx.DAO.FindOneUserByUUID(l.ctx, *in.FriendUuid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
			}
			return nil, err
		}
	}

	isFriend := true
	_, err = l.svcCtx.DAO.FindOneFriend(l.ctx, userID, u.Id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		isFriend = false
	}

	return &core.GetUserFriendProfileResp{
		Code: int32(errx.SUCCESS),
		UserInfo: &core.UserInfo{
			Id:            uint32(u.Id),
			Uuid:          u.Uuid,
			Name:          u.NickName,
			Email:         u.Email,
			Avatar:        u.Avatar,
			Cover:         u.Cover,
			StatusMessage: u.StatusMessage,
		}, IsFriend: isFriend,
	}, nil
}

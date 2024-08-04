package userservicelogic

import (
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"api/app/internal/models"
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
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
	if in.Uuid == nil && in.UserId == nil {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}

	userID := uint(*in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var u *models.User
	if userID != 0 {
		_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
			}
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
	} else {
		_, err := l.svcCtx.DAO.FindOneUserByUUID(l.ctx, *in.Uuid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
			}
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
	}

	isFriend := true
	_, err = l.svcCtx.DAO.FindOneFriend(l.ctx, userID, u.Id)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
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

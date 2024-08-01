package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFriendProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserFriendProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFriendProfileLogic {
	return &GetUserFriendProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFriendProfileLogic) GetUserFriendProfile(req *types.GetUserFriendProfileReq) (resp *types.GetUserFriendProfileResp, err error) {
	// todo: add your logic here and delete this line

	if req.UUID == "" && req.UserID == 0 {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}

	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var u *models.User
	if req.UserID != 0 {
		u, err = l.svcCtx.DAO.FindOneUser(l.ctx, req.UserID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
			}
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
	} else {
		u, err = l.svcCtx.DAO.FindOneUserByUUID(l.ctx, req.UUID)
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
	return &types.GetUserFriendProfileResp{
		Code: uint(http.StatusOK),
		UserInfo: types.CommonUserInfo{
			ID:            u.Id,
			Uuid:          u.Uuid,
			NickName:      u.NickName,
			Email:         u.Email,
			Avatar:        u.Avatar,
			Cover:         u.Cover,
			StatusMessage: u.StatusMessage,
		}, IsFriend: isFriend,
	}, nil
}

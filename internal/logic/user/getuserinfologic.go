package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"gorm.io/gorm"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	if req.UUID == "" && req.UserID == 0 {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}

	var u *models.UserModel
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

	return &types.GetUserInfoResp{
		Code:          uint(http.StatusOK),
		UUID:          u.Uuid,
		Email:         u.Email,
		Name:          u.NickName,
		Avatar:        u.Avatar,
		Cover:         u.Cover,
		StatusMessage: u.StatusMessage,
	}, nil
}

package userservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"api/app/internal/models"
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *core.GetUserInfoReq) (*core.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	if in.Uuid == nil && in.UserId == nil {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}

	userId := uint(*in.UserId)
	var u *models.User
	if userId != 0 {
		_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userId)
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

	return &core.GetUserInfoResp{
		Code:          int32(errx.SUCCESS),
		Uuid:          u.Uuid,
		Email:         u.Email,
		Name:          u.NickName,
		Avatar:        u.Avatar,
		Cover:         u.Cover,
		StatusMessage: u.StatusMessage,
	}, nil
}

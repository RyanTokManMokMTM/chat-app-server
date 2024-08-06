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
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "request parameter invaild")
	}

	userId := uint(*in.UserId)
	var u *models.User
	if userId != 0 {
		logx.Info("Get user by id")
		found, err := l.svcCtx.DAO.FindOneUser(l.ctx, userId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
			}
			logx.WithContext(l.ctx).Errorf("Error : %+v", err)
			return nil, err
		}
		u = found
	} else {
		found, err := l.svcCtx.DAO.FindOneUserByUUID(l.ctx, *in.Uuid)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
			}
			logx.WithContext(l.ctx).Errorf("Error : %+v", err)
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

		u = found
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

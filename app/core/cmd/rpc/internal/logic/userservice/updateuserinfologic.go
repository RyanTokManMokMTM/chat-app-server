package userservicelogic

import (
	"api/app/common/errx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(in *core.UpdateUserInfoReq) (*core.UpdateUserInfoResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	u, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error: %+v", err)
		}
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}
	if strings.Compare(in.Name, u.NickName) == 0 {
		return &core.UpdateUserInfoResp{
			Code: int32(errx.SUCCESS),
		}, nil
	}

	if err := l.svcCtx.DAO.UpdateUserProfile(l.ctx, userID, in.Name); err != nil {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	return &core.UpdateUserInfoResp{
		Code: int32(errx.SUCCESS),
	}, nil
}

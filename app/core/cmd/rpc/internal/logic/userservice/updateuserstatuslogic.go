package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserStatusLogic {
	return &UpdateUserStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserStatusLogic) UpdateUserStatus(in *core.UpdateUserStatusReq) (*core.UpdateUserStatusResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist,error : %+v", err)
		}
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	if err := l.svcCtx.DAO.UpdateUserStatusMessage(l.ctx, userID, in.Status); err != nil {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	return &core.UpdateUserStatusResp{
		Code: int32(errx.SUCCESS),
	}, nil
}

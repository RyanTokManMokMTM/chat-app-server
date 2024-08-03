package groupservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountUserGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountUserGroupLogic {
	return &CountUserGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CountUserGroupLogic) CountUserGroup(in *core.CountUserGroupReq) (*core.CountUserGroupResp, error) {
	// todo: add your logic here and delete this line

	return &core.CountUserGroupResp{}, nil
}

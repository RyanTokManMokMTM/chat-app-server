package group

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountUserGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Count user group
func NewCountUserGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountUserGroupLogic {
	return &CountUserGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountUserGroupLogic) CountUserGroup(req *types.CountUserGroupReq) (resp *types.CountUserGroupResp, err error) {
	// todo: add your logic here and delete this line

	return
}

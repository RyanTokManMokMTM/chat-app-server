package group

import (
	"api/app/common/ctxtool"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.GroupService.CountUserGroup(l.ctx, &core.CountUserGroupReq{
		UserId: uint32(userID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &types.CountUserGroupResp{
		Code:  uint(rpcResp.Code),
		Total: uint(rpcResp.Total),
	}, nil
}

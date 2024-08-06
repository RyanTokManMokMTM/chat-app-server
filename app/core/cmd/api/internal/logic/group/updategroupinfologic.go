package group

import (
	"api/app/common/ctxtool"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update and update group info
func NewUpdateGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupInfoLogic {
	return &UpdateGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateGroupInfoLogic) UpdateGroupInfo(req *types.UpdateGroupInfoReq) (resp *types.UpdateGroupInfoResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.GroupService.UpdateGroupInfo(l.ctx, &core.UpdateGroupInfoReq{
		UserId:    uint32(userID),
		GroupId: uint32(req.GroupID),
		GroupName: req.GroupName,
		GroupDesc: req.GroupDesc,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}s

	return &types.UpdateGroupInfoResp{
		Code: uint(rpcResp.Code),
	}, nil
}

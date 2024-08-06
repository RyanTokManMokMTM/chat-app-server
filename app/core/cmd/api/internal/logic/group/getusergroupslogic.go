package group

import (
	"api/app/common/ctxtool"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"net/http"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user joined group
func NewGetUserGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupsLogic {
	return &GetUserGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserGroupsLogic) GetUserGroups(req *types.GetUserGroupReq) (resp *types.GetUserGroupResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.GroupService.GetUserGroups(l.ctx, &core.GetUserGroupReq{
		UserId:   uint32(userID),
		Page:     uint32(req.Page),
		Limit:    uint32(req.Limit),
		LatestId: uint32(req.LatestID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	userGroups := make([]types.GroupInfo, 0)
	for _, group := range rpcResp.GroupInfo {
		userGroups = append(userGroups, types.GroupInfo{
			ID:        uint(group.Id),
			Uuid:      group.Uuid,
			Name:      group.Name,
			Avatar:    group.Avatar,
			Desc:      group.Desc,
			CreatedAt: uint(group.CreatedAt),
		})
	}
	return &types.GetUserGroupResp{
		Code:   uint(http.StatusOK),
		Groups: userGroups,
	}, nil
}

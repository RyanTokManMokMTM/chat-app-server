package group

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountUserGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Count user group
func NewCountUserGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountUserGroupsLogic {
	return &CountUserGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CountUserGroupsLogic) CountUserGroups(req *types.CountUserGroupsReq) (resp *types.CountUserGroupsResp, err error) {
	// todo: add your logic here and delete this line

	return
}

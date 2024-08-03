package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStorySeenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStorySeenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStorySeenLogic {
	return &UpdateStorySeenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStorySeenLogic) UpdateStorySeen(in *core.UpdateStorySeenReq) (*core.UpdateStorySeenResp, error) {
	// todo: add your logic here and delete this line

	return &core.UpdateStorySeenResp{}, nil
}

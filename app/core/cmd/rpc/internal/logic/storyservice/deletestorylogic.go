package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoryLogic {
	return &DeleteStoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStoryLogic) DeleteStory(in *core.DeleteStoryReq) (*core.DeleteStoryResp, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteStoryResp{}, nil
}

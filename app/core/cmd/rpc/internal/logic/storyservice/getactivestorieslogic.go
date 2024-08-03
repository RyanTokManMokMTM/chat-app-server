package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActiveStoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActiveStoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActiveStoriesLogic {
	return &GetActiveStoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActiveStoriesLogic) GetActiveStories(in *core.GetActiveStoryReq) (*core.GetActiveStoryResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetActiveStoryResp{}, nil
}

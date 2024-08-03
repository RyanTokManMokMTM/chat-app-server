package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorySeenListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStorySeenListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorySeenListInfoLogic {
	return &GetStorySeenListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStorySeenListInfoLogic) GetStorySeenListInfo(in *core.GetStorySeenListReq) (*core.GetStorySeenListResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetStorySeenListResp{}, nil
}

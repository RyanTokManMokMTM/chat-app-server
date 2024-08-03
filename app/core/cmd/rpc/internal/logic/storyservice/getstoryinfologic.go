package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoryInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoryInfoLogic {
	return &GetStoryInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStoryInfoLogic) GetStoryInfo(in *core.GetStoryInfoByIdRep) (*core.GetStoryInfoByIdResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetStoryInfoByIdResp{}, nil
}

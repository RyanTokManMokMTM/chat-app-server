package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStoriesByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserStoriesByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStoriesByUserIdLogic {
	return &GetUserStoriesByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserStoriesByUserIdLogic) GetUserStoriesByUserId(in *core.GetUserStoryReq) (*core.GetUserStoryResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserStoryResp{}, nil
}

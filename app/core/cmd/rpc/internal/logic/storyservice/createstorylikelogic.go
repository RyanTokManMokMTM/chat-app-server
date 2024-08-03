package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStoryLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStoryLikeLogic {
	return &CreateStoryLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStoryLikeLogic) CreateStoryLike(in *core.CreateStoryLikeReq) (*core.CreateStoryLikeResp, error) {
	// todo: add your logic here and delete this line

	return &core.CreateStoryLikeResp{}, nil
}

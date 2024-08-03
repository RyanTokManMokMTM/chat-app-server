package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoryLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoryLikeLogic {
	return &DeleteStoryLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStoryLikeLogic) DeleteStoryLike(in *core.DeleteStoryReq) (*core.DeleteStoryLikeResp, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteStoryLikeResp{}, nil
}

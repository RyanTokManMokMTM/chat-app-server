package story

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStoryLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create a story like
func NewCreateStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStoryLikeLogic {
	return &CreateStoryLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStoryLikeLogic) CreateStoryLike(req *types.CreateStoryLikeReq) (resp *types.CreateStoryLikeResp, err error) {
	// todo: add your logic here and delete this line

	return
}

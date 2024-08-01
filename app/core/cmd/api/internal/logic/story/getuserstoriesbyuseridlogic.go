package story

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStoriesByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get owner stories
func NewGetUserStoriesByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStoriesByUserIdLogic {
	return &GetUserStoriesByUserIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserStoriesByUserIdLogic) GetUserStoriesByUserId(req *types.GetUserStoryReq) (resp *types.GetUserStoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}

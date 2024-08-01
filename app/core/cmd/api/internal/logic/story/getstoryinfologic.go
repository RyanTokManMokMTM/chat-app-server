package story

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoryInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get an instance story by storyID
func NewGetStoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoryInfoLogic {
	return &GetStoryInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStoryInfoLogic) GetStoryInfo(req *types.GetStoryInfoByIdRep) (resp *types.GetStoryInfoByIdResp, err error) {
	// todo: add your logic here and delete this line

	return
}

package story

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorySeenListInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get the seen user list of instance story by storyID
func NewGetStorySeenListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorySeenListInfoLogic {
	return &GetStorySeenListInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStorySeenListInfoLogic) GetStorySeenListInfo(req *types.GetStorySeenListReq) (resp *types.GetStorySeenListResp, err error) {
	// todo: add your logic here and delete this line

	return
}

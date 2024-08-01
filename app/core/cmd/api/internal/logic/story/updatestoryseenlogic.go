package story

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStorySeenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update story Id which is latest seen
func NewUpdateStorySeenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStorySeenLogic {
	return &UpdateStorySeenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStorySeenLogic) UpdateStorySeen(req *types.UpdateStorySeenReq) (resp *types.UpdateStorySeenResp, err error) {
	// todo: add your logic here and delete this line

	return
}

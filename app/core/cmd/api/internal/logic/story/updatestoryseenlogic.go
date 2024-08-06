package story

import (
	"api/app/common/ctxtool"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.StoryService.UpdateStorySeen(l.ctx, &core.UpdateStorySeenReq{
		UserId:   uint32(userID),
		FriendId: uint32(req.FriendId),
		StoryId:  uint32(req.StoryId),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, rpcErr
	}

	return &types.UpdateStorySeenResp{
		Code: uint(rpcResp.Code),
	}, nil
}

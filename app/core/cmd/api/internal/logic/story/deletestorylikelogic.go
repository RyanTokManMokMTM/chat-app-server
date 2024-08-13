package story

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoryLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete a story like
func NewDeleteStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoryLikeLogic {
	return &DeleteStoryLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStoryLikeLogic) DeleteStoryLike(req *types.DeleteStoryLikeReq) (resp *types.DeleteStoryLikeResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.StoryService.DeleteStoryLike(l.ctx, &core.DeleteStoryLikeReq{
		UserId:  uint32(userID),
		StoryId: uint32(req.StoryId),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}
	return &types.DeleteStoryLikeResp{
		Code: uint(rpcResp.Code),
	}, nil
}

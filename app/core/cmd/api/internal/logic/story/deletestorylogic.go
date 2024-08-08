package story

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete an existing story
func NewDeleteStoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoryLogic {
	return &DeleteStoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStoryLogic) DeleteStory(req *types.DeleteStoryReq) (resp *types.DeleteStoryResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.StoryService.DeleteStory(l.ctx, &core.DeleteStoryReq{
		UserId:  uint32(userID),
		StoryId: uint32(req.StoryID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, rpcErr
	}
	return &types.DeleteStoryResp{
		Code: uint(rpcResp.Code),
	}, nil
}

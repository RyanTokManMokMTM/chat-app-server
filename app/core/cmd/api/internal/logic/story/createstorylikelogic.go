package story

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.StoryService.CreateStoryLike(l.ctx, &core.CreateStoryLikeReq{
		UserId:  uint32(userID),
		StoryId: uint32(req.StoryId),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, rpcErr
	}

	return &types.CreateStoryLikeResp{
		Code: uint(rpcResp.Code),
	}, nil
}

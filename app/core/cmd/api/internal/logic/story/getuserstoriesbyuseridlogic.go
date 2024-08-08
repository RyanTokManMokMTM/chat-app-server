package story

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.StoryService.GetUserStoriesByUserId(l.ctx, &core.GetUserStoryReq{
		UserId:           uint32(userID),
		StoryCreatedTime: uint32(req.StoryCreatedTime),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, rpcErr
	}

	storiesList := make([]types.StoryInfo, 0)
	for _, story := range rpcResp.Stories {
		storiesList = append(storiesList, types.StoryInfo{
			StoryID:       uint(story.StoryId),
			StoryUUID:     story.StoryUUID,
			StoryMediaURL: story.StoryURL,
		})
	}

	return &types.GetUserStoryResp{
		Code:        uint(rpcResp.Code),
		Stories:     storiesList,
		LastStoryId: uint(rpcResp.LastStoryId),
	}, nil
}

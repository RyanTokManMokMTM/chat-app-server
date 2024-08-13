package story

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetActiveStoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get friends active story
func NewGetActiveStoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActiveStoriesLogic {
	return &GetActiveStoriesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetActiveStoriesLogic) GetActiveStories(req *types.GetActiveStoryReq) (resp *types.GetActiveStoryResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.StoryService.GetActiveStories(l.ctx, &core.GetActiveStoryReq{
		UserId:   uint32(userID),
		Page:     uint32(req.Page),
		Limit:    uint32(req.Limit),
		LatestID: uint32(req.LatestID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	activeStories := make([]types.FriendStroy, 0)
	for _, story := range rpcResp.FriendStories {
		activeStories = append(activeStories, types.FriendStroy{
			UserID:               uint(story.UserId),
			Uuid:                 story.Uuid,
			UserName:             story.Username,
			UserAvatar:           story.Avatar,
			IsSeen:               story.IsSeen,
			LatestStoryTimeStamp: uint(story.LatestStoryTimeStamp),
		})
	}

	return &types.GetActiveStoryResp{
		Code:             uint(rpcResp.Code),
		FriendStroies:    activeStories,
		CurrentStoryTime: uint(rpcResp.CurrentStoryTime),
		PageableInfo: types.PageableInfo{
			TotalPage: uint(rpcResp.PageInfo.TotalPage),
			Page:      req.Page,
		},
	}, nil
}

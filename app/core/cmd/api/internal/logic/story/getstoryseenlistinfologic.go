package story

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.StoryService.GetStorySeenListInfo(l.ctx, &core.GetStorySeenListReq{
		UserId:  uint32(userID),
		StoryId: uint32(req.StoryId),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	seenInfoList := make([]types.StorySeenInfo, 0)
	for _, info := range rpcResp.SeenList {
		seenInfoList = append(seenInfoList, types.StorySeenInfo{
			UserID:     uint(info.UserId),
			Uuid:       info.Uuid,
			UserName:   info.Username,
			UserAvatar: info.Avatar,
			IsLiked:    info.IsLiked,
			CreatedAt:  uint(info.CreatedAt),
		})
	}
	return &types.GetStorySeenListResp{
		Code:      uint(http.StatusOK),
		SeenList:  seenInfoList,
		TotalSeen: uint(rpcResp.TotalSeen),
	}, nil
}

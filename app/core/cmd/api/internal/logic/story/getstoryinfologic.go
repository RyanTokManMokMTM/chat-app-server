package story

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	rpcResp, rpcErr := l.svcCtx.StoryService.GetStoryInfo(l.ctx, &core.GetStoryInfoByIdRep{
		UserId:  uint32(userID),
		StoryId: uint32(req.StoryID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	seenList := make([]types.StorySeenUserBasicInfo, 0)
	for _, seenInfo := range rpcResp.StorySeenList {
		seenList = append(seenList, types.StorySeenUserBasicInfo{
			Id:     uint(seenInfo.Id),
			Avatar: seenInfo.Avatar,
		})
	}

	return &types.GetStoryInfoByIdResp{
		Code: uint(rpcResp.Code),
		StoryInfo: types.StoryInfo{
			StoryID:       uint(rpcResp.Info.StoryId),
			StoryUUID:     rpcResp.Info.StoryUUID,
			StoryMediaURL: rpcResp.Info.StoryURL,
		},
		IsLiked:       rpcResp.IsLiked,
		CreateAt:      uint(rpcResp.CreatedAt),
		StorySeenList: seenList,
	}, nil
}

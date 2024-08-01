package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorySeenListInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

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
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.DAO.FindOneStory(l.ctx, req.StoryId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.STORY_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	storyList, err := l.svcCtx.DAO.GetStorySeenUserList(l.ctx, req.StoryId, 20)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	count, err := l.svcCtx.DAO.CountOneStorySeen(l.ctx, req.StoryId)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var seenList = make([]types.StorySeenInfo, 0)
	for _, seen := range storyList {
		likes, err := l.svcCtx.DAO.FindOneUserStoryLike(l.ctx, seen.UserId, req.StoryId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

		var isLikes = false
		if likes != nil {
			isLikes = true
		}
		seenList = append(seenList, types.StorySeenInfo{
			UserID:     seen.UserInfo.Id,
			Uuid:       seen.UserInfo.Uuid,
			UserAvatar: seen.UserInfo.Avatar,
			UserName:   seen.UserInfo.NickName,
			IsLiked:    isLikes,
			CreatedAt:  uint(seen.CreatedAt.Unix()),
		})
	}
	return &types.GetStorySeenListResp{
		Code:      uint(http.StatusOK),
		SeenList:  seenList,
		TotalSeen: uint(count),
	}, nil
}

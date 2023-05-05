package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"gorm.io/gorm"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActiveStoriesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

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
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//friend
	list, err := l.svcCtx.DAO.GetUserFriendList(l.ctx, userID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	activeStories := make([]types.FriendStroy, 0)
	for _, fd := range list {
		stories, err := l.svcCtx.DAO.GetUserStories(l.ctx, fd.FriendID)
		if err != nil {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
		if len(stories) == 0 {
			continue
		}

		activeStories = append(activeStories, types.FriendStroy{
			UserID:     fd.FriendInfo.ID,
			Uuid:       fd.FriendInfo.Uuid,
			UserName:   fd.FriendInfo.NickName,
			UserAvatar: fd.FriendInfo.Avatar,
			IsSeen:     false, //TODO: set as false for testing
			StoriesIDs: stories,
		})

	}

	return &types.GetActiveStoryResp{
		Code:          uint(http.StatusOK),
		FriendStroies: activeStories,
	}, nil
}

package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/pagerx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"gorm.io/gorm"
	"net/http"
	"time"

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

	var storyTimeStamp = int64(req.StoryCreatedTime)
	if storyTimeStamp == 0 {
		storyTimeStamp = time.Now().Unix()
	}

	//friend
	total, err := l.svcCtx.DAO.CountActiveStoryByTimeStamp(l.ctx, userID, storyTimeStamp)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	logx.Info(total)
	pageLimit := pagerx.GetLimit(req.Limit)                            //page limit
	totalPage := pagerx.GetTotalPageByPageSize(uint(total), pageLimit) //HOW MANY PAGE
	pageOffset := pagerx.PageOffset(pageLimit, req.Page)               //TO WHICH PAGE
	logx.Info(pageLimit, totalPage, pageOffset)

	stories, err := l.svcCtx.DAO.GetFriendActiveStoriesByTimeStamp(l.ctx, userID, int(pageOffset), int(pageLimit), storyTimeStamp)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	activeStories := make([]types.FriendStroy, 0)
	for _, fd := range stories {
		//Get story seen record...
		seenStory, err := l.svcCtx.DAO.FindOneUserStorySeen(l.ctx, userID, fd.UserId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

		userStories, err := l.svcCtx.DAO.GetUserStories(l.ctx, fd.UserId)
		if err != nil {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

		var isSeen = false
		if seenStory != nil {
			isSeen = seenStory.StoryInfo.Id == userStories[len(userStories)-1]
		}
		activeStories = append(activeStories, types.FriendStroy{
			UserID:               fd.UserInfo.Id,
			Uuid:                 fd.UserInfo.Uuid,
			UserName:             fd.UserInfo.NickName,
			UserAvatar:           fd.UserInfo.Avatar,
			IsSeen:               isSeen,
			LatestStoryTimeStamp: uint(fd.LatestTime.Unix()),
		})

	}

	return &types.GetActiveStoryResp{
		Code:             uint(http.StatusOK),
		FriendStroies:    activeStories,
		CurrentStoryTime: uint(storyTimeStamp),
		PageableInfo: types.PageableInfo{
			TotalPage: totalPage,
			Page:      req.Page,
		},
	}, nil
}

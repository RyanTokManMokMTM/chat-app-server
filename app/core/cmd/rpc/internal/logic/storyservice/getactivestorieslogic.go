package storyservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/pagerx"
	"gorm.io/gorm"
	"time"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetActiveStoriesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetActiveStoriesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetActiveStoriesLogic {
	return &GetActiveStoriesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetActiveStoriesLogic) GetActiveStories(in *core.GetActiveStoryReq) (*core.GetActiveStoryResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "uer not exist, error %+v", err)
		}
		return nil, err
	}

	var storyTimeStamp = int64(in.StoryCreatedTime)
	if storyTimeStamp == 0 {
		storyTimeStamp = time.Now().Unix()
	}

	//friend
	total, err := l.svcCtx.DAO.CountActiveStoryByTimeStamp(l.ctx, userID, storyTimeStamp)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	pageLimit := pagerx.GetLimit(uint(in.Limit))                       //page limit
	totalPage := pagerx.GetTotalPageByPageSize(uint(total), pageLimit) //HOW MANY PAGE
	pageOffset := pagerx.PageOffset(pageLimit, uint(in.Page))          //TO WHICH PAGE
	logx.Info(pageLimit, totalPage, pageOffset)

	stories, err := l.svcCtx.DAO.GetFriendActiveStoriesByTimeStamp(l.ctx, userID, int(pageOffset), int(pageLimit), storyTimeStamp)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	activeStories := make([]*core.FriendStory, 0)
	for _, story := range stories {
		//Get story seen record...
		lastStorySeen, err := l.svcCtx.DAO.FindOneLatestUserStorySeen(l.ctx, userID, story.UserId)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}

		userStories, err := l.svcCtx.DAO.GetUserStories(l.ctx, story.UserId)
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}

		var isSeen = false
		if lastStorySeen != nil {
			logx.Infof("latest story id %d", lastStorySeen.StoryId)
			isSeen = lastStorySeen.StoryInfo.Id == userStories[len(userStories)-1]
		}
		activeStories = append(activeStories, &core.FriendStory{
			UserId:               uint32(story.UserInfo.Id),
			Uuid:                 story.UserInfo.Uuid,
			Username:             story.UserInfo.NickName,
			Avatar:               story.UserInfo.Avatar,
			IsSeen:               isSeen,
			LatestStoryTimeStamp: uint32(uint(story.LatestTime.Unix())),
		})

	}

	return &core.GetActiveStoryResp{
		Code:             uint32(errx.SUCCESS),
		FriendStories:    activeStories,
		CurrentStoryTime: uint32(uint(storyTimeStamp)),
		PageInfo: &core.PageableInfo{
			TotalPage: uint32(totalPage),
			Page:      in.Page,
		},
	}, nil
}

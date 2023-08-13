package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/helper/toolx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/pagerx"
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
	total, err := l.svcCtx.DAO.CountActiveStory(l.ctx, userID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	logx.Info(total)
	pageLimit := pagerx.GetLimit(req.Limit)                            //page limit
	totalPage := pagerx.GetTotalPageByPageSize(uint(total), pageLimit) //HOW MANY PAGE
	pageOffset := pagerx.PageOffset(pageLimit, req.Page)               //TO WHICH PAGE
	logx.Info(pageLimit, totalPage, pageOffset)

	stories, err := l.svcCtx.DAO.GetFriendActiveStories(l.ctx, userID, int(pageOffset), int(pageLimit))
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	activeStories := make([]types.FriendStroy, 0)
	for _, fd := range stories {
		ids, err := toolx.StringTouIntArray(fd.Ids)
		if err != nil {
			return nil, errx.NewCustomErrCode(errx.DB_ERROR)
		}
		activeStories = append(activeStories, types.FriendStroy{
			UserID:     fd.UserInfo.Id,
			Uuid:       fd.UserInfo.Uuid,
			UserName:   fd.UserInfo.NickName,
			UserAvatar: fd.UserInfo.Avatar,
			IsSeen:     false, //for testing...
			StoriesIDs: ids,
		})

	}

	return &types.GetActiveStoryResp{
		Code:          uint(http.StatusOK),
		FriendStroies: activeStories,
		PageableInfo: types.PageableInfo{
			TotalPage: totalPage,
			Page:      req.Page,
		},
	}, nil
}

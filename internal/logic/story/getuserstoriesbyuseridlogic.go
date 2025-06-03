package story

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStoriesByuserIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserStoriesByuserIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStoriesByuserIDLogic {
	return &GetUserStoriesByuserIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserStoriesByuserIDLogic) GetUserStoriesByuserID(req *types.GetUserStoryReq) (resp *types.GetUserStoryResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userID)
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

	storys, err := l.svcCtx.Uow.StoryRepo().FindAllUserStoriesByTimeStamp(l.ctx, req.UserID, storyTimeStamp)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var lastStoryID uint = 0
	seenStory, err := l.svcCtx.Uow.UserStorySeenRepo().FindLatestOne(l.ctx, userID, req.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if seenStory != nil {
		lastStoryID = seenStory.StoryID
	}

	var storiesList = make([]types.StoryInfo, 0)
	for _, s := range storys {
		storiesList = append(storiesList, types.StoryInfo{
			StoryID:       s.ID,
			StoryUUID:     s.UUID,
			StoryMediaURL: s.StoryMediaPath,
		})
	}
	return &types.GetUserStoryResp{
		Code:        uint(http.StatusOK),
		Stories:     storiesList,
		LastStoryID: lastStoryID,
	}, nil
}

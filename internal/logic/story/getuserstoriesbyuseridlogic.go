package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStoriesByUserIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

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

	storys, err := l.svcCtx.DAO.GetUserStoriesByTimeStamp(l.ctx, req.UserID, storyTimeStamp)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var lastStoryID uint = 0
	seenStory, err := l.svcCtx.DAO.FindOneLatestUserStorySeen(l.ctx, userID, req.UserID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if seenStory != nil {
		lastStoryID = seenStory.StoryId
	}

	var storiesList = make([]types.StoryInfo, 0)
	for _, s := range storys {
		storiesList = append(storiesList, types.StoryInfo{
			StoryID:       s.Id,
			StoryUUID:     s.Uuid.String(),
			StoryMediaURL: s.StoryMediaPath,
		})
	}
	return &types.GetUserStoryResp{
		Code:        uint(http.StatusOK),
		Stories:     storiesList,
		LastStoryId: lastStoryID,
	}, nil
}

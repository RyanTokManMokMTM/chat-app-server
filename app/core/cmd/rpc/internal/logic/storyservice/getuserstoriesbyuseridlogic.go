package storyservicelogic

import (
	"api/app/common/errx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"time"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStoriesByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserStoriesByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStoriesByUserIdLogic {
	return &GetUserStoriesByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserStoriesByUserIdLogic) GetUserStoriesByUserId(in *core.GetUserStoryReq) (*core.GetUserStoryResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}

	var storyTimeStamp = int64(in.StoryCreatedTime)
	if storyTimeStamp == 0 {
		storyTimeStamp = time.Now().Unix()
	}

	stories, err := l.svcCtx.DAO.GetUserStoriesByTimeStamp(l.ctx, uint(in.UserId), storyTimeStamp)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	var lastStoryID uint = 0
	seenStory, err := l.svcCtx.DAO.FindOneLatestUserStorySeen(l.ctx, userID, uint(in.UserId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	if seenStory != nil {
		lastStoryID = seenStory.StoryId
	}

	var storiesList = make([]*core.StoryInfo, 0)
	for _, s := range stories {
		storiesList = append(storiesList, &core.StoryInfo{
			StoryId:   uint32(s.Id),
			StoryUUID: s.Uuid.String(),
			StoryURL:  s.StoryMediaPath,
		})
	}

	return &core.GetUserStoryResp{
		Code:        uint32(errx.SUCCESS),
		Stories:     storiesList,
		LastStoryId: uint32(lastStoryID),
	}, nil
}

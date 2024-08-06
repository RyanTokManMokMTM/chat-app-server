package storyservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStorySeenListInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStorySeenListInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStorySeenListInfoLogic {
	return &GetStorySeenListInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStorySeenListInfoLogic) GetStorySeenListInfo(in *core.GetStorySeenListReq) (*core.GetStorySeenListResp, error) {
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

	_, err = l.svcCtx.DAO.FindOneStory(l.ctx, uint(in.StoryId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STORY_NOT_EXIST), "story not exist")
		}
		return nil, err
	}

	storyList, err := l.svcCtx.DAO.GetStorySeenUserList(l.ctx, uint(in.StoryId), 20)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	count, err := l.svcCtx.DAO.CountOneStorySeen(l.ctx, uint(in.StoryId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	var seenList = make([]*core.StorySeenInfo, 0)
	for _, seen := range storyList {
		likes, err := l.svcCtx.DAO.FindOneUserStoryLike(l.ctx, seen.UserId, uint(in.StoryId))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}

		var isLikes = false
		if likes != nil {
			isLikes = true
		}
		seenList = append(seenList, &core.StorySeenInfo{
			UserId:    uint32(seen.UserInfo.Id),
			Uuid:      seen.UserInfo.Uuid,
			Avatar:    seen.UserInfo.Avatar,
			Username:  seen.UserInfo.NickName,
			IsLiked:   isLikes,
			CreatedAt: uint32(uint(seen.CreatedAt.Unix())),
		})
	}

	return &core.GetStorySeenListResp{
		Code:      uint32(errx.SUCCESS),
		SeenList:  seenList,
		TotalSeen: uint32(uint(count)),
	}, nil
}

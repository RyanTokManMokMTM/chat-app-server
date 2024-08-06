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

type GetStoryInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoryInfoLogic {
	return &GetStoryInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStoryInfoLogic) GetStoryInfo(in *core.GetStoryInfoByIdRep) (*core.GetStoryInfoByIdResp, error) {
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

	story, err := l.svcCtx.DAO.FindOneStory(l.ctx, uint(in.StoryId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STORY_NOT_EXIST), "story not exist, error : %+v", err)
		}
		return nil, err
	}

	//If current story is no belong to ourselves -> am I liked?
	//If current story is mine -> any one liked

	var isLiked = false
	var seenUserList []*core.StorySeenUserBasicInfo = nil
	if story.UserId == userID {
		count, err := l.svcCtx.DAO.CountStoryLikes(l.ctx, uint(in.StoryId))
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}
		isLiked = count > 0

		users, err := l.svcCtx.DAO.GetStorySeenUserList(l.ctx, uint(in.StoryId), 3) // MARK: the latest 3 user
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}

		for _, u := range users {
			seenUserList = append(seenUserList, &core.StorySeenUserBasicInfo{
				Id:     uint32(u.UserInfo.Id),
				Avatar: u.UserInfo.Avatar,
			})
		}
	} else {
		userLiked, err := l.svcCtx.DAO.FindOneUserStoryLike(l.ctx, userID, uint(in.StoryId))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}

		if userLiked != nil {
			isLiked = true
		}

	}

	return &core.GetStoryInfoByIdResp{
		Code: uint32(errx.SUCCESS),
		Info: &core.StoryInfo{
			StoryId:   uint32(story.Id),
			StoryUUID: story.Uuid.String(),
			StoryURL:  story.StoryMediaPath,
		},
		IsLiked:       isLiked,
		CreatedAt:     uint32(story.CreatedAt.Unix()),
		StorySeenList: seenUserList,
	}, nil
}

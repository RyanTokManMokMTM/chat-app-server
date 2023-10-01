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

type GetStoryInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoryInfoLogic {
	return &GetStoryInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStoryInfoLogic) GetStoryInfo(req *types.GetStoryInfoByIdRep) (resp *types.GetStoryInfoByIdResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	story, err := l.svcCtx.DAO.FindOneStory(l.ctx, req.StoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.STORY_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//If current story is no belong to ourselves -> am I liked?
	//If current story is mine -> any one liked

	var isLiked = false
	var seenUserList []types.StorySeenUserBasicInfo = nil
	if story.UserId == userID {
		count, err := l.svcCtx.DAO.CountStoryLikes(l.ctx, req.StoryID)
		if err != nil {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
		isLiked = count > 0

		users, err := l.svcCtx.DAO.GetStorySeenUserList(l.ctx, req.StoryID, 3) // MARK: the latest 3 user
		if err != nil {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

		for _, u := range users {
			seenUserList = append(seenUserList, types.StorySeenUserBasicInfo{
				Id:     u.UserInfo.Id,
				Avatar: u.UserInfo.Avatar,
			})
		}
	} else {
		userLiked, err := l.svcCtx.DAO.FindOneUserStoryLike(l.ctx, userID, req.StoryID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

		if userLiked != nil {
			isLiked = true
		}

	}

	return &types.GetStoryInfoByIdResp{
		Code:          uint(http.StatusOK),
		StoryID:       story.Id,
		StoryMediaURL: story.StoryMediaPath,
		IsLiked:       isLiked,
		CreateAt:      uint(story.CreatedAt.Unix()),
		StorySeenList: seenUserList,
	}, nil
}

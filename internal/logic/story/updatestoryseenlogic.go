package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStorySeenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateStorySeenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStorySeenLogic {
	return &UpdateStorySeenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStorySeenLogic) UpdateStorySeen(req *types.UpdateStorySeenReq) (resp *types.UpdateStorySeenResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	storySeen, err := l.svcCtx.DAO.FindOneUserStorySeen(l.ctx, userID, req.FriendId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := l.svcCtx.DAO.InsertOneUserStorySeen(l.ctx, userID, req.FriendId, req.StoryId); err != nil {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

	} else {
		currentStory, err := l.svcCtx.DAO.FindOneUserStory(l.ctx, req.StoryId, req.FriendId)
		if err != nil {
			//this story should be existed
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}

		if currentStory.CreatedAt.After(storySeen.StoryInfo.CreatedAt) {
			if err := l.svcCtx.DAO.UpdateOneUserStorySeen(l.ctx, storySeen.ID, req.StoryId); err != nil {
				return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
			}
		}

	}

	return &types.UpdateStorySeenResp{
		Code: uint(http.StatusOK),
	}, nil
}

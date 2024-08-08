package storyservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoryLogic {
	return &DeleteStoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStoryLogic) DeleteStory(in *core.DeleteStoryReq) (*core.DeleteStoryResp, error) {
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

	//TODO： Get story if the story is available and belongs to the user
	_, err = l.svcCtx.DAO.FindOneUserStory(l.ctx, uint(in.StoryId), userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STORY_NOT_EXIST), "story not exist, error : %+v", err)
		}
		return nil, err
	}

	if err = l.svcCtx.DAO.DeleteStories(l.ctx, uint(in.StoryId)); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.DeleteStoryResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

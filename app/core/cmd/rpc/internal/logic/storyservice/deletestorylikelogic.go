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

type DeleteStoryLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoryLikeLogic {
	return &DeleteStoryLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStoryLikeLogic) DeleteStoryLike(in *core.DeleteStoryReq) (*core.DeleteStoryLikeResp, error) {
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

	likes, err := l.svcCtx.DAO.FindOneUserStoryLike(l.ctx, userID, uint(in.StoryId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	err = l.svcCtx.DAO.DeleteOneUserStoryLike(l.ctx, likes.ID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.DeleteStoryLikeResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

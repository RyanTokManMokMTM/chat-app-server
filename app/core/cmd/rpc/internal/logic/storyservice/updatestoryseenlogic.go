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

type UpdateStorySeenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStorySeenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStorySeenLogic {
	return &UpdateStorySeenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStorySeenLogic) UpdateStorySeen(in *core.UpdateStorySeenReq) (*core.UpdateStorySeenResp, error) {
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

	_, err = l.svcCtx.DAO.FindOneUserStorySeen(l.ctx, userID, uint(in.FriendId), uint(in.StoryId))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := l.svcCtx.DAO.InsertOneUserStorySeen(l.ctx, userID, uint(in.FriendId), uint(in.StoryId)); err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}
	}

	return &core.UpdateStorySeenResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

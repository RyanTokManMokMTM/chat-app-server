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

type CreateStoryLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStoryLikeLogic {
	return &CreateStoryLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStoryLikeLogic) CreateStoryLike(in *core.CreateStoryLikeReq) (*core.CreateStoryLikeResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error :%+v", err)
		}
		return nil, err
	}

	_, err = l.svcCtx.DAO.FindOneStory(l.ctx, uint(in.StoryId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STORY_NOT_EXIST), "story not exist, error : %+v", err)
		}
		return nil, err
	}

	err = l.svcCtx.DAO.InsertOneUserStoryLike(l.ctx, userID, uint(in.StoryId))
	if err != nil {
		return nil, err
	}

	return &core.CreateStoryLikeResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

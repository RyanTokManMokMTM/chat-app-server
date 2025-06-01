package story

import (
	"context"
	"errors"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStoryLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStoryLikeLogic {
	return &CreateStoryLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStoryLikeLogic) CreateStoryLike(req *types.CreateStoryLikeReq) (resp *types.CreateStoryLikeResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.Uow.StoryRepo().FindOneByID(l.ctx, req.StoryId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.STORY_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.Uow.UserStoryLikesRepo().CreateOne(l.ctx, models.UserStoryLikes{
		UserId:  userId,
		StoryId: req.StoryId,
	})
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	return &types.CreateStoryLikeResp{
		Code: uint(http.StatusOK),
	}, nil
}

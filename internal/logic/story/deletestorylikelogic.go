package story

import (
	"context"
	"errors"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStoryLikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteStoryLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStoryLikeLogic {
	return &DeleteStoryLikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteStoryLikeLogic) DeleteStoryLike(req *types.DeleteStoryLikeReq) (resp *types.DeleteStoryLikeResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	likes, err := l.svcCtx.Uow.UserStoryLikesRepo().FindOneByUserIdAndStoryId(l.ctx, userId, req.StoryId)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	err = l.svcCtx.Uow.UserStoryLikesRepo().DeleteOne(l.ctx, likes.ID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	return &types.DeleteStoryLikeResp{
		Code: uint(http.StatusOK),
	}, nil
}

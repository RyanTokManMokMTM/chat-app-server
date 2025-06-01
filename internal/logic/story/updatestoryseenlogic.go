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
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.Uow.UserStorySeenRepo().FindOneByUserIdAndFriendIdAndStroyId(l.ctx, userId, req.FriendId, req.StoryId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		_, err := l.svcCtx.Uow.UserStorySeenRepo().CreateOne(l.ctx, models.UserStorySeen{
			UserId:   userId,
			FriendId: req.FriendId,
			StoryId:  req.StoryId,
		})
		if err != nil {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
	}

	return &types.UpdateStorySeenResp{
		Code: uint(http.StatusOK),
	}, nil
}

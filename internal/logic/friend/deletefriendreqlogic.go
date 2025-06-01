package friend

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

type DeleteFriendReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFriendReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendReqLogic {
	return &DeleteFriendReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFriendReqLogic) DeleteFriendReq(req *types.DeleteFriendReq) (resp *types.DeleteFriendResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: Check is friend
	_, err = l.svcCtx.Uow.UserFriendsRepo().FindOneByUserIdAndFriendId(l.ctx, userId, req.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.NOT_YET_FRIEND)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: Break the friend relationship
	err = l.svcCtx.Uow.UserFriendsRepo().DeleteOneByUserIdAndFriendID(l.ctx, userId, req.UserId)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	return &types.DeleteFriendResp{
		Code: uint(http.StatusOK),
	}, nil
}

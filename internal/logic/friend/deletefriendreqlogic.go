package friend

import (
	"context"
	"errors"
	"github.com/ryantokmanmok/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmok/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	"github.com/ryantokmanmok/chat-app-server/internal/types"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: Check is friend
	ok, err := l.svcCtx.DAO.FindOneFriend(l.ctx, req.UserID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if !ok {
		return nil, errx.NewCustomErrCode(errx.NOT_YET_FRIEND)
	}

	//TODO: Break the friend relationship
	err = l.svcCtx.DAO.DeleteOneFriend(l.ctx, req.UserID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	return &types.DeleteFriendResp{
		Code: uint(http.StatusOK),
	}, nil
}

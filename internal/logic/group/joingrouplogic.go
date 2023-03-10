package group

import (
	"context"
	"errors"
	"github.com/ryantokmanmok/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmok/chat-app-server/common/errx"
	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	"github.com/ryantokmanmok/chat-app-server/internal/types"
	"gorm.io/gorm"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinGroupLogic) JoinGroup(req *types.JoinGroupReq) (resp *types.JoinGroupResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
	}

	_, err = l.svcCtx.DAO.FindOneGroup(l.ctx, req.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	member, err := l.svcCtx.DAO.FindOneGroupMember(l.ctx, req.GroupID, userID)
	if member != nil {
		return nil, errx.NewCustomErrCode(errx.ALREADY_IN_GROUP)
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: Add it to group
	err = l.svcCtx.DAO.InsertOneGroupMember(l.ctx, req.GroupID, userID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &types.JoinGroupResp{
		Code: uint(http.StatusOK),
	}, nil

}

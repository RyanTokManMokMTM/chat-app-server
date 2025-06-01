package group

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler/ws"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LeaveGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLeaveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveGroupLogic {
	return &LeaveGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LeaveGroupLogic) LeaveGroup(req *types.LeaveGroupReq) (resp *types.LeaveGroupResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	u, err := l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	g, err := l.svcCtx.Uow.GroupRepo().FindOneById(l.ctx, req.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.Uow.UserGroupRepo().FindOneByGroupIdAndUserId(l.ctx, req.GroupID, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.NOT_JOIN_GROUP_YET)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.Uow.UserGroupRepo().DeleteOneByGroupIdAndUserId(l.ctx, req.GroupID, userId); err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	go func() {
		logx.Info("sending a system message")
		sysMessage := fmt.Sprintf("%s left the group", u.NickName)
		ws.SendGroupSystemNotification(u.Uuid, g.Uuid, sysMessage)
	}()

	return &types.LeaveGroupResp{
		Code: uint(http.StatusOK),
	}, nil
}

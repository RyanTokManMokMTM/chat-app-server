package group

import (
	"context"
	"errors"
	"fmt"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler/ws"
	"gorm.io/gorm"
	"net/http"

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	u, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	g, err := l.svcCtx.DAO.FindOneGroup(l.ctx, req.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.DAO.FindOneGroupMember(l.ctx, req.GroupID, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.NOT_JOIN_GROUP_YET)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.DAO.DeleteGroupMember(l.ctx, req.GroupID, userID); err != nil {
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

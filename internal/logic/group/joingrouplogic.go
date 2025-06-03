package group

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler/ws"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"gorm.io/gorm"

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
	u, err := l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
	}

	g, err := l.svcCtx.Uow.GroupRepo().FindOneByID(l.ctx, req.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	member, err := l.svcCtx.Uow.UserGroupRepo().FindOneByGroupIDAndUserID(l.ctx, req.GroupID, userID)
	if member != nil {
		return nil, errx.NewCustomErrCode(errx.ALREADY_IN_GROUP)
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: Add it to group
	_, err = l.svcCtx.Uow.UserGroupRepo().CreateOne(l.ctx, &models.UserGroup{
		GroupID: req.GroupID,
		UserID:  userID,
	})
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	go func() {
		logx.Info("sending a system message")
		sysMessage := fmt.Sprintf("%s joined the group", u.NickName)
		ws.SendGroupSystemNotification(u.UUID, g.UUID, sysMessage)
	}()
	return &types.JoinGroupResp{
		Code: uint(http.StatusOK),
	}, nil

}

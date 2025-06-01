package group

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

type DeleteGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteGroupLogic) DeleteGroup(req *types.DeleteGroupReq) (resp *types.DeleteGroupResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: GET GROUP INFO
	group, err := l.svcCtx.Uow.GroupRepo().FindOneById(l.ctx, req.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if group.GroupLead != userId {
		return nil, errx.NewCustomErrCode(errx.NO_GROUP_AUTHORITY)
	}

	// Fixme: Should me in a transaction.

	//TODO: Remove all user/members that joined the group
	if err := l.svcCtx.Uow.UserGroupRepo().DeleteAll(l.ctx, req.GroupID); err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: Remove entire group
	if err := l.svcCtx.Uow.GroupRepo().DeleteOneById(l.ctx, req.GroupID); err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &types.DeleteGroupResp{
		Code: uint(http.StatusOK),
	}, nil
}

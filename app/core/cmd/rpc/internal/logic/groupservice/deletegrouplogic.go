package groupservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteGroupLogic {
	return &DeleteGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteGroupLogic) DeleteGroup(in *core.DeleteGroupReq) (*core.DeleteGroupResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}

	//TODO: GET GROUP INFO
	group, err := l.svcCtx.DAO.FindOneGroup(l.ctx, uint(in.GroupId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.GROUP_NOT_EXIST), "group not exist, error : %+v", err)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if group.GroupLead != userID {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NO_GROUP_AUTHORITY), "user no group authority")
	}

	//TODO: Remove all user/members that joined the group
	if err := l.svcCtx.DAO.DeleteAllGroupMembers(l.ctx, uint(in.GroupId)); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	//TODO: Remove entire group
	if err := l.svcCtx.DAO.DeleteOneGroup(l.ctx, uint(in.GroupId)); err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &core.DeleteGroupResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

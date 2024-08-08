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

type LeaveGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLeaveGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LeaveGroupLogic {
	return &LeaveGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LeaveGroupLogic) LeaveGroup(in *core.LeaveGroupReq) (*core.LeaveGroupResp, error) {
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
	_, err = l.svcCtx.DAO.FindOneGroup(l.ctx, uint(in.GroupId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.GROUP_NOT_EXIST), "group not exist")
		}
		return nil, err
	}

	_, err = l.svcCtx.DAO.FindOneGroupMember(l.ctx, uint(in.GroupId), userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NOT_JOIN_GROUP_YET), "not joined group yet")
		}
		return nil, err
	}

	if err := l.svcCtx.DAO.DeleteGroupMember(l.ctx, uint(in.GroupId), userID); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	//go func() {
	//	logx.Info("sending a system message")
	//	sysMessage := fmt.Sprintf("%s left the group", u.NickName)
	//	ws.SendGroupSystemNotification(u.Uuid, g.Uuid, sysMessage)
	//}()

	return &core.LeaveGroupResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

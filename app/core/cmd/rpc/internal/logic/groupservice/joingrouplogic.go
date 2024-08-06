package groupservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *JoinGroupLogic) JoinGroup(in *core.JoinGroupReq) (*core.JoinGroupResp, error) {
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
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.GROUP_NOT_EXIST), "group not exist, error : %+v", err)
		}
		return nil, err
	}

	member, err := l.svcCtx.DAO.FindOneGroupMember(l.ctx, uint(in.GroupId), userID)
	if member != nil {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.ALREADY_IN_GROUP), "already in group")
	}
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	//TODO: Add it to group
	err = l.svcCtx.DAO.InsertOneGroupMember(l.ctx, uint(in.GroupId), userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}
	//
	//go func() {
	//	logx.Info("sending a system message")
	//	sysMessage := fmt.Sprintf("%s joined the group", u.NickName)
	//	ws.SendGroupSystemNotification(u.Uuid, g.Uuid, sysMessage)
	//}()

	return &core.JoinGroupResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

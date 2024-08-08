package chatservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMessageLogic {
	return &DeleteMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteMessageLogic) DeleteMessage(in *core.DeleteMessageReq) (*core.DeleteMessageResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, with error %+v", err)
		}
		return nil, err
	}

	msg, err := l.svcCtx.DAO.FindOneMessage(l.ctx, uint(in.MsgId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.MESSAGE_NOT_EXIST), "message not exist, error :%+v", err)
		}
		return nil, err
	}

	if msg.FromUserID != userID {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NO_MESSAGE_DELETE_AUTHORITY), "unable to remove the message")
	}

	if err := l.svcCtx.DAO.DeleteOneMessage(l.ctx, uint(in.MsgId)); err != nil {
		return nil, err
	}
	return &core.DeleteMessageResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

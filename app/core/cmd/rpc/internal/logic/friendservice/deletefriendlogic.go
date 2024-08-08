package friendservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteFriendLogic) DeleteFriend(in *core.DeleteFriendReq) (*core.DeleteFriendResp, error) {
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

	//TODO: Find Friend
	_, err = l.svcCtx.DAO.FindOneFriend(l.ctx, userID, uint(in.FriendId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NOT_YET_FRIEND), "user not yet friend, userId : %d", in.FriendId)
		}
		return nil, err
	}

	if err := l.svcCtx.DAO.DeleteOneFriend(l.ctx, userID, uint(in.FriendId)); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.DeleteFriendResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

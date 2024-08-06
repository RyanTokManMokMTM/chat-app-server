package friendservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddFriendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddFriendLogic {
	return &AddFriendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddFriendLogic) AddFriend(in *core.AddFriendReq) (*core.AddFriendResp, error) {
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

	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, uint(in.UserId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist")
		}
		return nil, err
	}

	//TODO: Check is friend
	_, err = l.svcCtx.DAO.FindOneFriend(l.ctx, userID, uint(in.UserId))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			//TODO: Create FriendID Relationship
			err = l.svcCtx.DAO.InsertOneFriend(l.ctx, userID, uint(in.UserId))
			if err != nil {
				logx.WithContext(l.ctx).Error(err)
				return nil, err
			}
			return &core.AddFriendResp{
				Code: uint32(errx.SUCCESS),
			}, nil
		} else {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}
	}

	return nil, errors.Wrapf(errx.NewCustomErrCode(errx.IS_FRIEND_ALREADY), "friended already")
}

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

type GetFriendInformationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInformationLogic {
	return &GetFriendInformationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendInformationLogic) GetFriendInformation(in *core.GetFriendInfoReq) (*core.GetFriendInfoResp, error) {
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

	friend, err := l.svcCtx.DAO.FindOneUserByUUID(l.ctx, in.Uuid)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist")
		}
		return nil, err
	}

	return &core.GetFriendInfoResp{
		Code: uint32(errx.SUCCESS),
		Info: &core.FriendInfo{
			Id:       uint32(friend.Id),
			Uuid:     friend.Uuid,
			NickName: friend.NickName,
			Avatar:   friend.Avatar,
		},
	}, nil
}

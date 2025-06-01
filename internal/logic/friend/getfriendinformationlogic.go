package friend

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

type GetFriendInformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInformationLogic {
	return &GetFriendInformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendInformationLogic) GetFriendInformation(req *types.GetFriendInfoReq) (resp *types.GetFriendInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	friend, err := l.svcCtx.Uow.UserRepo().FindOneUserByUUID(l.ctx, req.Uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &types.GetFriendInfoResp{
		Code: uint(http.StatusOK),
		FriendInfo: types.FriendInfo{
			ID:       friend.Id,
			Uuid:     friend.Uuid,
			NickName: friend.NickName,
			Avatar:   friend.Avatar,
		},
	}, nil
}

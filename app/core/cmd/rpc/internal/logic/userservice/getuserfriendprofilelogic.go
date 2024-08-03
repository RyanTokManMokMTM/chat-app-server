package userservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFriendProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFriendProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFriendProfileLogic {
	return &GetUserFriendProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserFriendProfileLogic) GetUserFriendProfile(in *core.GetUserFriendProfileReq) (*core.GetUserFriendProfileResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserFriendProfileResp{}, nil
}

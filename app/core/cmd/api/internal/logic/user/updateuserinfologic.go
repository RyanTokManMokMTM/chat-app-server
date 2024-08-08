package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update user profile
func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.UpdateUserInfoResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.CoreRPC.UpdateUserInfo(l.ctx, &core.UpdateUserInfoReq{
		UserId: uint32(userID),
		Name:   req.Name,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}
	return &types.UpdateUserInfoResp{
		Code: uint(rpcResp.Code),
	}, nil
}

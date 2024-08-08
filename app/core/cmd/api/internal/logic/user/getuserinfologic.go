package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get User Profile - Other/Own
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	// todo: add your logic here and delete this line

	if req.UUID == "" && req.UserID == 0 {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}
	userId := uint32(req.UserID)
	rpcResp, rpcErr := l.svcCtx.UserService.GetUserInfo(l.ctx, &core.GetUserInfoReq{
		Uuid:   &(req.UUID),
		UserId: &(userId),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}
	return &types.GetUserInfoResp{
		Code:          uint(rpcResp.Code),
		UUID:          rpcResp.Uuid,
		Email:         rpcResp.Email,
		Name:          rpcResp.Name,
		Avatar:        rpcResp.Avatar,
		Cover:         rpcResp.Cover,
		StatusMessage: rpcResp.StatusMessage,
	}, nil
}

package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// User account sign in
func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignInLogic) SignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, rpcErr := l.svcCtx.UserService.SignIn(
		l.ctx,
		&core.SignInReq{
			Email:    req.Email,
			Password: req.Password,
		})
	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.SignInResp{
		Code:        uint(rpcResp.Code),
		Token:       rpcResp.Token,
		ExpiredTime: uint(rpcResp.ExpiredTime),
		UserInfo: types.CommonUserInfo{
			ID:            uint(rpcResp.UserInfo.Id),
			Uuid:          rpcResp.UserInfo.Uuid,
			NickName:      rpcResp.UserInfo.Name,
			Avatar:        rpcResp.UserInfo.Avatar,
			Email:         rpcResp.UserInfo.Email,
			Cover:         rpcResp.UserInfo.Cover,
			StatusMessage: rpcResp.UserInfo.StatusMessage,
		},
	}, nil

}

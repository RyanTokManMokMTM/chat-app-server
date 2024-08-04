package user

import (
	"api/app/core/cmd/rpc/types/core"
	"context"
	"net/http"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

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
	response, rpcErr := l.svcCtx.UserService.SignIn(
		l.ctx,
		&core.SignInReq{
			Email:    req.Email,
			Password: req.Password,
		})
	if rpcErr != nil {
		return nil, rpcErr
	}

	return &types.SignInResp{
		Code:        uint(http.StatusOK),
		Token:       response.Token,
		ExpiredTime: uint(response.ExpiredTime),
		UserInfo: types.CommonUserInfo{
			ID:            uint(response.UserInfo.Id),
			Uuid:          response.UserInfo.Uuid,
			NickName:      response.UserInfo.Name,
			Avatar:        response.UserInfo.Avatar,
			Email:         response.UserInfo.Email,
			Cover:         response.UserInfo.Cover,
			StatusMessage: response.UserInfo.StatusMessage,
		},
	}, nil

}

package user

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// User account sign in
func NewUserSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignInLogic {
	return &UserSignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignInLogic) UserSignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	// todo: add your logic here and delete this line

	return
}

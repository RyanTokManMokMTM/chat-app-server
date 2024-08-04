package user

import (
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// User accout sign up
func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUpLogic) SignUp(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("Call User SignUp API with email : %v, name : %v", req.Email, req.Name)
	rpcResp, rpcErr := l.svcCtx.UserService.SignUp(l.ctx, &core.SignUpReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if rpcErr != nil {
		return nil, rpcErr
	}

	return &types.SignUpResp{
		Code:        http.StatusOK,
		Token:       rpcResp.Token,
		ExpiredTime: uint(rpcResp.ExpiredTime),
	}, nil

}

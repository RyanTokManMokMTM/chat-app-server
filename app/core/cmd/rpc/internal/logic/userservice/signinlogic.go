package userservicelogic

import (
	"api/app/common/cryptox"
	"api/app/common/ctxtool"
	"api/app/common/errx"
	"api/app/common/jwtx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
	"time"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignInLogic {
	return &SignInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignInLogic) SignIn(in *core.SignInReq) (*core.SignInResp, error) {
	// todo: add your logic here and delete this line
	u, err := l.svcCtx.DAO.FindOneUserByEmail(l.ctx, in.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not found error :%v", err)
		}
	}

	encryptedPW := cryptox.PasswordEncrypt(in.Password, l.svcCtx.Config.Salt)
	if strings.Compare(encryptedPW, u.Password) != 0 {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_SIGN_IN_FAILED), "user sign in error :%v", err)
	}

	now := time.Now().Unix()
	exp := now + l.svcCtx.Config.AuthConf.AccessExpire
	payLoad := map[string]interface{}{
		ctxtool.CTXJWTUserID: u.Id,
	}

	token, err := jwtx.GetToken(now, exp, l.svcCtx.Config.AuthConf.AccessSecret, payLoad)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.TOKEN_GENERATE_ERROR), "token generated error :%v", err)
	}

	return &core.SignInResp{
		Code:        int32(errx.SUCCESS),
		Token:       token,
		ExpiredTime: int32(uint(exp)),
		UserInfo: &core.UserInfo{
			Id:            uint32(u.Id),
			Uuid:          u.Uuid,
			Name:          u.NickName,
			Avatar:        u.Avatar,
			Email:         u.Email,
			Cover:         u.Cover,
			StatusMessage: u.StatusMessage,
		},
	}, nil
}

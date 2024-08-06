package userservicelogic

import (
	"api/app/common/cryptox"
	"api/app/common/errx"
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/jwtx"
	"gorm.io/gorm"
	"time"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignUpLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignUpLogic {
	return &SignUpLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SignUpLogic) SignUp(in *core.SignUpReq) (*core.SignUpResp, error) {
	// todo: add your logic here and delete this line
	found, err := l.svcCtx.DAO.FindOneUserByEmail(l.ctx, in.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	if found != nil {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.EMAIL_HAS_BEEN_REGISTERED), "User existed ,email: %s ,error : %+v", in.Email, err)
	}

	encryptedPW := cryptox.PasswordEncrypt(in.Password, l.svcCtx.Config.Salt)
	u, err := l.svcCtx.DAO.InsertOneUser(l.ctx, in.Name, in.Email, encryptedPW)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	now := time.Now().Unix()
	exp := now + l.svcCtx.Config.AuthConf.AccessExpire
	payLoad := map[string]interface{}{
		ctxtool.CTXJWTUserID: u.Id,
	}

	token, err := jwtx.GetToken(now, exp, l.svcCtx.Config.AuthConf.AccessSecret, payLoad)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.TOKEN_GENERATE_ERROR), "token generated error : %+v", err)
	}

	return &core.SignUpResp{
		Code:        int32(errx.SUCCESS),
		Token:       token,
		ExpiredTime: int32(exp),
	}, nil
}

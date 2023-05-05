package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/cryptox"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/jwtx"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignUpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSignUpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignUpLogic {
	return &UserSignUpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignUpLogic) UserSignUp(req *types.SignUpReq) (resp *types.SignUpResp, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("Call User SignUp API with email : %v, name : %v", req.Email, req.Name)
	found, err := l.svcCtx.DAO.FindOneUserByEmail(l.ctx, req.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if found != nil {
		return nil, errx.NewCustomErrCode(errx.EMAIL_HAS_BEEN_REGISTERED)
	}

	encryptedPW := cryptox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt)
	u, err := l.svcCtx.DAO.InsertOneUser(l.ctx, req.Name, req.Email, encryptedPW)
	if err != nil {
		return nil, errx.NewCustomError(errx.USER_SIGN_UP_FAILED, err.Error())
	}

	now := time.Now().Unix()
	exp := now + l.svcCtx.Config.Auth.AccessExpire
	payLoad := map[string]interface{}{
		ctxtool.CTXJWTUserID: u.ID,
	}

	token, err := jwtx.GetToken(now, exp, l.svcCtx.Config.Auth.AccessSecret, payLoad)
	if err != nil {
		return nil, errx.NewCustomErrCode(errx.TOKEN_GENERATE_ERROR)
	}
	return &types.SignUpResp{
		Code:        http.StatusOK,
		Token:       token,
		ExpiredTime: uint(exp),
	}, nil
}

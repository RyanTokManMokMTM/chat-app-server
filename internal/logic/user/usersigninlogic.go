package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/cryptox"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/jwtx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserSignInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserSignInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserSignInLogic {
	return &UserSignInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserSignInLogic) UserSignIn(req *types.SignInReq) (resp *types.SignInResp, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("Call User Sign In with email: %v ", req.Email)
	u, err := l.svcCtx.DAO.FindOneUserByEmail(l.ctx, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
	}

	encryptedPW := cryptox.PasswordEncrypt(req.Password, l.svcCtx.Config.Salt)
	if strings.Compare(encryptedPW, u.Password) != 0 {
		return nil, errx.NewCustomErrCode(errx.USER_SIGN_IN_FAILED)
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

	return &types.SignInResp{
		Code:        uint(http.StatusOK),
		Token:       token,
		ExpiredTime: uint(exp),
		UserInfo: types.CommonUserInfo{
			ID:            u.ID,
			Uuid:          u.Uuid,
			NickName:      u.NickName,
			Avatar:        u.Avatar,
			Email:         u.Email,
			Cover:         u.Cover,
			StatusMessage: u.StatusMessage,
		},
	}, nil
}

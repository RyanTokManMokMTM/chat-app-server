package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmok/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmok/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	"github.com/ryantokmanmok/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserLogic) SearchUser(req *types.SearchUserReq) (resp *types.SearchUserResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	if len(req.Qurey) == 0 {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, "Missing Search Keyword.")
	}
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	results, err := l.svcCtx.DAO.FindUsers(l.ctx, req.Qurey)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &types.SearchUserResp{ //User Not Exist
				Code: uint(http.StatusNotFound),
			}, nil
		}

		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var users []types.CommonUserInfo
	for _, info := range results {
		if info.ID == userID {
			continue
		}
		users = append(users, types.CommonUserInfo{
			ID:       info.ID,
			Uuid:     info.Uuid,
			NickName: info.NickName,
			Email:    info.Email,
			Avatar:   info.Avatar,
		})
	}

	return &types.SearchUserResp{
		Code:    uint(http.StatusOK),
		Results: users,
	}, nil
}

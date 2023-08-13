package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

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

	var users = make([]types.SearchUserResult, 0)
	for _, info := range results {
		if info.Id == userID {
			continue
		}

		var isFriend = true
		err := l.svcCtx.DAO.FindOneFriend(l.ctx, userID, info.Id)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			isFriend = false
		}

		users = append(users, types.SearchUserResult{
			UserInfo: types.CommonUserInfo{
				ID:            info.Id,
				Uuid:          info.Uuid,
				NickName:      info.NickName,
				Email:         info.Email,
				Avatar:        info.Avatar,
				Cover:         info.Cover,
				StatusMessage: info.StatusMessage,
			}, IsFriend: isFriend,
		})
	}
	if len(users) == 0 {
		return &types.SearchUserResp{ //User Not Exist
			Code: uint(http.StatusNotFound),
		}, nil
	}
	return &types.SearchUserResp{
		Code:    uint(http.StatusOK),
		Results: users,
	}, nil
}

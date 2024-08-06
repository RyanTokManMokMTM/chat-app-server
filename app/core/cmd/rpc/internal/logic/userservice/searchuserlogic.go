package userservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchUserLogic) SearchUser(in *core.SearchUserReq) (*core.SearchUserResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)

	if len(in.Query) == 0 {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "request parameter error : missing searching keyword")
	}
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error %+v", err)
		}
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	results, err := l.svcCtx.DAO.FindUsers(l.ctx, in.Query)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error %+v", err)
		}
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	var users = make([]*core.SearchUserRespResult, 0)
	for _, info := range results {
		if info.Id == userID {
			continue
		}

		var isFriend = true
		_, err := l.svcCtx.DAO.FindOneFriend(l.ctx, userID, info.Id)
		if err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				continue
			}
			isFriend = false
		}

		users = append(users, &core.SearchUserRespResult{
			UserInfo: &core.UserInfo{
				Id:            uint32(info.Id),
				Uuid:          info.Uuid,
				Name:          info.NickName,
				Email:         info.Email,
				Avatar:        info.Avatar,
				Cover:         info.Cover,
				StatusMessage: info.StatusMessage,
			}, IsFriend: isFriend,
		})
	}
	//if len(users) == 0 {
	//	return nil, errx.NewCustomError(errx.NOT_FOUND, err.Error())
	//}

	return &core.SearchUserResp{
		Code:          int32(errx.SUCCESS),
		SearchResults: users,
	}, nil
}

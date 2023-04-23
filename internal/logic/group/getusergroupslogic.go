package group

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

type GetUserGroupsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupsLogic {
	return &GetUserGroupsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserGroupsLogic) GetUserGroups(req *types.GetUserGroupReq) (resp *types.GetUserGroupResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	groups, err := l.svcCtx.DAO.GetUserGroups(l.ctx, userID)
	if err != nil {
		logx.Infof(err.Error())
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var userGroups = make([]types.GroupInfo, 0)
	for _, g := range groups {
		userGroups = append(userGroups, types.GroupInfo{
			ID:        g.GroupInfo.ID,
			Uuid:      g.GroupInfo.Uuid,
			Name:      g.GroupInfo.GroupName,
			Avatar:    g.GroupInfo.GroupAvatar,
			CreatedAt: uint(g.GroupInfo.CreatedAt.Unix()),
		})
	}
	return &types.GetUserGroupResp{
		Code:   uint(http.StatusOK),
		Groups: userGroups,
	}, nil
}

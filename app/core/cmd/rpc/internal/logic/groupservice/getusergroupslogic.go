package groupservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/pagerx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserGroupsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserGroupsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserGroupsLogic {
	return &GetUserGroupsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserGroupsLogic) GetUserGroups(in *core.GetUserGroupReq) (*core.GetUserGroupResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}

	total := l.svcCtx.DAO.CountUserGroups(l.ctx, userID)
	//
	pageLimit := pagerx.GetLimit(uint(in.Limit))
	pageSize := pagerx.GetTotalPageByPageSize(uint(total), pageLimit)
	pageOffset := pagerx.PageOffset(pageSize, uint(in.Page))

	groups, err := l.svcCtx.DAO.GetUserGroups(l.ctx, userID, int(pageOffset), int(pageLimit))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	userGroups := make([]*core.GroupInfo, 0)
	for _, g := range groups {
		userGroups = append(userGroups, &core.GroupInfo{
			Id:        uint32(g.GroupInfo.Id),
			Uuid:      g.GroupInfo.Uuid,
			Name:      g.GroupInfo.GroupName,
			Avatar:    g.GroupInfo.GroupAvatar,
			CreatedAt: uint32(g.GroupInfo.CreatedAt.Unix()),
		})
	}

	return &core.GetUserGroupResp{
		Code:      uint32(errx.SUCCESS),
		GroupInfo: userGroups,
	}, nil
}

package friendservicelogic

import (
	"api/app/common/errx"
	"api/app/common/pagerx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendListLogic) GetFriendList(in *core.GetFriendListReq) (*core.GetFriendListResp, error) {
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

	total, err := l.svcCtx.DAO.CountUserFriend(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	pageLimit := pagerx.GetLimit(uint(in.Limit))
	pageSize := pagerx.GetTotalPageByPageSize(uint(total), pageLimit)
	pageOffset := pagerx.PageOffset(pageLimit, uint(in.Page))

	list, err := l.svcCtx.DAO.GetUserFriendListByPageSize(l.ctx, userID, int(pageOffset), int(pageLimit))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	var respList = make([]*core.UserInfo, 0)
	for _, info := range list {
		respList = append(respList, &core.UserInfo{
			Id:     uint32(info.FriendInfo.Id),
			Uuid:   info.FriendInfo.Uuid,
			Name:   info.FriendInfo.NickName,
			Avatar: info.FriendInfo.Avatar,
			Email:  info.FriendInfo.Email,
			Cover:  info.FriendInfo.Cover,
		})
	}

	return &core.GetFriendListResp{
		Code: uint32(errx.SUCCESS),
		PageInfo: &core.PageableInfo{
			TotalPage: uint32(pageSize),
			Page:      in.Page,
		},
		FriendList: respList,
	}, nil
}

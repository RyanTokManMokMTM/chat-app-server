package friend

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/pagerx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendListLogic {
	return &GetFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendListLogic) GetFriendList(req *types.GetFriendListReq) (resp *types.GetFriendListResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	total, err := l.svcCtx.DAO.CountUserFriend(l.ctx, userID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	pageLimit := pagerx.GetLimit(req.Limit)
	pageSize := pagerx.GetTotalPageByPageSize(uint(total), pageLimit)
	pageOffset := pagerx.PageOffset(pageSize, req.Page)
	list, err := l.svcCtx.DAO.GetUserFriendListByPageSize(l.ctx, userID, int(pageOffset), int(pageLimit))
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	var respList = make([]types.CommonUserInfo, 0)
	for _, info := range list {
		respList = append(respList, types.CommonUserInfo{
			ID:       info.FriendInfo.Id,
			Uuid:     info.FriendInfo.Uuid,
			NickName: info.FriendInfo.NickName,
			Avatar:   info.FriendInfo.Avatar,
			Email:    info.FriendInfo.Email,
			Cover:    info.FriendInfo.Cover,
		})
	}

	//TODO : response to user type
	return &types.GetFriendListResp{
		FriendList: respList,
	}, nil
}

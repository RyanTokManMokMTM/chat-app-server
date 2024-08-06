package friend

import (
	"api/app/common/ctxtool"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get all user friends
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
	rpcResp, rpcErr := l.svcCtx.FriendService.GetFriendList(l.ctx, &core.GetFriendListReq{
		UserId: uint32(userID),
		Page:   uint32(req.Page),
		Limit:  uint32(req.Limit),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	friendList := make([]types.CommonUserInfo, 0)
	for _, fd := range rpcResp.FriendList {
		friendList = append(friendList, types.CommonUserInfo{
			ID:            uint(fd.Id),
			Uuid:          fd.Uuid,
			NickName:      fd.Name,
			Avatar:        fd.Avatar,
			Email:         fd.Email,
			Cover:         fd.Cover,
			StatusMessage: fd.StatusMessage,
		})
	}

	return &types.GetFriendListResp{
		Code: uint(rpcResp.Code),
		PageableInfo: types.PageableInfo{
			TotalPage: uint(rpcResp.PageInfo.TotalPage),
			Page:      uint(rpcResp.PageInfo.Page),
		},
		FriendList: friendList,
	}, nil
}

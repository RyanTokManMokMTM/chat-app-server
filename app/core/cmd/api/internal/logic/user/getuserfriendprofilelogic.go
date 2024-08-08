package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFriendProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get User Friend Profile - with `isFriend` data
func NewGetUserFriendProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFriendProfileLogic {
	return &GetUserFriendProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFriendProfileLogic) GetUserFriendProfile(req *types.GetUserFriendProfileReq) (resp *types.GetUserFriendProfileResp, err error) {
	// todo: add your logic here and delete this line

	if req.UUID == "" && req.UserID == 0 {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}

	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	fdUserID := uint32(req.UserID)
	rpcResp, rpcErr := l.svcCtx.UserService.GetUserFriendProfile(l.ctx, &core.GetUserFriendProfileReq{
		UserId:       uint32(userID),
		FriendUserId: &(fdUserID),
		FriendUuid:   &(req.UUID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.GetUserFriendProfileResp{
		Code: uint(rpcResp.Code),
		UserInfo: types.CommonUserInfo{
			ID:            uint(rpcResp.UserInfo.Id),
			Uuid:          rpcResp.UserInfo.Uuid,
			NickName:      rpcResp.UserInfo.Name,
			Email:         rpcResp.UserInfo.Email,
			Avatar:        rpcResp.UserInfo.Avatar,
			Cover:         rpcResp.UserInfo.Cover,
			StatusMessage: rpcResp.UserInfo.StatusMessage,
		}, IsFriend: rpcResp.IsFriend,
	}, nil
}

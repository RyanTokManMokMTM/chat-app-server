package friend

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendInformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get one friend information
func NewGetFriendInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInformationLogic {
	return &GetFriendInformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendInformationLogic) GetFriendInformation(req *types.GetFriendInfoReq) (resp *types.GetFriendInfoResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.FriendService.GetFriendInformation(l.ctx, &core.GetFriendInfoReq{
		UserId: uint32(userID),
		Uuid:   req.Uuid,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.GetFriendInfoResp{
		Code: uint(rpcResp.Code),
		FriendInfo: types.FriendInfo{
			ID:       uint(rpcResp.Info.Id),
			Uuid:     rpcResp.Info.Uuid,
			NickName: rpcResp.Info.NickName,
			Avatar:   rpcResp.Info.Avatar,
		},
	}, nil
}

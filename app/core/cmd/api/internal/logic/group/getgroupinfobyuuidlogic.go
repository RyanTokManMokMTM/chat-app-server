package group

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoByUUIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get group info by UUID
func NewGetGroupInfoByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoByUUIDLogic {
	return &GetGroupInfoByUUIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupInfoByUUIDLogic) GetGroupInfoByUUID(req *types.GetGroupInfoByUUIDReq) (resp *types.GetGroupInfoByUUIDResp, err error) {
	// todo: add your logic here and delete this line

	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.GroupService.GetGroupInfoByUUID(l.ctx, &core.GetGroupInfoByUUIDReq{
		UserId: uint32(userID),
		Uuid:   req.UUID,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &types.GetGroupInfoByUUIDResp{
		Code: uint(http.StatusOK),
		Result: types.FullGroupInfo{
			GroupInfo: types.GroupInfo{
				ID:        uint(rpcResp.Result.Info.Id),
				Uuid:      rpcResp.Result.Info.Uuid,
				Name:      rpcResp.Result.Info.Name,
				Avatar:    rpcResp.Result.Info.Avatar,
				Desc:      rpcResp.Result.Info.Desc,
				CreatedAt: uint(rpcResp.Result.Info.CreatedAt),
			},
			Members:   uint(rpcResp.Result.Members),
			IsJoined:  rpcResp.Result.IsJoined,
			IsOwner:   rpcResp.Result.IsOwner,
			CreatedBy: rpcResp.Result.CreatedBy,
		},
	}, nil
}

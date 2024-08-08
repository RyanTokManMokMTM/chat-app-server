package sticker

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStickerGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStickerGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupInfoLogic {
	return &GetStickerGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStickerGroupInfoLogic) GetStickerGroupInfo(req *types.GetStickerInfoReq) (resp *types.GetStickerInfoResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, rpcErr := l.svcCtx.StickerService.GetStickerGroupInfo(l.ctx, &core.GetStickerInfoReq{
		StickerGroupUUID: req.StickerUUID,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &types.GetStickerInfoResp{
		Code: uint(rpcResp.Code),
		StickerInfo: types.StickerInfo{
			StickerID:   rpcResp.Info.StickerId,
			StickerName: rpcResp.Info.StickerName,
			StickerThum: rpcResp.Info.StickerThum,
		},
	}, nil
}

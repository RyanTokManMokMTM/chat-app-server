package sticker

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStickerGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStickerGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupListLogic {
	return &GetStickerGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStickerGroupListLogic) GetStickerGroupList(req *types.GetStickerListReq) (resp *types.GetStickerListResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, rpcErr := l.svcCtx.StickerService.GetStickerGroupList(l.ctx, &core.GetStickerListReq{})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	stickers := make([]types.StickerInfo, 0)
	for _, sticker := range rpcResp.Sticker {
		stickers = append(stickers, types.StickerInfo{
			StickerID:   sticker.StickerId,
			StickerName: sticker.StickerName,
			StickerThum: sticker.StickerThum,
		})
	}

	return &types.GetStickerListResp{
		Code:     uint(rpcResp.Code),
		Stickers: stickers,
	}, nil
}

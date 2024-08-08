package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStickersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get user sticker group
func NewGetUserStickersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStickersLogic {
	return &GetUserStickersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserStickersLogic) GetUserStickers(req *types.GetUserStickerReq) (resp *types.GetUserStickerResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.UserService.GetUserStickers(l.ctx, &core.GetUserStickerReq{
		UserId: uint32(userID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	stickers := make([]types.StickerInfo, 0)
	for _, sticker := range rpcResp.StickerInfo {
		stickers = append(stickers, types.StickerInfo{
			StickerID:   sticker.StickerId,
			StickerName: sticker.StickerName,
			StickerThum: sticker.StickerThum,
		})
	}

	return &types.GetUserStickerResp{
		Code:     uint(rpcResp.Code),
		Stickers: stickers,
	}, nil
}

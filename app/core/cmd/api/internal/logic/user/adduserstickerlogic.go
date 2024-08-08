package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserStickerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Add an existing sticker to user account
func NewAddUserStickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserStickerLogic {
	return &AddUserStickerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserStickerLogic) AddUserSticker(req *types.AddStickerReq) (resp *types.AddStickerResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	rpcResp, rpcErr := l.svcCtx.UserService.AddUserSticker(l.ctx, &core.AddStickerReq{
		UserId:      uint32(userID),
		StickerUUID: req.StickerUUID,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.AddStickerResp{
		Code: uint(rpcResp.Code),
	}, nil
}

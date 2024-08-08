package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserStickerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Remove the sticker is added
func NewDeleteUserStickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserStickerLogic {
	return &DeleteUserStickerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserStickerLogic) DeleteUserSticker(req *types.DeleteStickerReq) (resp *types.DeleteStickerResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.UserService.DeleteUserSticker(l.ctx, &core.DeleteStickerReq{
		UserId:      uint32(userID),
		StickerUUID: req.StickerUUID,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.DeleteStickerResp{
		Code: uint(rpcResp.Code),
	}, nil
}

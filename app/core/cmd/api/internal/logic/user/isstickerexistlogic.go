package user

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsStickerExistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Check an existing sticker has been added to user
func NewIsStickerExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsStickerExistLogic {
	return &IsStickerExistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IsStickerExistLogic) IsStickerExist(req *types.IsStickerExistReq) (resp *types.IsStickerExistResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.UserService.IsStickerExist(l.ctx, &core.IsStickerExistReq{
		UserId:      uint32(userID),
		StickerUUID: req.StickerUUID,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.IsStickerExistResp{
		Code:    uint(rpcResp.Code),
		IsExist: rpcResp.IsExist,
	}, nil
}

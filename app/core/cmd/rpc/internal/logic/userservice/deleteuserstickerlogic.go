package userservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserStickerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserStickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserStickerLogic {
	return &DeleteUserStickerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserStickerLogic) DeleteUserSticker(in *core.DeleteStickerReq) (*core.DeleteStickerResp, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteStickerResp{}, nil
}

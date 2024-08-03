package userservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserStickerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserStickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserStickerLogic {
	return &AddUserStickerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserStickerLogic) AddUserSticker(in *core.AddStickerReq) (*core.AddStickerResp, error) {
	// todo: add your logic here and delete this line

	return &core.AddStickerResp{}, nil
}

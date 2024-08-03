package userservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStickersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserStickersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStickersLogic {
	return &GetUserStickersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserStickersLogic) GetUserStickers(in *core.GetUserStickerReq) (*core.GetUserStickerResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserStickerResp{}, nil
}

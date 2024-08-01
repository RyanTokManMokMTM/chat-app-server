package sticker

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

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

	return
}

package sticker

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

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

	return
}

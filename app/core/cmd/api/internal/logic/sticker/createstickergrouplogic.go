package sticker

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStickerGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateStickerGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStickerGroupLogic {
	return &CreateStickerGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateStickerGroupLogic) CreateStickerGroup(req *types.CreateStickerGroupReq) (resp *types.CreateStickerGroupResp, err error) {
	// todo: add your logic here and delete this line

	return
}

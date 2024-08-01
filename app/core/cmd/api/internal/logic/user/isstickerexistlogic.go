package user

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

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

	return
}

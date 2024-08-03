package userservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsStickerExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsStickerExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsStickerExistLogic {
	return &IsStickerExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsStickerExistLogic) IsStickerExist(in *core.IsStickerExistReq) (*core.IsStickerExistResp, error) {
	// todo: add your logic here and delete this line

	return &core.IsStickerExistResp{}, nil
}

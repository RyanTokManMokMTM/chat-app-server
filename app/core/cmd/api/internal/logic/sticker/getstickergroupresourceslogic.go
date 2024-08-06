package sticker

import (
	"api/app/core/cmd/rpc/types/core"
	"context"
	"net/http"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStickerGroupResourcesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStickerGroupResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupResourcesLogic {
	return &GetStickerGroupResourcesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStickerGroupResourcesLogic) GetStickerGroupResources(req *types.GetStickerResourcesReq) (resp *types.GetStickerResourcesResp, err error) {
	// todo: add your logic here and delete this line
	rpcResp, rpcErr := l.svcCtx.StickerService.GetStickerGroupResources(l.ctx, &core.GetStickerResourcesReq{
		StickerGroupUUID: req.StickerGroupUUID,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &types.GetStickerResourcesResp{
		Code:          http.StatusOK,
		StickerId:     rpcResp.StickerId,
		ResourcesPath: rpcResp.Paths,
	}, nil
}

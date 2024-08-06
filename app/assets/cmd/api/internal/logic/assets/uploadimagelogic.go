package assets

import (
	"api/app/assets/cmd/api/internal/svc"
	"api/app/assets/cmd/api/internal/types"
	"api/app/assets/cmd/rpc/types/assets_api"
	"api/app/common/errx"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Upload only image
func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(req *types.UploadImageReq) (resp *types.UploadImageResp, err error) {
	// todo: add your logic here and delete this line
	if len(req.Data) == 0 {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, "data is empty")
	}

	rpcResp, rpcErr := l.svcCtx.AssetRPC.UploadImage(l.ctx, &assets_api.UploadImageReq{
		Format:    req.ImageType,
		Base64Str: req.Data,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.UploadImageResp{
		Code: uint(rpcResp.Code),
		Path: rpcResp.Path,
	}, nil

}

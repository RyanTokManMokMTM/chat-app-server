package assets

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageByBase64Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Upload image by base64
func NewUploadImageByBase64Logic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageByBase64Logic {
	return &UploadImageByBase64Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageByBase64Logic) UploadImageByBase64(req *types.UploadImageByBase64Req) (resp *types.UploadImageByBase64Resp, err error) {
	// todo: add your logic here and delete this line
	if len(req.Data) == 0 {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, "data is empty")
	}

	rpcResp, rpcErr := l.svcCtx.AssetRPC.UploadImageByBase64(l.ctx, &assets_api.UploadImageReq{
		Format: req.ImageType,
		Data:   []byte(req.Data),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.UploadImageByBase64Resp{
		Code: uint(rpcResp.Code),
		Path: rpcResp.Path,
	}, nil
}

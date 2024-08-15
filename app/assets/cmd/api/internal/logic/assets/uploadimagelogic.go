package assets

import (
	"bytes"
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"io"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// Upload only image
func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadImageLogic) UploadImage(req *types.UploadImageReq) (resp *types.UploadImageResp, err error) {
	// todo: add your logic here and delete this line
	file, header, err := l.r.FormFile("image")
	if err != nil {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, err.Error())
	}
	defer file.Close()
	fileName := header.Filename
	fileFormat := strings.Split(fileName, ".")
	if len(fileFormat) < 2 {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, "image type incorrect")
	}

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errx.NewCustomErrCode(errx.SERVER_COMMON_ERROR)
	}

	rpcResp, rpcErr := l.svcCtx.AssetRPC.UploadImageByByte(l.ctx, &assets_api.UploadImageReq{
		Format: fileFormat[1],
		Data:   buffer.Bytes(),
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

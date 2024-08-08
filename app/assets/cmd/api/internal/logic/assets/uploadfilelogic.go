package assets

import (
	"bytes"
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"io"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadFileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// Upload any file
func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadFileLogic {
	return &UploadFileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadFileLogic) UploadFile(req *types.UploadFileReq) (resp *types.UploadFileResp, err error) {
	// todo: add your logic here and delete this line
	file, header, err := l.r.FormFile("file")
	if err != nil {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, err.Error())
	}
	defer file.Close()

	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errx.NewCustomErrCode(errx.SERVER_COMMON_ERROR)
	}

	rpcResp, rpcErr := l.svcCtx.AssetRPC.UploadFile(l.ctx, &assets_api.UploadFileReq{
		FileName: header.Filename,
		Data:     buffer.Bytes(),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.UploadFileResp{
		Code: uint(rpcResp.Code),
		Path: rpcResp.Path,
	}, nil
}

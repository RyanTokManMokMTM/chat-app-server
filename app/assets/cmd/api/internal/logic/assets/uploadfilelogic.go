package assets

import (
	"api/app/assets/cmd/api/internal/svc"
	"api/app/assets/cmd/api/internal/types"
	"api/app/assets/cmd/rpc/types/assets_api"
	"api/app/common/errx"
	"context"
	"net/http"

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
	//return
	file, header, err := l.r.FormFile("file")
	if err != nil {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, err.Error())
	}
	defer file.Close()

	fileBytes := make([]byte, 0)
	_, err = file.Read(fileBytes)
	if err != nil {
		return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
	}

	stream, err := l.svcCtx.AssetRPC.UploadFile(l.ctx)
	err = stream.Send(&assets_api.UploadFileReq{
		FileName: header.Filename,
		Type:     "file",
		Data:     fileBytes,
	})
	if err != nil {
		return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
	}
	//
	streamResp, err := stream.CloseAndRecv()
	return &types.UploadFileResp{
		Code: uint(http.StatusOK),
		Path: streamResp.Path,
	}, nil

}

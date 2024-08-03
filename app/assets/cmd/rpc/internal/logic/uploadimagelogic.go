package logic

import (
	"api/app/common/errx"
	"api/app/common/uploadx"
	"context"
	"io"

	"api/app/assets/cmd/rpc/internal/svc"
	"api/app/assets/cmd/rpc/types/assets_api"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadImageLogic) UploadImage(stream assets_api.AssetRPC_UploadImageServer) error {
	// todo: add your logic here and delete this line
	imageData := make([]byte, 0)
	imageName := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if imageName == "" {
			imageName = req.GetFileName()
		}

		data := req.GetData()
		imageData = append(imageData, data...)
	}

	path, err := uploadx.SaveBytesIntoFile(imageName, imageData, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return stream.SendAndClose(&assets_api.UploadImageResp{
			Code: int32(errx.FILE_UPLOAD_FAILED),
			Path: "",
		})
	}
	return stream.SendAndClose(&assets_api.UploadImageResp{
		Code: int32(errx.SUCCESS),
		Path: path,
	})
}

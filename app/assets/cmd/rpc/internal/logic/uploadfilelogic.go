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

type UploadFileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadFileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadFileLogic {
	return &UploadFileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadFileLogic) UploadFile(stream assets_api.AssetRPC_UploadFileServer) error {
	// todo: add your logic here and delete this line
	//TODO: Receiving the file chunks
	fileData := make([]byte, 0)
	fileName := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		chunk := req.GetData()
		if fileName == "" {
			fileName = req.GetFileName()
		}

		fileData = append(fileData, chunk...) //added data to fileData
	}

	path, err := uploadx.SaveBytesIntoFile(fileName, fileData, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return stream.SendAndClose(&assets_api.UploadFileResp{
			Code: int32(errx.SERVER_COMMON_ERROR),
			Path: "",
		})
	}
	//save and
	return stream.SendAndClose(&assets_api.UploadFileResp{
		Code: int32(errx.SUCCESS),
		Path: path,
	})
}

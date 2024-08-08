package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"

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

func (l *UploadFileLogic) UploadFile(in *assets_api.UploadFileReq) (*assets_api.UploadFileResp, error) {
	// todo: add your logic here and delete this line
	if len(in.Data) == 0 {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "data is nill")
	}
	path, err := uploadx.SaveBytesIntoFile(in.FileName, in.Data, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return nil, err
	}
	return &assets_api.UploadFileResp{
		Code: int32(errx.SUCCESS),
		Path: path,
	}, nil
}

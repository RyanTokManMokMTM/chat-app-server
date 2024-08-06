package logic

import (
	"api/app/assets/cmd/rpc/internal/svc"
	"api/app/assets/cmd/rpc/types/assets_api"
	"api/app/common/errx"
	"api/app/common/uploadx"
	"context"
	"github.com/pkg/errors"

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

func (l *UploadImageLogic) UploadImage(in *assets_api.UploadImageReq) (*assets_api.UploadImageResp, error) {
	// todo: add your logic here and delete this line
	if len(in.Base64Str) == 0 {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "base64 is nil")
	}

	path, err := uploadx.SaveImageByBase64(in.Base64Str, in.Format, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
	}

	return &assets_api.UploadImageResp{
		Code: int32(errx.SUCCESS),
		Path: path,
	}, nil
}

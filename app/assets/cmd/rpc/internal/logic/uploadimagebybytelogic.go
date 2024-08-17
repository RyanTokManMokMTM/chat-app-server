package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageByByteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadImageByByteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageByByteLogic {
	return &UploadImageByByteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadImageByByteLogic) UploadImageByByte(in *assets_api.UploadImageReq) (*assets_api.UploadImageResp, error) {
	// todo: add your logic here and delete this line
	logx.WithContext(l.ctx).Info("Testing................")
	if len(in.Data) == 0 {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "image bytes data is nil")
	}

	imageName := fmt.Sprintf("%s.%s", uuid.NewString(), in.Format)
	logx.Info(imageName)
	path, err := uploadx.SaveBytesIntoFile(imageName, in.Data, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
	}

	return &assets_api.UploadImageResp{
		Code: int32(errx.SUCCESS),
		Path: path,
	}, nil
}

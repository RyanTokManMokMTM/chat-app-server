package user

import (
	"api/app/common/ctxtool"
	"api/app/common/uploadx"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"bytes"
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// Upload and update user avatar
func NewUploadUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadUserAvatarLogic) UploadUserAvatar(req *types.UploadUserAvatarReq) (resp *types.UploadUserAvatarResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	file, header, err := l.r.FormFile(uploadx.AvatarFileField)
	if err != nil {
		return nil, errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED)
	}

	buffer := bytes.NewBuffer(nil)
	fileName := header.Filename

	_, err = io.Copy(buffer, file)
	if err != nil {
		return nil, errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED)
	}

	rpcResp, rpcErr := l.svcCtx.UserService.UploadUserAvatar(l.ctx, &core.UploadUserAvatarReq{
		UserId:   uint32(userID),
		FileName: fileName,
		Data:     buffer.Bytes(),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.UploadUserAvatarResp{
		Code: uint(rpcResp.Code),
		Path: rpcResp.Path,
	}, nil
}

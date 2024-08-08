package user

import (
	"bytes"
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"io"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// Upload and update user cover
func NewUploadUserCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadUserCoverLogic {
	return &UploadUserCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadUserCoverLogic) UploadUserCover(req *types.UploadUserAvatarReq) (resp *types.UploadUserAvatarResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	file, header, err := l.r.FormFile(uploadx.CoverFileField)
	if err != nil {
		return nil, errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED)
	}

	buffer := bytes.NewBuffer(nil)
	fileName := header.Filename

	_, err = io.Copy(buffer, file)
	if err != nil {
		return nil, errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED)
	}

	rpcResp, rpcErr := l.svcCtx.UserService.UploadUserCover(l.ctx, &core.UploadUserCoverReq{
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

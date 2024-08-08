package group

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
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

type UploadGroupAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// Update and update group avatar
func NewUploadGroupAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadGroupAvatarLogic {
	return &UploadGroupAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadGroupAvatarLogic) UploadGroupAvatar(req *types.UploadGroupAvatarReq) (resp *types.UploadGroupAvatarResp, err error) {
	// todo: add your logic here and delete this line
	file, header, err := l.r.FormFile(uploadx.AvatarFileField)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "avatar file not exist")
	}

	fileName := header.Filename
	buffer := bytes.NewBuffer(nil)
	_, err = io.Copy(buffer, file)
	if err != nil {
		return nil, err
	}

	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.GroupService.UploadGroupAvatar(l.ctx, &core.UploadGroupAvatarReq{
		UserId:         uint32(userID),
		GroupId:        uint32(req.GroupID),
		AvatarFileName: fileName,
		Data:           buffer.Bytes(),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &types.UploadGroupAvatarResp{
		Code: uint(rpcResp.Code),
		Path: rpcResp.Path,
	}, nil
}

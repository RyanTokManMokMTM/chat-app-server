package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/uploadx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

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
	userId := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	name, err := uploadx.UploadFileFromRequest(l.r, l.svcCtx.Config.MaxFileSize, uploadx.AvatarFileField, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return nil, errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED)
	}

	path := fmt.Sprintf("%v", name)
	l.svcCtx.Uow.UserRepo().UpdateOneUserAvatar(l.ctx, userId, path)
	return &types.UploadUserAvatarResp{
		Code: uint(http.StatusOK),
		Path: path,
	}, nil
}

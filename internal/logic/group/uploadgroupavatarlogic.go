package group

import (
	"context"
	"errors"
	"fmt"
	"github.com/ryantokmanmok/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmok/chat-app-server/common/errx"
	"github.com/ryantokmanmok/chat-app-server/common/uploadx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	"github.com/ryantokmanmok/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadGroupAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

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
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	group, err := l.svcCtx.DAO.FindOneGroup(l.ctx, req.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if group.GroupLead != userID {
		return nil, errx.NewCustomErrCode(errx.NO_GROUP_AUTHORITY)
	}

	//upload
	name, err := uploadx.UploadFileFromRequest(l.r, l.svcCtx.Config.MaxFileSize, uploadx.AvatarFileField, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
	}

	path := fmt.Sprintf("/%s", name)
	if err := l.svcCtx.DAO.UpdateOneGroupAvatar(l.ctx, group.ID, path); err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	return &types.UploadGroupAvatarResp{
		Code: uint(http.StatusOK),
		Path: path,
	}, nil
}

package groupservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadGroupAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadGroupAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadGroupAvatarLogic {
	return &UploadGroupAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadGroupAvatarLogic) UploadGroupAvatar(in *core.UploadGroupAvatarReq) (*core.UploadGroupAvatarResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist")
		}
		return nil, err
	}

	group, err := l.svcCtx.DAO.FindOneGroup(l.ctx, uint(in.GroupId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.GROUP_NOT_EXIST), "group not exist")
		}
		return nil, err
	}

	if group.GroupLead != userID {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NO_GROUP_AUTHORITY), "No group authrotiy")
	}

	//upload
	avatarPaths, err := uploadx.SaveBytesIntoFile(in.AvatarFileName, in.Data, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED), "File upload filed, error : %+v", err)
	}

	if err := l.svcCtx.DAO.UpdateOneGroupAvatar(l.ctx, group.Id, avatarPaths); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.UploadGroupAvatarResp{
		Code: uint32(errx.SUCCESS),
		Path: avatarPaths,
	}, nil
}

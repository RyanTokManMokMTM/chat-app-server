package groupservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

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

func (l *UploadGroupAvatarLogic) UploadGroupAvatar(stream core.GroupService_UploadGroupAvatarServer) error {
	// todo: add your logic here and delete this line

	return nil
}

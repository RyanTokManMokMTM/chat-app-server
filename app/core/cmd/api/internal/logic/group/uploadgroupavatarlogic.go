package group

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadGroupAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Update and update group avatar
func NewUploadGroupAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadGroupAvatarLogic {
	return &UploadGroupAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadGroupAvatarLogic) UploadGroupAvatar(req *types.UploadGroupAvatarReq) (resp *types.UploadGroupAvatarResp, err error) {
	// todo: add your logic here and delete this line

	return
}

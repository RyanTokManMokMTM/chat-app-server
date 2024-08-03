package userservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadUserCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserCoverLogic {
	return &UploadUserCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadUserCoverLogic) UploadUserCover(stream core.UserService_UploadUserCoverServer) error {
	// todo: add your logic here and delete this line

	return nil
}

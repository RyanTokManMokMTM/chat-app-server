package health

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HealthCheckLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHealthCheckLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HealthCheckLogic {
	return &HealthCheckLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HealthCheckLogic) HealthCheck() (resp *types.HealthCheckResp, err error) {
	// todo: add your logic here and delete this line

	return
}

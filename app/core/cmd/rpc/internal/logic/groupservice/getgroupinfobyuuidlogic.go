package groupservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoByUUIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupInfoByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoByUUIDLogic {
	return &GetGroupInfoByUUIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupInfoByUUIDLogic) GetGroupInfoByUUID(in *core.GetGroupInfoByUUIDReq) (*core.GetGroupInfoByUUIDResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetGroupInfoByUUIDResp{}, nil
}

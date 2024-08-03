package friendservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendInformationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInformationLogic {
	return &GetFriendInformationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFriendInformationLogic) GetFriendInformation(in *core.GetFriendInfoReq) (*core.GetFriendInfoResp, error) {
	// todo: add your logic here and delete this line

	return &core.GetFriendInfoResp{}, nil
}

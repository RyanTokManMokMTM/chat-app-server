package friend

import (
	"context"

	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendInformationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get one friend information
func NewGetFriendInformationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendInformationLogic {
	return &GetFriendInformationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFriendInformationLogic) GetFriendInformation(req *types.GetFriendInfoReq) (resp *types.GetFriendInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}

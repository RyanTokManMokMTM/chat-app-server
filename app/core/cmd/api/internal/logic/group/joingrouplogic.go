package group

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Join a new group by groupID
func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinGroupLogic) JoinGroup(req *types.JoinGroupReq) (resp *types.JoinGroupResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.GroupService.JoinGroup(l.ctx, &core.JoinGroupReq{
		UserId:  uint32(userID),
		GroupId: uint32(req.GroupID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}
	return &types.JoinGroupResp{
		Code: uint(rpcResp.Code),
	}, nil
}

package group

import (
	"api/app/common/ctxtool"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Create a new group with group members
func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupReq) (resp *types.CreateGroupResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	groupMember := make([]uint32, 0)
	for _, member := range req.GroupMembers {
		groupMember = append(groupMember, uint32(member))
	}
	rpcResp, rpcErr := l.svcCtx.GroupService.CreateGroup(l.ctx, &core.CreateGroupReq{
		UserId:       uint32(userID),
		GroupName:    req.GroupName,
		GroupMembers: groupMember,
		AvatarData:   []byte(req.GroupAvatar),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &types.CreateGroupResp{
		Code:        uint(rpcResp.Code),
		GroupUUID:   rpcResp.GroupUUID,
		GroupAvatar: rpcResp.Avatar,
	}, nil
}

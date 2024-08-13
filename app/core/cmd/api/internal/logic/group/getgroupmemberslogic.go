package group

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Get an group members
func NewGetGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMembersLogic {
	return &GetGroupMembersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupMembersLogic) GetGroupMembers(req *types.GetGroupMembersReq) (resp *types.GetGroupMembersResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.GroupService.GetGroupMembers(l.ctx, &core.GetGroupMembersReq{
		UserId:  uint32(userID),
		GroupID: uint32(req.GroupID),
		Page:    uint32(req.Page),
		Limit:   uint32(req.Limit),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	membersList := make([]types.GroupMemberInfo, 0)
	for _, member := range rpcResp.MemberInfo {
		membersList = append(membersList, types.GroupMemberInfo{
			CommonUserInfo: types.CommonUserInfo{
				ID:            uint(member.UserInfo.Id),
				Uuid:          member.UserInfo.Uuid,
				NickName:      member.UserInfo.Name,
				Avatar:        member.UserInfo.Avatar,
				Email:         member.UserInfo.Email,
				Cover:         member.UserInfo.Cover,
				StatusMessage: member.UserInfo.StatusMessage,
			},
			IsGroupLead: member.IsGroupLead,
		})
	}

	return &types.GetGroupMembersResp{
		Code:       uint(http.StatusOK),
		MemberList: membersList,
	}, nil
}

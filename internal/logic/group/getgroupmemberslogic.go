package group

import (
	"context"
	"errors"
	"github.com/ryantokmanmok/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmok/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmok/chat-app-server/internal/svc"
	"github.com/ryantokmanmok/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMembersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

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
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	_, err = l.svcCtx.DAO.FindOneGroup(l.ctx, req.GroupID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	members, err := l.svcCtx.DAO.GetGroupMembers(l.ctx, req.GroupID)

	var membersList = make([]types.GroupMemberInfo, 0)
	for _, mem := range members {
		membersList = append(membersList, types.GroupMemberInfo{
			CommonUserInfo: types.CommonUserInfo{
				ID:       mem.MemberInfo.ID,
				Uuid:     mem.MemberInfo.Uuid,
				NickName: mem.MemberInfo.NickName,
				Avatar:   mem.MemberInfo.Avatar,
			},
			IsGroupLead: mem.GroupInfo.GroupLead == mem.MemberInfo.ID,
		})
	}

	return &types.GetGroupMembersResp{
		Code:       uint(http.StatusOK),
		MemberList: membersList,
	}, nil
}

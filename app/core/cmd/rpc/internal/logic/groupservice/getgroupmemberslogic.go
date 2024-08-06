package groupservicelogic

import (
	"api/app/common/errx"
	"api/app/common/pagerx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupMembersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGroupMembersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupMembersLogic {
	return &GetGroupMembersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetGroupMembersLogic) GetGroupMembers(in *core.GetGroupMembersReq) (*core.GetGroupMembersResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}

	_, err = l.svcCtx.DAO.FindOneGroup(l.ctx, uint(in.GroupID))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.GROUP_NOT_EXIST), "group not exist, error : %+v", err)
		}
		return nil, err
	}

	total, err := l.svcCtx.DAO.CountGroupMembers(l.ctx, uint(in.GroupID))
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	pageLimit := pagerx.GetLimit(uint(in.Limit))
	pageSize := pagerx.GetTotalPageByPageSize(uint(total), pageLimit)
	pageOffset := pagerx.PageOffset(pageSize, uint(in.Page))

	members, err := l.svcCtx.DAO.GetGroupMembers(l.ctx, uint(in.GroupID), int(pageOffset), int(pageLimit))

	var membersList = make([]*core.GroupMemberInfo, 0)
	for _, mem := range members {
		membersList = append(membersList, &core.GroupMemberInfo{
			UserInfo: &core.UserInfo{
				Id:     uint32(mem.MemberInfo.Id),
				Uuid:   mem.MemberInfo.Uuid,
				Name:   mem.MemberInfo.NickName,
				Avatar: mem.MemberInfo.Avatar,
			},
			IsGroupLead: mem.GroupInfo.GroupLead == mem.MemberInfo.Id,
		})
	}

	return &core.GetGroupMembersResp{
		Code:       uint32(errx.SUCCESS),
		MemberInfo: membersList,
	}, nil
}

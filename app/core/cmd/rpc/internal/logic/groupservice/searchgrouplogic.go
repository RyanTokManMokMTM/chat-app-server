package groupservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGroupLogic {
	return &SearchGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchGroupLogic) SearchGroup(in *core.SearchGroupReq) (*core.SearchGroupResp, error) {
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

	if len(in.Query) == 0 {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "required query keywork")
	}

	groups, err := l.svcCtx.DAO.SearchGroup(l.ctx, in.Query)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	groupInfo := make([]*core.FullGroupInfo, 0)
	for _, group := range groups {

		count, err := l.svcCtx.DAO.CountGroupMembers(l.ctx, group.Id)
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			continue
		}

		u, err := l.svcCtx.DAO.FindOneGroupMember(l.ctx, group.Id, userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(l.ctx).Error(err)
			continue
		}

		isJoined := u != nil

		groupInfo = append(groupInfo, &core.FullGroupInfo{
			Info: &core.GroupInfo{
				Id:        uint32(group.Id),
				Uuid:      group.Uuid,
				Name:      group.GroupName,
				Avatar:    group.GroupAvatar,
				CreatedAt: uint32(group.CreatedAt.Unix()),
			},
			Members:   uint32(count),
			IsJoined:  isJoined,
			CreatedBy: group.LeadInfo.NickName,
			IsOwner:   group.GroupLead == userID,
		})
	}

	return &core.SearchGroupResp{
		Code:    uint32(errx.SUCCESS),
		Results: groupInfo,
	}, nil
}

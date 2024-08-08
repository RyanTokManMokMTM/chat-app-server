package groupservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

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
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}

	group, err := l.svcCtx.DAO.FindOneGroupByUUID(l.ctx, in.Uuid)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.GROUP_NOT_EXIST), "group not exist, error : %+v", err)
		}
		return nil, err
	}

	isJoined := true
	_, err = l.svcCtx.DAO.FindOneGroupMember(l.ctx, group.Id, userID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}
		isJoined = false
	}

	count, err := l.svcCtx.DAO.CountGroupMembers(l.ctx, group.Id)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.GetGroupInfoByUUIDResp{
		Code: uint32(errx.SUCCESS),
		Result: &core.FullGroupInfo{
			Info: &core.GroupInfo{
				Id:        uint32(group.Id),
				Uuid:      group.Uuid,
				Name:      group.GroupName,
				Avatar:    group.GroupAvatar,
				Desc:      group.GroupDesc,
				CreatedAt: uint32(uint(group.CreatedAt.Unix())),
			},
			Members:   uint32(uint(count)),
			IsJoined:  isJoined,
			IsOwner:   group.GroupLead == userID,
			CreatedBy: group.LeadInfo.NickName,
		},
	}, nil
}

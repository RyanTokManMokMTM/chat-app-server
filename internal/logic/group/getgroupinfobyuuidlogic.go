package group

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGroupInfoByUUIDLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGroupInfoByUUIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGroupInfoByUUIDLogic {
	return &GetGroupInfoByUUIDLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGroupInfoByUUIDLogic) GetGroupInfoByUUID(req *types.GetGroupInfoByUUIDReq) (resp *types.GetGroupInfoByUUIDResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	group, err := l.svcCtx.DAO.FindOneGroupByUUID(l.ctx, req.UUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.GROUP_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	isJoined := true
	_, err = l.svcCtx.DAO.FindOneGroupMember(l.ctx, group.Id, userID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
		isJoined = false
	}

	count, err := l.svcCtx.DAO.CountGroupMembers(l.ctx, group.Id)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &types.GetGroupInfoByUUIDResp{
		Code: uint(http.StatusOK),
		Result: types.FullGroupInfo{
			GroupInfo: types.GroupInfo{
				ID:        group.Id,
				Uuid:      group.Uuid,
				Name:      group.GroupName,
				Avatar:    group.GroupAvatar,
				Desc:      group.GroupDesc,
				CreatedAt: uint(group.CreatedAt.Unix()),
			},
			Members:   uint(count),
			IsJoined:  isJoined,
			IsOwner:   group.GroupLead == userID,
			CreatedBy: group.LeadInfo.NickName,
		},
	}, nil
}

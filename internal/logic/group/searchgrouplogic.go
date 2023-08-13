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

type SearchGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchGroupLogic {
	return &SearchGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchGroupLogic) SearchGroup(req *types.SearchGroupReq) (resp *types.SearchGroupResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if len(req.Qurey) == 0 {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}

	groups, err := l.svcCtx.DAO.SearchGroup(l.ctx, req.Qurey)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	groupInfo := make([]types.FullGroupInfo, 0)
	for _, group := range groups {

		count, err := l.svcCtx.DAO.CountGroupMembers(l.ctx, group.Id)
		if err != nil {
			logx.Error(err.Error())
			continue
		}

		u, err := l.svcCtx.DAO.FindOneGroupMember(l.ctx, group.Id, userID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logx.Error(err.Error())
			continue
		}

		isJoined := u != nil

		groupInfo = append(groupInfo, types.FullGroupInfo{
			GroupInfo: types.GroupInfo{
				ID:        group.Id,
				Uuid:      group.Uuid,
				Name:      group.GroupName,
				Avatar:    group.GroupAvatar,
				CreatedAt: uint(group.CreatedAt.Unix()),
			},
			Members:  uint(count),
			IsJoined: isJoined,
			IsOwner:  group.GroupLead == userID,
		})
	}
	return &types.SearchGroupResp{
		Code:    uint(http.StatusOK),
		Results: groupInfo,
	}, nil
}

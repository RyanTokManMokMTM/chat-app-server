package group

import (
	"api/app/common/ctxtool"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Search group by name
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
	rpcResp, rpcErr := l.svcCtx.GroupService.SearchGroup(l.ctx, &core.SearchGroupReq{
		UserId: uint32(userID),
		Query:  req.Qurey,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	groupInfo := make([]types.FullGroupInfo, 0)
	for _, group := range rpcResp.Results {
		groupInfo = append(groupInfo, types.FullGroupInfo{
			GroupInfo: types.GroupInfo{
				ID:        uint(group.Info.Id),
				Uuid:      group.Info.Uuid,
				Name:      group.Info.Uuid,
				Avatar:    group.Info.Avatar,
				Desc:      group.Info.Desc,
				CreatedAt: uint(group.Info.CreatedAt),
			},
			Members:   uint(group.Members),
			IsJoined:  group.IsJoined,
			IsOwner:   group.IsOwner,
			CreatedBy: group.CreatedBy,
		})
	}

	return &types.SearchGroupResp{
		Code:    uint(rpcResp.Code),
		Results: groupInfo,
	}, nil
}

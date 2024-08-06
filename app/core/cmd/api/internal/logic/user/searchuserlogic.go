package user

import (
	"api/app/common/ctxtool"
	"api/app/common/errx"
	"api/app/core/cmd/api/internal/svc"
	"api/app/core/cmd/api/internal/types"
	"api/app/core/cmd/rpc/types/core"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Search user by name
func NewSearchUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchUserLogic {
	return &SearchUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchUserLogic) SearchUser(req *types.SearchUserReq) (resp *types.SearchUserResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	if len(req.Qurey) == 0 {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, "Missing Search Keyword.")
	}

	rpcResp, rpcErr := l.svcCtx.UserService.SearchUser(l.ctx, &core.SearchUserReq{
		UserId: uint32(userID),
		Query:  req.Qurey,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	results := make([]types.SearchUserResult, 0)
	for _, result := range rpcResp.SearchResults {
		results = append(results, types.SearchUserResult{
			UserInfo: types.CommonUserInfo{
				ID:            uint(result.UserInfo.Id),
				Uuid:          result.UserInfo.Uuid,
				NickName:      result.UserInfo.Name,
				Avatar:        result.UserInfo.Avatar,
				Email:         result.UserInfo.Email,
				Cover:         result.UserInfo.Cover,
				StatusMessage: result.UserInfo.StatusMessage,
			},
			IsFriend: result.IsFriend,
		})
	}

	return &types.SearchUserResp{
		Code:    uint(rpcResp.Code),
		Results: results,
	}, nil
}

package friend

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

type DeleteFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLogic {
	return &DeleteFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteFriendLogic) DeleteFriend(req *types.DeleteFriendReq) (resp *types.DeleteFriendResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//TODO: Find Friend
	_, err = l.svcCtx.DAO.FindOneFriend(l.ctx, userID, req.UserID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.NOT_YET_FRIEND)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if err := l.svcCtx.DAO.DeleteOneFriend(l.ctx, userID, req.UserID); err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	return &types.DeleteFriendResp{
		Code: uint(http.StatusOK),
	}, nil
}

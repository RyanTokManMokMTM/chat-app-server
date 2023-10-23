package user

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserStickerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserStickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserStickerLogic {
	return &AddUserStickerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserStickerLogic) AddUserSticker(req *types.AddStickerReq) (resp *types.AddStickerResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return
}

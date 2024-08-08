package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserStickerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserStickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserStickerLogic {
	return &DeleteUserStickerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserStickerLogic) DeleteUserSticker(in *core.DeleteStickerReq) (*core.DeleteStickerResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not found ,error : %+v", err)
		}
		return nil, err
	}

	sticker, err := l.svcCtx.DAO.FindOneStickerFromUser(l.ctx, userID, in.StickerUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STICKER_NOT_EXIST), "sticker not exist,error :%+v", err)
		}
		return nil, err
	}

	if err := l.svcCtx.DAO.DeleteOneStickerFromUser(l.ctx, userID, sticker); err != nil {
		return nil, err
	}

	return &core.DeleteStickerResp{
		Code: int32(errx.SUCCESS),
	}, nil
}

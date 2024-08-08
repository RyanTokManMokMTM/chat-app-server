package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/models"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserStickerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserStickerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserStickerLogic {
	return &AddUserStickerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddUserStickerLogic) AddUserSticker(in *core.AddStickerReq) (*core.AddStickerResp, error) {
	// todo: add your logic here and delete this line
	userId := uint(in.UserId)

	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not found,error : %+v", err)
		}
		return nil, err
	}

	_, err = l.svcCtx.DAO.FindOneStickerGroupByStickerUUID(l.ctx, in.StickerUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STICKER_NOT_EXIST), "Sticker not found,error %+v,", err)
		}
		return nil, err //db error
	}

	_, err = l.svcCtx.DAO.FindOneStickerFromUser(l.ctx, userId, in.StickerUUID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err := l.svcCtx.DAO.InsertOneStickerToUser(l.ctx, userId, &models.Sticker{Uuid: in.StickerUUID}); err != nil {
		return nil, err
	}

	return &core.AddStickerResp{
		Code: int32(errx.SUCCESS), //return system error code
	}, nil
}

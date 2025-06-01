package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/models"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"gorm.io/gorm"

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
	userId := ctxtool.GetUserIDFromCTX(l.ctx)

	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	_, err = l.svcCtx.Uow.StickerResourcesRepo().FindOneByUuid(l.ctx, req.StickerUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.STICKER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	found, err := l.svcCtx.Uow.UserRepo().FindOneSticker(l.ctx, userId, req.StickerUUID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	logx.Info(found)

	if err := l.svcCtx.Uow.UserRepo().InsertOneSticker(l.ctx, userId, models.Sticker{Uuid: req.StickerUUID}); err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &types.AddStickerResp{
		Code: http.StatusOK,
	}, nil
}

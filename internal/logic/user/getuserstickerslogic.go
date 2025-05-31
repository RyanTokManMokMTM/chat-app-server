package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStickersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserStickersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStickersLogic {
	return &GetUserStickersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserStickersLogic) GetUserStickers(req *types.GetUserStickerReq) (resp *types.GetUserStickerResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)

	_, err = l.svcCtx.Uow.UserRepo().FindOneUserByID(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	stickerList := make([]types.StickerInfo, 0)
	stickers, err := l.svcCtx.Uow.UserRepo().FindAllSticker(l.ctx, userID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	for _, s := range stickers {
		stickerList = append(stickerList, types.StickerInfo{
			StickerID:   s.Uuid,
			StickerName: s.StickerName,
			StickerThum: s.StickerThum,
		})
	}
	return &types.GetUserStickerResp{
		Code:     uint(http.StatusOK),
		Stickers: stickerList,
	}, nil
}

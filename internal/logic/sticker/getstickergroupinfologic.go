package sticker

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStickerGroupInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStickerGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupInfoLogic {
	return &GetStickerGroupInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStickerGroupInfoLogic) GetStickerGroupInfo(req *types.GetStickerInfoReq) (resp *types.GetStickerInfoResp, err error) {
	// todo: add your logic here and delete this line
	sticker, err := l.svcCtx.DAO.FindOneStickerGroupByStickerUUID(l.ctx, req.StickerUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.STICKER_NOT_EXIST)
		}

		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	return &types.GetStickerInfoResp{
		Code: http.StatusOK,
		StickerInfo: types.StickerInfo{
			StickerID:   sticker.Uuid,
			StickerName: sticker.StickerName,
			StickerThum: sticker.StickerThum,
		},
	}, nil
}

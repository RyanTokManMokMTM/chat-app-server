package stickerservicelogic

import (
	"api/app/common/errx"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"net/http"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStickerGroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStickerGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupInfoLogic {
	return &GetStickerGroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStickerGroupInfoLogic) GetStickerGroupInfo(in *core.GetStickerInfoReq) (*core.GetStickerInfoResp, error) {
	// todo: add your logic here and delete this line
	sticker, err := l.svcCtx.DAO.FindOneStickerGroupByStickerUUID(l.ctx, in.StickerGroupUUID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STICKER_NOT_EXIST), "sticker not exist")
		}
		return nil, err
	}

	return &core.GetStickerInfoResp{Code: http.StatusOK,
		Info: &core.StickerInfo{
			StickerId:   sticker.Uuid,
			StickerName: sticker.StickerName,
			StickerThum: sticker.StickerThum,
		}}, nil
}

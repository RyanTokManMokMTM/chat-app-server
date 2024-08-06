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

type GetStickerGroupResourcesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStickerGroupResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupResourcesLogic {
	return &GetStickerGroupResourcesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStickerGroupResourcesLogic) GetStickerGroupResources(in *core.GetStickerResourcesReq) (*core.GetStickerResourcesResp, error) {
	// todo: add your logic here and delete this line
	sticker, err := l.svcCtx.DAO.FindOneStickerGroupByStickerUUID(l.ctx, in.StickerGroupUUID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.STICKER_NOT_EXIST), "sticker not exist")
		}
		return nil, err
	}

	resources := make([]string, 0)
	for _, stickerPath := range sticker.Resources {
		resources = append(resources, stickerPath.Path)
	}

	return &core.GetStickerResourcesResp{
		Code:      http.StatusOK,
		StickerId: sticker.Uuid,
		Paths:     resources,
	}, nil
}

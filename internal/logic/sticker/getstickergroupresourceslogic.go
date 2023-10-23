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

type GetStickerGroupResourcesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStickerGroupResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupResourcesLogic {
	return &GetStickerGroupResourcesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStickerGroupResourcesLogic) GetStickerGroupResources(req *types.GetStickerGroupReq) (resp *types.GetStickerGroupResp, err error) {
	// todo: add your logic here and delete this line
	sticker, err := l.svcCtx.DAO.FindOneStickerGroupByStickerUUID(l.ctx, req.StickerGroupUUID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.STICKER_NOT_EXIST)
		}

		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	resources := make([]string, 0)
	for _, stickerPath := range sticker.Resources {
		resources = append(resources, stickerPath.Path)
	}
	return &types.GetStickerGroupResp{
		Code:          http.StatusOK,
		StickerId:     sticker.Uuid,
		ResourcesPath: resources,
	}, nil
}

package sticker

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStickerGroupListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStickerGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupListLogic {
	return &GetStickerGroupListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStickerGroupListLogic) GetStickerGroupList(req *types.GetStickerListReq) (resp *types.GetStickerListResp, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.DAO.GetStickerGroupList(l.ctx)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	stickerInfos := make([]types.StickerInfo, 0)
	for _, info := range list {
		stickerInfos = append(stickerInfos, types.StickerInfo{
			StickerID:   info.Uuid,
			StickerName: info.StickerName,
			StickerThum: info.StickerThum,
		})
	}
	return &types.GetStickerListResp{
		Code:     http.StatusOK,
		Stickers: stickerInfos,
	}, nil
}

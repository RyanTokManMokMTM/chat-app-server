package stickerservicelogic

import (
	"context"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStickerGroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStickerGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStickerGroupListLogic {
	return &GetStickerGroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStickerGroupListLogic) GetStickerGroupList(in *core.GetStickerListReq) (*core.GetStickerListResp, error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.DAO.GetStickerGroupList(l.ctx)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}
	stickerInfos := make([]*core.StickerInfo, 0)
	for _, info := range list {
		stickerInfos = append(stickerInfos, &core.StickerInfo{
			StickerId:   info.Uuid,
			StickerName: info.StickerName,
			StickerThum: info.StickerThum,
		})
	}
	return &core.GetStickerListResp{
		Code:    http.StatusOK,
		Sticker: stickerInfos,
	}, nil
}

package stickerservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStickerGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStickerGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStickerGroupLogic {
	return &CreateStickerGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStickerGroupLogic) CreateStickerGroup(in *core.CreateStickerGroupReq) (*core.CreateStickerGroupResp, error) {
	// todo: add your logic here and delete this line
	//RPC Transaction?
	stickerModel, err := l.svcCtx.DAO.InsertOneStickerGroup(l.ctx, in.StickerName)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	assetsData := make(map[string]*assets_api.StickerFileMap)
	for key, data := range in.StickerData {
		stickerData := make([]*assets_api.StickerData, 0)
		for _, i := range data.Data {
			stickerData = append(stickerData, &assets_api.StickerData{
				Name: i.Name,
				Data: i.Data,
			})
		}
		assetsData[key] = &assets_api.StickerFileMap{
			Data: stickerData,
		}
	}

	rpcResp, rpcErr := l.svcCtx.AssetsRPC.UploadStickerGroup(l.ctx, &assets_api.UploadStickerGroupReq{
		StickerId:   stickerModel.Uuid,
		StickerName: in.StickerName,
		StickerData: assetsData,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	filePaths := make([]string, 0)
	for _, stickerMap := range rpcResp.StickerInfos {
		key := stickerMap.GetName()
		if key == "thum" {
			if len(stickerMap.GetPaths()) > 1 {
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "too many thum sticker")
			}
			stickerModel.StickerThum = stickerMap.GetPaths()[0]
			if err := l.svcCtx.DAO.UpdateOneStickerGroup(l.ctx, stickerModel); err != nil {
				logx.WithContext(l.ctx).Error(err)
				return nil, err
			}
		} else {
			for _, path := range stickerMap.GetPaths() {
				filePaths = append(filePaths, path)
			}
		}

	}
	//
	if err := l.svcCtx.DAO.InsertStickerListIntoGroup(l.ctx, stickerModel, filePaths); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.CreateStickerGroupResp{
		Code:             int32(errx.SUCCESS),
		StickerGroupUUID: stickerModel.Uuid,
	}, nil
}

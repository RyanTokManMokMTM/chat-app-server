package stickerservicelogic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/variable"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"github.com/ryantokmanmokmtm/chat-app-server/app/internal/models"
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
	stickerModelUUID := uuid.NewString()

	assetsData := make(map[string]*assets_api.StickerFileMap)
	resourcesPaths := make(map[string][]string)
	for key, data := range in.StickerData {
		stickerData := make([]*assets_api.StickerData, 0)
		resources := make([]string, 0)
		for _, i := range data.Data {
			name := uploadx.RandomFileName(i.Name)
			stickerData = append(stickerData, &assets_api.StickerData{
				Name: name,
				Data: i.Data,
			})
			resources = append(resources, fmt.Sprintf("/%s/%s", stickerModelUUID, name))
		}
		assetsData[key] = &assets_api.StickerFileMap{
			Data: stickerData,
		}
		resourcesPaths[key] = resources
	}

	_, rpcErr := l.svcCtx.AssetsRPC.UploadStickerGroup(l.ctx, &assets_api.UploadStickerGroupReq{
		StickerId:   stickerModelUUID,
		StickerData: assetsData,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	stickerModel := new(models.Sticker)
	stickerModel.Uuid = stickerModelUUID
	stickerModel.StickerName = in.StickerName
	stickerModel.StickerThum = resourcesPaths[variable.STICKER_THUM][0]

	stickerResources := make([]models.StickerResource, 0)
	for _, p := range resourcesPaths[variable.STICKER_RESOURCES] {
		stickerResources = append(stickerResources, models.StickerResource{
			Path: p,
		})
	}
	stickerModel.Resources = stickerResources
	_, err := l.svcCtx.DAO.InsertOneStickerGroupWithResources(l.ctx, stickerModel)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.CreateStickerGroupResp{
		Code:             int32(errx.SUCCESS),
		StickerGroupUUID: stickerModel.Uuid,
	}, nil
}

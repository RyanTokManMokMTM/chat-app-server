package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/util"
	"os"
	"path"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadStickerGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadStickerGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadStickerGroupLogic {
	return &UploadStickerGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadStickerGroupLogic) UploadStickerGroup(in *assets_api.UploadStickerGroupReq) (*assets_api.UploadStickerGroupResp, error) {
	// todo: add your logic here and delete this line
	root := util.GetRootDir()
	stickerDir := path.Join(root, fmt.Sprintf("%s/sticker", l.svcCtx.Config.ResourcesPath))

	if _, err := os.Stat(stickerDir); os.IsNotExist(err) {
		err := os.Mkdir(stickerDir, 0777)
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}
	}

	//stickerModel, err := l.svcCtx.DAO.InsertOneStickerGroup(l.ctx, in.StickerName)
	//if err != nil {
	//	logx.WithContext(l.ctx).Error(err)
	//	return nil, err
	//}
	//TODO: Create an sticker file
	stickerGroupDir := path.Join(root, fmt.Sprintf("%s/sticker/%s", l.svcCtx.Config.ResourcesPath, in.StickerId))
	if err := os.Mkdir(stickerGroupDir, 0777); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	fileMap := in.StickerData
	filePaths := make([]*assets_api.StickerUploadedInfo, 0)
	for key, stickerMap := range fileMap {
		if key == "thum" {
			thumFile := stickerMap.Data[0]

			p, err := uploadx.SaveFileWithRandomName(thumFile.Data, thumFile.Name, stickerGroupDir)
			if err != nil {
				logx.WithContext(l.ctx).Error(err)
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED), "Sticker upload failed")
			}

			filePaths = append(filePaths, &assets_api.StickerUploadedInfo{
				Name: key,
				Paths: []string{
					p,
				},
			})
		} else {
			stickerPaths := make([]string, 0)
			for _, stickerData := range stickerMap.Data {
				p, err := uploadx.SaveFileWithRandomName(stickerData.Data, stickerData.Name, stickerGroupDir)
				if err != nil {
					logx.WithContext(l.ctx).Error(err)
					return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
				}
				stickerPaths = append(stickerPaths, fmt.Sprintf("/%s%s", in.StickerId, p))
			}
			filePaths = append(filePaths, &assets_api.StickerUploadedInfo{
				Name:  key,
				Paths: stickerPaths,
			})
		}
	}

	return &assets_api.UploadStickerGroupResp{
		Code:         int32(errx.SUCCESS),
		StickerInfos: filePaths,
	}, nil

}

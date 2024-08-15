package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/util"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/variable"
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

	//TODO: Create an sticker file
	stickerGroupDir := path.Join(root, fmt.Sprintf("%s/sticker/%s", l.svcCtx.Config.ResourcesPath, in.StickerId))
	if err := os.Mkdir(stickerGroupDir, 0777); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	fileMap := in.StickerData
	for key, stickerMap := range fileMap {
		if key == variable.STICKER_THUM {
			thumFile := stickerMap.Data[0]

			_, err := uploadx.SaveFileWithName(thumFile.Data, thumFile.Name, stickerGroupDir)
			if err != nil {
				logx.WithContext(l.ctx).Error(err)
				go func() {
					err := os.RemoveAll(stickerGroupDir)
					logx.Errorf("Remove sticker dir error %+v", err)
				}()
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED), "Sticker upload failed")
			}

		} else if key == variable.STICKER_RESOURCES {
			for _, stickerData := range stickerMap.Data {
				_, err := uploadx.SaveFileWithName(stickerData.Data, stickerData.Name, stickerGroupDir)
				if err != nil {
					logx.WithContext(l.ctx).Error(err)
					go func() {
						err := os.RemoveAll(stickerGroupDir)
						logx.Errorf("Remove sticker dir error %+v", err)
					}()
					return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
				}
			}
		}
	}
	return &assets_api.UploadStickerGroupResp{
		Code: int32(errx.SUCCESS),
	}, nil

}

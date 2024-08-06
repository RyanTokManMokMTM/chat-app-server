package stickerservicelogic

import (
	"api/app/common/errx"
	"api/app/common/uploadx"
	"api/app/common/util"
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

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
	root := util.GetRootDir()
	stickerDir := path.Join(root, fmt.Sprintf("%s/sticker", l.svcCtx.Config.ResourcesPath))

	if _, err := os.Stat(stickerDir); os.IsNotExist(err) {
		err := os.Mkdir(stickerDir, 0777)
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}
	}

	stickerModel, err := l.svcCtx.DAO.InsertOneStickerGroup(l.ctx, in.StickerName)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}
	//TODO: Create an sticker file
	stickerGroupDir := path.Join(root, fmt.Sprintf("%s/sticker/%s", l.svcCtx.Config.ResourcesPath, stickerModel.Uuid))
	if err := os.Mkdir(stickerGroupDir, 0777); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	fileMap := in.StickerData
	filePaths := make([]string, 0)
	for key, stickerMap := range fileMap {
		if key == "thum" {
			thumFile := stickerMap.Data[0]

			p, err := uploadx.SaveFileWithRandomName(thumFile.Data, thumFile.Name, stickerGroupDir)
			if err != nil {
				logx.WithContext(l.ctx).Error(err)
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED), "Sticker upload failed")
			}

			stickerModel.StickerThum = fmt.Sprintf("/%s%s", stickerModel.Uuid, p)
			if err := l.svcCtx.DAO.UpdateOneStickerGroup(l.ctx, stickerModel); err != nil {
				logx.WithContext(l.ctx).Error(err)
				return nil, err
			}
		} else {
			for _, stickerData := range stickerMap.Data {
				p, err := uploadx.SaveFileWithRandomName(stickerData.Data, stickerData.Name, stickerGroupDir)
				if err != nil {
					logx.WithContext(l.ctx).Error(err)
					return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
				}
				filePaths = append(filePaths, fmt.Sprintf("/%s%s", stickerModel.Uuid, p))
			}
		}

	}
	logx.Info(filePaths)

	if err := l.svcCtx.DAO.InsertStickerListIntoGroup(l.ctx, stickerModel, filePaths); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.CreateStickerGroupResp{
		Code:             int32(errx.SUCCESS),
		StickerGroupUUID: stickerModel.Uuid,
	}, nil
}

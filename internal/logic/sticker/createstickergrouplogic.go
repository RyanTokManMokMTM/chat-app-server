package sticker

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"os"
)

type CreateStickerGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewCreateStickerGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *CreateStickerGroupLogic {
	return &CreateStickerGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *CreateStickerGroupLogic) CreateStickerGroup(req *types.CreateStickerGroupReq) (resp *types.CreateStickerGroupResp, err error) {
	// todo: add your logic here and delete this line
	stickerDir := fmt.Sprintf("%s/sticker", l.svcCtx.Config.ResourcesPath)
	if _, err := os.Stat(stickerDir); os.IsNotExist(err) {
		err := os.Mkdir(stickerDir, 666)
		if err != nil {
			return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
		}
	}

	stickerModel, err := l.svcCtx.DAO.InsertOneStickerGroup(l.ctx, req.StickerName)

	//TODO: Create an sticker file
	stickerGroupDir := fmt.Sprintf("%s/sticker/%s", l.svcCtx.Config.ResourcesPath, stickerModel.Uuid)
	if err := os.Mkdir(stickerGroupDir, 0777); err != nil {
		return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
	}
	fileMap := l.r.MultipartForm.File
	filePaths := make([]string, 0)
	for key, files := range fileMap {
		if key == "thum" {
			thumFile := files[0]
			f, err := thumFile.Open()
			if err != nil {
				return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
			}

			path, err := uploadx.UploadFileWithCustomName(f, thumFile, uuid.NewString(), stickerGroupDir)
			if err != nil {
				return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
			}
			f.Close()

			stickerModel.StickerThum = fmt.Sprintf("/%s%s", stickerModel.Uuid, path)
			if err := l.svcCtx.DAO.UpdateOneStickerGroup(l.ctx, stickerModel); err != nil {
				return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
			}
		} else {
			for _, file := range files {

				f, err := file.Open()
				if err != nil {
					return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
				}

				path, err := uploadx.UploadFileWithCustomName(f, file, uuid.NewString(), stickerGroupDir)
				if err != nil {
					return nil, errx.NewCustomError(errx.FILE_UPLOAD_FAILED, err.Error())
				}

				filePaths = append(filePaths, fmt.Sprintf("/%s%s", stickerModel.Uuid, path))
				f.Close()
			}
		}

	}
	logx.Info(filePaths)

	if err := l.svcCtx.DAO.InsertStickerListIntoGroup(l.ctx, stickerModel, filePaths); err != nil {
		return nil, errx.NewCustomError(errx.STORY_CREATED_FAILED, err.Error())
	}

	return &types.CreateStickerGroupResp{
		Code:             uint(http.StatusOK),
		StickerGroupUUID: stickerModel.Uuid,
	}, nil
}

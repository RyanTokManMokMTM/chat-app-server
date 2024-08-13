package sticker

import (
	"bytes"
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
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
	//uploadx.UploadFileWithCustomName(f, thumFile, uuid.NewString(), stickerGroupDir)
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	fileMap := l.r.MultipartForm.File
	stickerMap := make(map[string]*core.StickerFileMap)

	var nameExpo = func(file *multipart.FileHeader) (*core.StickerData, error) {
		f, err := file.Open()
		if err != nil {
			return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
		}
		defer f.Close()
		fileType := strings.Split(file.Filename, ".")[1]
		fileName := fmt.Sprintf("%s.%s", strings.ToLower(uuid.NewString()), fileType)

		buffer := bytes.NewBuffer(nil)
		_, err = io.Copy(buffer, f)
		if err != nil {
			return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
		}
		return &core.StickerData{
			Name: fileName,
			Data: buffer.Bytes(),
		}, nil
	}

	for key, files := range fileMap {
		if key == "thum" {
			thumFile := files[0]
			data, err := nameExpo(thumFile)
			if err != nil {
				return nil, err
			}
			stickerMap[key] = &core.StickerFileMap{
				Data: []*core.StickerData{
					data,
				},
			}

		} else {
			stickerDatas := make([]*core.StickerData, 0)
			for _, stickerFile := range files {
				data, err := nameExpo(stickerFile)
				if err != nil {
					return nil, err
				}
				stickerDatas = append(stickerDatas, data)
			}
			stickerMap[key] = &core.StickerFileMap{
				Data: stickerDatas,
			}
		}
	}
	rpcResp, rpcErr := l.svcCtx.StickerService.CreateStickerGroup(l.ctx, &core.CreateStickerGroupReq{
		UserId:      int32(userID),
		StickerName: req.StickerName,
		StickerData: stickerMap,
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.CreateStickerGroupResp{
		Code:             uint(rpcResp.Code),
		StickerGroupUUID: rpcResp.StickerGroupUUID,
	}, nil
}

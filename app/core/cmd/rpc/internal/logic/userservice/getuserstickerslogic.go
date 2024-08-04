package userservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserStickersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserStickersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserStickersLogic {
	return &GetUserStickersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserStickersLogic) GetUserStickers(in *core.GetUserStickerReq) (*core.GetUserStickerResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)

	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	stickerList := make([]*core.StickerInfo, 0)
	stickers, err := l.svcCtx.DAO.FindAllSticker(l.ctx, userID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	for _, s := range stickers {
		stickerList = append(stickerList, &core.StickerInfo{
			StickerId:   s.Uuid,
			StickerName: s.StickerName,
			StickerThum: s.StickerThum,
		})
	}

	return &core.GetUserStickerResp{
		Code:        int32(errx.SUCCESS),
		StickerInfo: stickerList,
	}, nil
}
package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
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
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error: %+v", err)
		}
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	stickerList := make([]*core.StickerInfo, 0)
	stickers, err := l.svcCtx.DAO.FindAllSticker(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
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

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

type IsStickerExistLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsStickerExistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsStickerExistLogic {
	return &IsStickerExistLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsStickerExistLogic) IsStickerExist(in *core.IsStickerExistReq) (*core.IsStickerExistResp, error) {
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

	var isExist = false
	sticker, err := l.svcCtx.DAO.FindOneStickerFromUser(l.ctx, userID, in.StickerUUID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if sticker != nil && sticker.Id != 0 {
		isExist = true
	}

	return &core.IsStickerExistResp{
		Code:    int32(errx.SUCCESS),
		IsExist: isExist,
	}, nil
}

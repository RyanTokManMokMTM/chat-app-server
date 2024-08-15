package userservicelogic

import (
	"context"

	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/assetrpc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadUserAvatarLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadUserAvatarLogic) UploadUserAvatar(in *core.UploadUserAvatarReq) (*core.UploadUserAvatarResp, error) {
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

	rpcResp, rpcErr := l.svcCtx.AssetsRPC.UploadImageByByte(l.ctx, &assetrpc.UploadImageReq{
		Format: in.Format,
		Data:   in.Data,
	})

	//name, err := uploadx.SaveBytesIntoFile(in.FileName, in.Data, l.svcCtx.Config.ResourcesPath)
	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}
	//
	//path := fmt.Sprintf("%v", name)
	err = l.svcCtx.DAO.UpdateUserAvatar(l.ctx, userID, rpcResp.Path)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	return &core.UploadUserAvatarResp{
		Code: int32(errx.SUCCESS),
		Path: rpcResp.Path,
	}, nil
}

package userservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/assetrpc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UploadUserCoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadUserCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserCoverLogic {
	return &UploadUserCoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadUserCoverLogic) UploadUserCover(in *core.UploadUserCoverReq) (*core.UploadUserCoverResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist ,error : %+v", err)
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
	if err = l.svcCtx.DAO.UpdateUserCover(l.ctx, userID, rpcResp.Path); err != nil {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	return &core.UploadUserCoverResp{
		Code: int32(errx.SUCCESS),
		Path: rpcResp.Path,
	}, nil
}

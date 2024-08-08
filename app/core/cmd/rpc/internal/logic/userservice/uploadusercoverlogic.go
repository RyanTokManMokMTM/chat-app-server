package userservicelogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
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

	name, err := uploadx.SaveBytesIntoFile(in.FileName, in.Data, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.FILE_UPLOAD_FAILED), "upload file failed, error: %+v", err)
	}

	path := fmt.Sprintf("%v", name)
	if err = l.svcCtx.DAO.UpdateUserCover(l.ctx, userID, path); err != nil {
		logx.WithContext(l.ctx).Errorf("Error : %+v", err)
		return nil, err
	}

	return &core.UploadUserCoverResp{
		Code: int32(errx.SUCCESS),
		Path: path,
	}, nil
}

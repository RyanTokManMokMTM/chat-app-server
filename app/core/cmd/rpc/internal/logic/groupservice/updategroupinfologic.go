package groupservicelogic

import (
	"api/app/common/errx"
	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateGroupInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateGroupInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateGroupInfoLogic {
	return &UpdateGroupInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateGroupInfoLogic) UpdateGroupInfo(in *core.UpdateGroupInfoReq) (*core.UpdateGroupInfoResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}

	group, err := l.svcCtx.DAO.FindOneGroup(l.ctx, uint(in.GroupId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.GROUP_NOT_EXIST), "group not exist")
		}
		return nil, err
	}

	if group.GroupLead != userID {
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NO_GROUP_AUTHORITY), "no group authority")
	}

	if err := l.svcCtx.DAO.UpdateOneGroup(l.ctx, group.Id, in.GroupName, in.GroupDesc); err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	return &core.UpdateGroupInfoResp{
		Code: uint32(errx.SUCCESS),
	}, nil
}

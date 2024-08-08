package groupservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupLogic) CreateGroup(in *core.CreateGroupReq) (*core.CreateGroupResp, error) {
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
	avatar := "/defaultGroup.jpg"

	if len(in.AvatarData) != 0 {
		path, err := uploadx.SaveImageByBase64(string(in.AvatarData), "jpg", l.svcCtx.Config.ResourcesPath)
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			return nil, err
		}
		avatar = path
	}

	group, err := l.svcCtx.DAO.InsertOneGroup(l.ctx, in.GroupName, avatar, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	//TODO: Put it into MQ, and sent it by MQ
	//sysMessage := fmt.Sprintf("%s created the group.", u.NickName)
	//if len(req.GroupMembers) > 0 {
	//	var members []string
	//	for _, memberID := range req.GroupMembers {
	//		err := l.svcCtx.DAO.InsertOneGroupMember(l.ctx, group.Id, memberID)
	//		if err != nil {
	//			logx.Error(err.Error())
	//			continue
	//		}
	//
	//		mem, err := l.svcCtx.DAO.FindOneUser(l.ctx, memberID)
	//		if err != nil {
	//			logx.Error(err.Error())
	//			continue
	//		}
	//
	//		members = append(members, mem.NickName)
	//	}
	//	sysMessage = fmt.Sprintf("%s added %s to the group.", u.NickName, strings.Join(members, ","))
	//}

	//go func() {
	//	logx.Info("sending a system message")
	//	ws.SendGroupSystemNotification(u.Uuid, group.Uuid, sysMessage)
	//}()

	return &core.CreateGroupResp{
		Code:      uint32(errx.SUCCESS),
		GroupUUID: group.Uuid,
		Avatar:    group.GroupAvatar,
	}, nil
}

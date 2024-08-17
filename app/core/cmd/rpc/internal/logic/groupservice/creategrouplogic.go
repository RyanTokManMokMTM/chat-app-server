package groupservicelogic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/assetrpc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/redisx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/redisx/types"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/util"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"gorm.io/gorm"
	"strings"

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
	u, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}
	avatar := "/default/defaultGroup.jpg"

	if len(in.AvatarData) != 0 {
		imgFormat := util.ExtractImgTypeFromBase64(string(in.AvatarData))
		if imgFormat == "" {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "Avatar data format incorrect")
		}

		rpcResp, rpcErr := l.svcCtx.AssetsRPC.UploadImageByBase64(l.ctx, &assetrpc.UploadImageReq{
			Format: imgFormat,
			Data:   in.AvatarData,
		})
		if rpcErr != nil {
			logx.WithContext(l.ctx).Error(rpcErr)
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "save file error, err: %+v", rpcErr)
		}
		avatar = rpcResp.Path
	}

	group, err := l.svcCtx.DAO.InsertOneGroup(l.ctx, in.GroupName, avatar, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	//TODO: Put it into MQ, and sent it by MQ
	sysMessage := fmt.Sprintf("%s created the group.", u.NickName)

	if len(in.GroupMembers) > 0 {
		var members []string
		for _, memberID := range in.GroupMembers {
			err := l.svcCtx.DAO.InsertOneGroupMember(l.ctx, group.Id, uint(memberID))
			if err != nil {
				logx.Error(err.Error())
				continue
			}
			logx.Info(memberID)
			mem, err := l.svcCtx.DAO.FindOneUser(l.ctx, uint(memberID))
			if err != nil {
				logx.Error(err.Error())
				continue
			}

			members = append(members, mem.NickName)
		}
		sysMessage = fmt.Sprintf("%s added %s to the group.", u.NickName, strings.Join(members, ","))
	}

	go func() {
		logx.Info("sending a system message", sysMessage)
		redisContext := context.Background()
		redisx.SendMessageToChannel(l.svcCtx.RedisCli, redisContext, redisx.NOTIFICATION_CHANNEL, types.NotificationMessage{
			To:      group.Uuid,
			From:    u.Uuid,
			Content: sysMessage,
		})
		//ws.SendGroupSystemNotification(u.Uuid, group.Uuid, sysMessage)
	}()

	return &core.CreateGroupResp{
		Code:      uint32(errx.SUCCESS),
		GroupUUID: group.Uuid,
		Avatar:    group.GroupAvatar,
	}, nil
}

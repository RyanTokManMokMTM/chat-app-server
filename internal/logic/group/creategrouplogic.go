package group

import (
	"context"
	"errors"
	"fmt"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/handler/ws"
	"gorm.io/gorm"
	"net/http"
	"strings"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.CreateGroupReq) (resp *types.CreateGroupResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	u, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}
	var avatar string = "/defaultGroup.jpg"

	if len(req.GroupAvatar) != 0 {
		path, err := uploadx.UploadImageByBase64(req.GroupAvatar, "jpg", l.svcCtx.Config.ResourcesPath)
		if err != nil {
			return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
		}

		avatar = path
	}

	group, err := l.svcCtx.DAO.InsertOneGroup(l.ctx, req.GroupName, avatar, userID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	sysMessage := fmt.Sprintf("%s created the group.", u.NickName)
	if len(req.GroupMembers) > 0 {
		var members []string
		for _, memberID := range req.GroupMembers {
			err := l.svcCtx.DAO.InsertOneGroupMember(l.ctx, group.ID, memberID)
			if err != nil {
				logx.Error(err.Error())
				continue
			}

			mem, err := l.svcCtx.DAO.FindOneUser(l.ctx, memberID)
			if err != nil {
				logx.Error(err.Error())
				continue
			}

			members = append(members, mem.NickName)
		}
		sysMessage = fmt.Sprintf("%s added %s to the group.", u.NickName, strings.Join(members, ","))
	}

	go func() {
		logx.Info("sending a system message")
		ws.SendGroupSystemNotification(u.Uuid, group.Uuid, sysMessage)
	}()

	return &types.CreateGroupResp{
		Code:        uint(http.StatusOK),
		GroupUUID:   group.Uuid,
		GroupAvatar: group.GroupAvatar,
	}, nil
}

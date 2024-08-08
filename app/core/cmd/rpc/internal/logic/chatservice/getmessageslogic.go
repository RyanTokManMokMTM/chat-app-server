package chatservicelogic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/pagerx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/variable"
	"gorm.io/gorm"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessagesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessagesLogic {
	return &GetMessagesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessagesLogic) GetMessages(in *core.GetMessagesReq) (*core.GetMessagesResp, error) {
	// todo: add your logic here and delete this line
	var respMessages = make([]*core.MessageUser, 0)
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not found, error : %+v", err)
		}
		return nil, err
	}

	if in.MessageType == variable.MESSAGE_TYPE_USERCHAT {
		//TODO: Check User is friend
		_, err = l.svcCtx.DAO.FindOneFriend(l.ctx, userID, uint(in.SourceId))
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NOT_YET_FRIEND), "Not your friend, error : %+v", err)
			}
			return nil, err
		}
	} else if in.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
		//TODO: Check User is group member
		_, err := l.svcCtx.DAO.FindOneGroupMember(l.ctx, uint(in.SourceId), userID)
		if err != nil {
			logx.WithContext(l.ctx).Error(err)
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.Wrapf(errx.NewCustomErrCode(errx.NOT_JOIN_GROUP_YET), "not joing the group yet,error : %+v", err)
			}
			return nil, err
		}
	}

	pageLimit := pagerx.GetLimit(uint(in.Limit))
	messages, err := l.svcCtx.DAO.GetMessage(l.ctx, userID, uint(in.SourceId), uint(in.MessageType), int(pageLimit), uint(in.LatestId))
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, err
	}

	for _, msg := range messages {
		respMessages = append(respMessages, &core.MessageUser{
			MessageId:   uint32(msg.ID),
			FromId:      uint32(msg.FromUserID),
			ToId:        uint32(msg.ToUserID),
			Content:     msg.Content,
			ContentType: msg.ContentType,
			MessageType: uint32(msg.MessageType),
			Url:         msg.Url,
			FileName:    msg.FileName,
			FileSize:    uint32(msg.FileSize),
			StoryTime:   uint32(msg.ContentAvailableTime),
			CreatedAt:   uint32(uint(msg.CreatedAt.Unix())),
		})
	}

	return &core.GetMessagesResp{
		Code:     uint32(errx.SUCCESS),
		Messages: respMessages,
	}, nil
}

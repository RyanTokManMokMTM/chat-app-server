package message

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/pagerx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/variable"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessagesLogic {
	return &GetMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessagesLogic) GetMessages(req *types.GetMessagesReq) (resp *types.GetMessagesResp, err error) {
	// todo: add your logic here and delete this line
	var respMessages = make([]types.MessageUser, 0)
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	if req.MessageType == variable.MESSAGE_TYPE_USERCHAT {
		//TODO: Check User is friend
		if err := l.svcCtx.DAO.FindOneFriend(l.ctx, userID, req.SouceId); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errx.NewCustomErrCode(errx.NOT_YET_FRIEND)
			}
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
	} else if req.MessageType == variable.MESSAGE_TYPE_GROUPCHAT {
		//TODO: Check User is group member
		_, err := l.svcCtx.DAO.FindOneGroupMember(l.ctx, req.SouceId, userID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errx.NewCustomErrCode(errx.NOT_JOIN_GROUP_YET)
			}
			return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
		}
	}

	pageLimit := pagerx.GetLimit(req.Limit)
	messages, err := l.svcCtx.DAO.GetMessage(l.ctx, userID, req.SouceId, req.MessageType, int(pageLimit), req.LatestID)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	for _, msg := range messages {
		respMessages = append(respMessages, types.MessageUser{
			MessageID:   msg.ID,
			FromID:      msg.FromUserID,
			ToID:        msg.ToUserID,
			Content:     msg.Content,
			ContentType: msg.ContentType,
			MessageType: msg.MessageType,
			Url:         msg.Url,
			FileName:    msg.FileName,
			FileSize:    msg.FileSize,
			StoryTime:   msg.StoryTime,
			CreatedAt:   uint(msg.CreatedAt.Unix()),
		})
	}
	return &types.GetMessagesResp{
		Code:     http.StatusOK,
		Messages: respMessages,
	}, nil
}

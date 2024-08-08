package message

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/chatservice"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// get room messages by roomID
func NewGetMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessagesLogic {
	return &GetMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessagesLogic) GetMessages(req *types.GetMessagesReq) (resp *types.GetMessagesResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.ChatService.GetMessages(l.ctx, &chatservice.GetMessagesReq{
		UserId:      uint32(userID),
		MessageType: uint32(req.MessageType),
		SourceId:    uint32(req.SouceId),
		Limit:       uint32(req.Limit),
		LatestId:    uint32(req.LatestID),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	messages := make([]types.MessageUser, 0)
	for _, msg := range rpcResp.Messages {
		messages = append(messages, types.MessageUser{
			MessageID:   uint(msg.MessageId),
			FromID:      uint(msg.FromId),
			ToID:        uint(msg.ToId),
			Content:     msg.Content,
			ContentType: msg.ContentType,
			MessageType: uint(msg.MessageType),
			Url:         msg.Url,
			FileName:    msg.FileName,
			FileSize:    uint(msg.FileSize),
			StoryTime:   uint(msg.StoryTime),
			CreatedAt:   uint(msg.CreatedAt),
		})
	}
	return &types.GetMessagesResp{
		Code:     uint(rpcResp.Code),
		Messages: messages,
	}, nil
}

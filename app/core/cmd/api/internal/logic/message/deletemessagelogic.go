package message

import (
	"context"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/chatservice"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// Delete one message
func NewDeleteMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMessageLogic {
	return &DeleteMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMessageLogic) DeleteMessage(req *types.DeleteMessageReq) (resp *types.DeleteMessageResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	rpcResp, rpcErr := l.svcCtx.ChatService.DeleteMessage(l.ctx, &chatservice.DeleteMessageReq{
		UserId: uint32(userID),
		MsgId:  uint32(req.MesssageID),
	})
	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}
	return &types.DeleteMessageResp{
		Code: uint(rpcResp.Code),
	}, nil
}

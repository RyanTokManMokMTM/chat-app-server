package svc

import (
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/config"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/chatservice"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/friendservice"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/groupservice"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/stickerservice"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/storyservice"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/client/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserService    userservice.UserService
	FriendService  friendservice.FriendService
	GroupService   groupservice.GroupService
	StickerService stickerservice.StickerService
	StoryService   storyservice.StoryService
	ChatService    chatservice.ChatService
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:         c,
		UserService:    userservice.NewUserService(zrpc.MustNewClient(c.CoreRPC)),
		FriendService:  friendservice.NewFriendService(zrpc.MustNewClient(c.CoreRPC)),
		GroupService:   groupservice.NewGroupService(zrpc.MustNewClient(c.CoreRPC)),
		StickerService: stickerservice.NewStickerService(zrpc.MustNewClient(c.CoreRPC)),
		StoryService:   storyservice.NewStoryService(zrpc.MustNewClient(c.CoreRPC)),
		ChatService:    chatservice.NewChatService(zrpc.MustNewClient(c.CoreRPC)),
	}
}

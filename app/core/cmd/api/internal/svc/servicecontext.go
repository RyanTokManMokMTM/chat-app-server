package svc

import (
	"api/app/core/cmd/api/internal/config"
	"api/app/core/cmd/rpc/client/friendservice"
	"api/app/core/cmd/rpc/client/groupservice"
	"api/app/core/cmd/rpc/client/stickerservice"
	"api/app/core/cmd/rpc/client/storyservice"
	"api/app/core/cmd/rpc/client/userservice"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	UserService    userservice.UserService
	FriendService  friendservice.FriendService
	GroupService   groupservice.GroupService
	StickerService stickerservice.StickerService
	StoryService   storyservice.StoryService
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:         c,
		UserService:    userservice.NewUserService(zrpc.MustNewClient(c.CoreRPC)),
		FriendService:  friendservice.NewFriendService(zrpc.MustNewClient(c.CoreRPC)),
		GroupService:   groupservice.NewGroupService(zrpc.MustNewClient(c.CoreRPC)),
		StickerService: stickerservice.NewStickerService(zrpc.MustNewClient(c.CoreRPC)),
		StoryService:   storyservice.NewStoryService(zrpc.MustNewClient(c.CoreRPC)),
	}
}

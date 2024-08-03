package storyservicelogic

import (
	"context"

	"api/app/core/cmd/rpc/internal/svc"
	"api/app/core/cmd/rpc/types/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddStoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddStoryLogic {
	return &AddStoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddStoryLogic) AddStory(stream core.StoryService_AddStoryServer) error {
	// todo: add your logic here and delete this line

	return nil
}

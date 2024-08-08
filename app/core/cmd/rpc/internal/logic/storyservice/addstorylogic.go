package storyservicelogic

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

func (l *AddStoryLogic) AddStory(in *core.AddStoryReq) (*core.AddStoryResp, error) {
	// todo: add your logic here and delete this line
	userID := uint(in.UserId)
	_, err := l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.Wrapf(errx.NewCustomErrCode(errx.USER_NOT_EXIST), "user not exist, error : %+v", err)
		}
		return nil, err
	}

	//Upload Image
	path, err := uploadx.SaveBytesIntoFile(in.StoryImgFileName, in.Data, l.svcCtx.Config.ResourcesPath)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "save file error, err: %+v", err)
	}

	story, err := l.svcCtx.DAO.InsertOneStory(l.ctx, userID, path)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &core.AddStoryResp{
		Code: uint32(errx.SUCCESS),
		Info: &core.StoryInfo{
			StoryId:   uint32(story.Id),
			StoryUUID: story.Uuid.String(),
			StoryURL:  path,
		},
	}, nil
}

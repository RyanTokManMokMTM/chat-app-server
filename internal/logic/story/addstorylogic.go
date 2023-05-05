package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/common/uploadx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewAddStoryLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *AddStoryLogic {
	return &AddStoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *AddStoryLogic) AddStory(req *types.AddStoryReq) (resp *types.AddStoryResp, err error) {
	// todo: add your logic here and delete this line
	userID := ctxtool.GetUserIDFromCTX(l.ctx)
	_, err = l.svcCtx.DAO.FindOneUser(l.ctx, userID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.USER_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	//Upload Image
	path, err := uploadx.UploadFileFromRequest(l.r, 1024, "story_media", l.svcCtx.Config.ResourcesPath)
	if err != nil {
		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, err.Error())
	}

	mediaPath := "/" + path
	logx.Info("media url " + mediaPath)
	storyID, err := l.svcCtx.DAO.InsertOneStory(l.ctx, userID, mediaPath)
	if err != nil {
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &types.AddStoryResp{
		Code:    uint(http.StatusOK),
		StoryID: storyID,
	}, nil
}

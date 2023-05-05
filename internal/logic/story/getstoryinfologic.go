package story

import (
	"context"
	"errors"
	"github.com/ryantokmanmokmtm/chat-app-server/common/errx"
	"gorm.io/gorm"
	"net/http"

	"github.com/ryantokmanmokmtm/chat-app-server/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStoryInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetStoryInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStoryInfoLogic {
	return &GetStoryInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStoryInfoLogic) GetStoryInfo(req *types.GetStoryInfoByIdRep) (resp *types.GetStoryInfoByIdResp, err error) {
	// todo: add your logic here and delete this line
	s, err := l.svcCtx.DAO.FindOneStory(l.ctx, req.StoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errx.NewCustomErrCode(errx.STORY_NOT_EXIST)
		}
		return nil, errx.NewCustomError(errx.DB_ERROR, err.Error())
	}

	return &types.GetStoryInfoByIdResp{
		Code:          uint(http.StatusOK),
		StoryID:       s.Id,
		StoryMediaURL: s.StoryMediaPath,
		CreateAt:      uint(s.CreatedAt.Unix()),
	}, nil
}

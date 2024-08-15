package story

import (
	"bytes"
	"context"
	"github.com/pkg/errors"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/ctxtool"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/errx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/common/uploadx"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/rpc/types/core"
	"io"
	"net/http"
	"strings"

	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/core/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddStoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

// Create a new instance story
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

	file, header, err := l.r.FormFile(uploadx.StoryMediaField)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.REQ_PARAM_ERROR), "Story media not exist")
	}

	buffer := bytes.NewBuffer(nil)
	fileName := header.Filename
	_, err = io.Copy(buffer, file)
	if err != nil {
		logx.WithContext(l.ctx).Error(err)
		return nil, errors.Wrapf(errx.NewCustomErrCode(errx.SERVER_COMMON_ERROR), "error:  %+v", err)
	}

	formatInfo := strings.Split(fileName, ".")
	if len(formatInfo) < 2 {
		return nil, errx.NewCustomErrCode(errx.REQ_PARAM_ERROR)
	}

	rpcResp, rpcErr := l.svcCtx.StoryService.AddStory(l.ctx, &core.AddStoryReq{
		UserId:      uint32(userID),
		StoryFormat: formatInfo[1],
		Data:        buffer.Bytes(),
	})

	if rpcErr != nil {
		logx.WithContext(l.ctx).Error(rpcErr)
		return nil, rpcErr
	}

	return &types.AddStoryResp{
		Code: uint(rpcResp.Code),
		Info: types.StoryInfo{
			StoryID:       uint(rpcResp.Info.StoryId),
			StoryUUID:     rpcResp.Info.StoryUUID,
			StoryMediaURL: rpcResp.Info.StoryURL,
		},
	}, nil
}

package file

//
//import (
//	"context"
//	"github.com/ryantokmanmok/chat-app-server/common/errx"
//	"github.com/ryantokmanmok/chat-app-server/common/uploadx"
//	"net/http"
//
//	"github.com/ryantokmanmok/chat-app-server/internal/svc"
//	"github.com/ryantokmanmok/chat-app-server/internal/types"
//
//	"github.com/zeromicro/go-zero/core/logx"
//)
//
//type UploadDocumentLogic struct {
//	logx.Logger
//	ctx    context.Context
//	svcCtx *svc.ServiceContext
//}
//
//func NewUploadDocumentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadDocumentLogic {
//	return &UploadDocumentLogic{
//		Logger: logx.WithContext(ctx),
//		ctx:    ctx,
//		svcCtx: svcCtx,
//	}
//}
//
//func (l *UploadDocumentLogic) UploadDocument(req *types.UploadFildReq) (resp *types.UploadFileResp, err error) {
//	// todo: add your logic here and delete this line
//	if len(req.Data) == 0 {
//		return nil, errx.NewCustomError(errx.REQ_PARAM_ERROR, "data can't be empty")
//	}
//
//	path, err := uploadx.UploadFileByBase64(req.Data, req.FileName, l.svcCtx.Config.ResourcesPath)
//	if err != nil {
//		return nil, errx.NewCustomError(errx.SERVER_COMMON_ERROR, err.Error())
//	}
//	return &types.UploadFileResp{
//		Code: uint(http.StatusOK),
//		Path: path,
//	}, nil
//
//}

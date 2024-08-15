// Code generated by goctl. DO NOT EDIT.
// Source: assets.proto

package server

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/logic"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/internal/svc"
	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"
)

type AssetRPCServer struct {
	svcCtx *svc.ServiceContext
	assets_api.UnimplementedAssetRPCServer
}

func NewAssetRPCServer(svcCtx *svc.ServiceContext) *AssetRPCServer {
	return &AssetRPCServer{
		svcCtx: svcCtx,
	}
}

func (s *AssetRPCServer) UploadImageByBase64(ctx context.Context, in *assets_api.UploadImageReq) (*assets_api.UploadImageResp, error) {
	l := logic.NewUploadImageByBase64Logic(ctx, s.svcCtx)
	return l.UploadImageByBase64(in)
}

func (s *AssetRPCServer) UploadFile(ctx context.Context, in *assets_api.UploadFileReq) (*assets_api.UploadFileResp, error) {
	l := logic.NewUploadFileLogic(ctx, s.svcCtx)
	return l.UploadFile(in)
}

func (s *AssetRPCServer) UploadImageByByte(ctx context.Context, in *assets_api.UploadImageReq) (*assets_api.UploadImageResp, error) {
	l := logic.NewUploadImageByByteLogic(ctx, s.svcCtx)
	return l.UploadImageByByte(in)
}

func (s *AssetRPCServer) UploadStickerGroup(ctx context.Context, in *assets_api.UploadStickerGroupReq) (*assets_api.UploadStickerGroupResp, error) {
	l := logic.NewUploadStickerGroupLogic(ctx, s.svcCtx)
	return l.UploadStickerGroup(in)
}

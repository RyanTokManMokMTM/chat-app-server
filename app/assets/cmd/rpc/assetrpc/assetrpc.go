// Code generated by goctl. DO NOT EDIT.
// Source: assets.proto

package assetrpc

import (
	"context"

	"github.com/ryantokmanmokmtm/chat-app-server/app/assets/cmd/rpc/types/assets_api"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	StickerData            = assets_api.StickerData
	StickerFileMap         = assets_api.StickerFileMap
	StickerUploadedInfo    = assets_api.StickerUploadedInfo
	UploadFileReq          = assets_api.UploadFileReq
	UploadFileResp         = assets_api.UploadFileResp
	UploadImageReq         = assets_api.UploadImageReq
	UploadImageResp        = assets_api.UploadImageResp
	UploadStickerGroupReq  = assets_api.UploadStickerGroupReq
	UploadStickerGroupResp = assets_api.UploadStickerGroupResp

	AssetRPC interface {
		UploadImage(ctx context.Context, in *UploadImageReq, opts ...grpc.CallOption) (*UploadImageResp, error)
		UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*UploadFileResp, error)
		UploadStickerGroup(ctx context.Context, in *UploadStickerGroupReq, opts ...grpc.CallOption) (*UploadStickerGroupResp, error)
	}

	defaultAssetRPC struct {
		cli zrpc.Client
	}
)

func NewAssetRPC(cli zrpc.Client) AssetRPC {
	return &defaultAssetRPC{
		cli: cli,
	}
}

func (m *defaultAssetRPC) UploadImage(ctx context.Context, in *UploadImageReq, opts ...grpc.CallOption) (*UploadImageResp, error) {
	client := assets_api.NewAssetRPCClient(m.cli.Conn())
	return client.UploadImage(ctx, in, opts...)
}

func (m *defaultAssetRPC) UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*UploadFileResp, error) {
	client := assets_api.NewAssetRPCClient(m.cli.Conn())
	return client.UploadFile(ctx, in, opts...)
}

func (m *defaultAssetRPC) UploadStickerGroup(ctx context.Context, in *UploadStickerGroupReq, opts ...grpc.CallOption) (*UploadStickerGroupResp, error) {
	client := assets_api.NewAssetRPCClient(m.cli.Conn())
	return client.UploadStickerGroup(ctx, in, opts...)
}

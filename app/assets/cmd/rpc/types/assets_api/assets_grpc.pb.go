// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: cmd/rpc/proto/assets.proto

package assets_api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AssetRPC_UploadImage_FullMethodName        = "/stream.AssetRPC/UploadImage"
	AssetRPC_UploadFile_FullMethodName         = "/stream.AssetRPC/UploadFile"
	AssetRPC_UploadStickerGroup_FullMethodName = "/stream.AssetRPC/UploadStickerGroup"
)

// AssetRPCClient is the client API for AssetRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetRPCClient interface {
	UploadImage(ctx context.Context, in *UploadImageReq, opts ...grpc.CallOption) (*UploadImageResp, error)
	UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*UploadFileResp, error)
	UploadStickerGroup(ctx context.Context, in *UploadStickerGroupReq, opts ...grpc.CallOption) (*UploadStickerGroupResp, error)
}

type assetRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetRPCClient(cc grpc.ClientConnInterface) AssetRPCClient {
	return &assetRPCClient{cc}
}

func (c *assetRPCClient) UploadImage(ctx context.Context, in *UploadImageReq, opts ...grpc.CallOption) (*UploadImageResp, error) {
	out := new(UploadImageResp)
	err := c.cc.Invoke(ctx, AssetRPC_UploadImage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetRPCClient) UploadFile(ctx context.Context, in *UploadFileReq, opts ...grpc.CallOption) (*UploadFileResp, error) {
	out := new(UploadFileResp)
	err := c.cc.Invoke(ctx, AssetRPC_UploadFile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetRPCClient) UploadStickerGroup(ctx context.Context, in *UploadStickerGroupReq, opts ...grpc.CallOption) (*UploadStickerGroupResp, error) {
	out := new(UploadStickerGroupResp)
	err := c.cc.Invoke(ctx, AssetRPC_UploadStickerGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetRPCServer is the server API for AssetRPC service.
// All implementations must embed UnimplementedAssetRPCServer
// for forward compatibility
type AssetRPCServer interface {
	UploadImage(context.Context, *UploadImageReq) (*UploadImageResp, error)
	UploadFile(context.Context, *UploadFileReq) (*UploadFileResp, error)
	UploadStickerGroup(context.Context, *UploadStickerGroupReq) (*UploadStickerGroupResp, error)
	mustEmbedUnimplementedAssetRPCServer()
}

// UnimplementedAssetRPCServer must be embedded to have forward compatible implementations.
type UnimplementedAssetRPCServer struct {
}

func (UnimplementedAssetRPCServer) UploadImage(context.Context, *UploadImageReq) (*UploadImageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedAssetRPCServer) UploadFile(context.Context, *UploadFileReq) (*UploadFileResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
}
func (UnimplementedAssetRPCServer) UploadStickerGroup(context.Context, *UploadStickerGroupReq) (*UploadStickerGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadStickerGroup not implemented")
}
func (UnimplementedAssetRPCServer) mustEmbedUnimplementedAssetRPCServer() {}

// UnsafeAssetRPCServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetRPCServer will
// result in compilation errors.
type UnsafeAssetRPCServer interface {
	mustEmbedUnimplementedAssetRPCServer()
}

func RegisterAssetRPCServer(s grpc.ServiceRegistrar, srv AssetRPCServer) {
	s.RegisterService(&AssetRPC_ServiceDesc, srv)
}

func _AssetRPC_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetRPCServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetRPC_UploadImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetRPCServer).UploadImage(ctx, req.(*UploadImageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetRPC_UploadFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadFileReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetRPCServer).UploadFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetRPC_UploadFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetRPCServer).UploadFile(ctx, req.(*UploadFileReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AssetRPC_UploadStickerGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadStickerGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetRPCServer).UploadStickerGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AssetRPC_UploadStickerGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetRPCServer).UploadStickerGroup(ctx, req.(*UploadStickerGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AssetRPC_ServiceDesc is the grpc.ServiceDesc for AssetRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.AssetRPC",
	HandlerType: (*AssetRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadImage",
			Handler:    _AssetRPC_UploadImage_Handler,
		},
		{
			MethodName: "UploadFile",
			Handler:    _AssetRPC_UploadFile_Handler,
		},
		{
			MethodName: "UploadStickerGroup",
			Handler:    _AssetRPC_UploadStickerGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/rpc/proto/assets.proto",
}

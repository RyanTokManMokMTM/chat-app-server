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
	AssetRPC_UploadImage_FullMethodName = "/stream.AssetRPC/UploadImage"
	AssetRPC_UploadFile_FullMethodName  = "/stream.AssetRPC/UploadFile"
)

// AssetRPCClient is the client API for AssetRPC service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetRPCClient interface {
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (AssetRPC_UploadImageClient, error)
	UploadFile(ctx context.Context, opts ...grpc.CallOption) (AssetRPC_UploadFileClient, error)
}

type assetRPCClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetRPCClient(cc grpc.ClientConnInterface) AssetRPCClient {
	return &assetRPCClient{cc}
}

func (c *assetRPCClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (AssetRPC_UploadImageClient, error) {
	stream, err := c.cc.NewStream(ctx, &AssetRPC_ServiceDesc.Streams[0], AssetRPC_UploadImage_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &assetRPCUploadImageClient{stream}
	return x, nil
}

type AssetRPC_UploadImageClient interface {
	Send(*UploadImageReq) error
	CloseAndRecv() (*UploadImageResp, error)
	grpc.ClientStream
}

type assetRPCUploadImageClient struct {
	grpc.ClientStream
}

func (x *assetRPCUploadImageClient) Send(m *UploadImageReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *assetRPCUploadImageClient) CloseAndRecv() (*UploadImageResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadImageResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *assetRPCClient) UploadFile(ctx context.Context, opts ...grpc.CallOption) (AssetRPC_UploadFileClient, error) {
	stream, err := c.cc.NewStream(ctx, &AssetRPC_ServiceDesc.Streams[1], AssetRPC_UploadFile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &assetRPCUploadFileClient{stream}
	return x, nil
}

type AssetRPC_UploadFileClient interface {
	Send(*UploadFileReq) error
	CloseAndRecv() (*UploadFileResp, error)
	grpc.ClientStream
}

type assetRPCUploadFileClient struct {
	grpc.ClientStream
}

func (x *assetRPCUploadFileClient) Send(m *UploadFileReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *assetRPCUploadFileClient) CloseAndRecv() (*UploadFileResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadFileResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AssetRPCServer is the server API for AssetRPC service.
// All implementations must embed UnimplementedAssetRPCServer
// for forward compatibility
type AssetRPCServer interface {
	UploadImage(AssetRPC_UploadImageServer) error
	UploadFile(AssetRPC_UploadFileServer) error
	mustEmbedUnimplementedAssetRPCServer()
}

// UnimplementedAssetRPCServer must be embedded to have forward compatible implementations.
type UnimplementedAssetRPCServer struct {
}

func (UnimplementedAssetRPCServer) UploadImage(AssetRPC_UploadImageServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedAssetRPCServer) UploadFile(AssetRPC_UploadFileServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadFile not implemented")
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

func _AssetRPC_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AssetRPCServer).UploadImage(&assetRPCUploadImageServer{stream})
}

type AssetRPC_UploadImageServer interface {
	SendAndClose(*UploadImageResp) error
	Recv() (*UploadImageReq, error)
	grpc.ServerStream
}

type assetRPCUploadImageServer struct {
	grpc.ServerStream
}

func (x *assetRPCUploadImageServer) SendAndClose(m *UploadImageResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *assetRPCUploadImageServer) Recv() (*UploadImageReq, error) {
	m := new(UploadImageReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _AssetRPC_UploadFile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AssetRPCServer).UploadFile(&assetRPCUploadFileServer{stream})
}

type AssetRPC_UploadFileServer interface {
	SendAndClose(*UploadFileResp) error
	Recv() (*UploadFileReq, error)
	grpc.ServerStream
}

type assetRPCUploadFileServer struct {
	grpc.ServerStream
}

func (x *assetRPCUploadFileServer) SendAndClose(m *UploadFileResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *assetRPCUploadFileServer) Recv() (*UploadFileReq, error) {
	m := new(UploadFileReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AssetRPC_ServiceDesc is the grpc.ServiceDesc for AssetRPC service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AssetRPC_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "stream.AssetRPC",
	HandlerType: (*AssetRPCServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _AssetRPC_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "UploadFile",
			Handler:       _AssetRPC_UploadFile_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "cmd/rpc/proto/assets.proto",
}
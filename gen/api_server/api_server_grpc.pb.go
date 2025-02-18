// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api_server.proto

package api_server

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
	ExternalAPIService_CreateStrategy_FullMethodName = "/api.server.ExternalAPIService/CreateStrategy"
)

// ExternalAPIServiceClient is the client API for ExternalAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExternalAPIServiceClient interface {
	// 创建策略
	CreateStrategy(ctx context.Context, in *CreateStrategyRequest, opts ...grpc.CallOption) (*CreateStrategyResponse, error)
}

type externalAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExternalAPIServiceClient(cc grpc.ClientConnInterface) ExternalAPIServiceClient {
	return &externalAPIServiceClient{cc}
}

func (c *externalAPIServiceClient) CreateStrategy(ctx context.Context, in *CreateStrategyRequest, opts ...grpc.CallOption) (*CreateStrategyResponse, error) {
	out := new(CreateStrategyResponse)
	err := c.cc.Invoke(ctx, ExternalAPIService_CreateStrategy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExternalAPIServiceServer is the server API for ExternalAPIService service.
// All implementations must embed UnimplementedExternalAPIServiceServer
// for forward compatibility
type ExternalAPIServiceServer interface {
	// 创建策略
	CreateStrategy(context.Context, *CreateStrategyRequest) (*CreateStrategyResponse, error)
	mustEmbedUnimplementedExternalAPIServiceServer()
}

// UnimplementedExternalAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExternalAPIServiceServer struct {
}

func (UnimplementedExternalAPIServiceServer) CreateStrategy(context.Context, *CreateStrategyRequest) (*CreateStrategyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateStrategy not implemented")
}
func (UnimplementedExternalAPIServiceServer) mustEmbedUnimplementedExternalAPIServiceServer() {}

// UnsafeExternalAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExternalAPIServiceServer will
// result in compilation errors.
type UnsafeExternalAPIServiceServer interface {
	mustEmbedUnimplementedExternalAPIServiceServer()
}

func RegisterExternalAPIServiceServer(s grpc.ServiceRegistrar, srv ExternalAPIServiceServer) {
	s.RegisterService(&ExternalAPIService_ServiceDesc, srv)
}

func _ExternalAPIService_CreateStrategy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStrategyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExternalAPIServiceServer).CreateStrategy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExternalAPIService_CreateStrategy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExternalAPIServiceServer).CreateStrategy(ctx, req.(*CreateStrategyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExternalAPIService_ServiceDesc is the grpc.ServiceDesc for ExternalAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExternalAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.server.ExternalAPIService",
	HandlerType: (*ExternalAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateStrategy",
			Handler:    _ExternalAPIService_CreateStrategy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_server.proto",
}

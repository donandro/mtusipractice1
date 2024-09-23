// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: mobileapi/v1/mobileapi.proto

package grpc

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

// MobileApiServiceClient is the client API for MobileApiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MobileApiServiceClient interface {
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error)
}

type mobileApiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMobileApiServiceClient(cc grpc.ClientConnInterface) MobileApiServiceClient {
	return &mobileApiServiceClient{cc}
}

func (c *mobileApiServiceClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*GetUserByIdResponse, error) {
	out := new(GetUserByIdResponse)
	err := c.cc.Invoke(ctx, "/mobileapi.v1.MobileApiService/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MobileApiServiceServer is the server API for MobileApiService service.
// All implementations must embed UnimplementedMobileApiServiceServer
// for forward compatibility
type MobileApiServiceServer interface {
	GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error)
	mustEmbedUnimplementedMobileApiServiceServer()
}

// UnimplementedMobileApiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMobileApiServiceServer struct {
}

func (UnimplementedMobileApiServiceServer) GetUserById(context.Context, *GetUserByIdRequest) (*GetUserByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedMobileApiServiceServer) mustEmbedUnimplementedMobileApiServiceServer() {}

// UnsafeMobileApiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MobileApiServiceServer will
// result in compilation errors.
type UnsafeMobileApiServiceServer interface {
	mustEmbedUnimplementedMobileApiServiceServer()
}

func RegisterMobileApiServiceServer(s grpc.ServiceRegistrar, srv MobileApiServiceServer) {
	s.RegisterService(&MobileApiService_ServiceDesc, srv)
}

func _MobileApiService_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MobileApiServiceServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mobileapi.v1.MobileApiService/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MobileApiServiceServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MobileApiService_ServiceDesc is the grpc.ServiceDesc for MobileApiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MobileApiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mobileapi.v1.MobileApiService",
	HandlerType: (*MobileApiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserById",
			Handler:    _MobileApiService_GetUserById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mobileapi/v1/mobileapi.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.0
// source: health_info/health_info.proto

package helloworld

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

// HealthInfoServiceClient is the client API for HealthInfoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthInfoServiceClient interface {
	CheckHealth(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	GetServiceInfo(ctx context.Context, in *ServiceInfoRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error)
}

type healthInfoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthInfoServiceClient(cc grpc.ClientConnInterface) HealthInfoServiceClient {
	return &healthInfoServiceClient{cc}
}

func (c *healthInfoServiceClient) CheckHealth(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/health_info.HealthInfoService/CheckHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthInfoServiceClient) GetServiceInfo(ctx context.Context, in *ServiceInfoRequest, opts ...grpc.CallOption) (*ServiceInfoResponse, error) {
	out := new(ServiceInfoResponse)
	err := c.cc.Invoke(ctx, "/health_info.HealthInfoService/GetServiceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthInfoServiceServer is the server API for HealthInfoService service.
// All implementations must embed UnimplementedHealthInfoServiceServer
// for forward compatibility
type HealthInfoServiceServer interface {
	CheckHealth(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	GetServiceInfo(context.Context, *ServiceInfoRequest) (*ServiceInfoResponse, error)
	mustEmbedUnimplementedHealthInfoServiceServer()
}

// UnimplementedHealthInfoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHealthInfoServiceServer struct {
}

func (UnimplementedHealthInfoServiceServer) CheckHealth(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedHealthInfoServiceServer) GetServiceInfo(context.Context, *ServiceInfoRequest) (*ServiceInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServiceInfo not implemented")
}
func (UnimplementedHealthInfoServiceServer) mustEmbedUnimplementedHealthInfoServiceServer() {}

// UnsafeHealthInfoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthInfoServiceServer will
// result in compilation errors.
type UnsafeHealthInfoServiceServer interface {
	mustEmbedUnimplementedHealthInfoServiceServer()
}

func RegisterHealthInfoServiceServer(s grpc.ServiceRegistrar, srv HealthInfoServiceServer) {
	s.RegisterService(&HealthInfoService_ServiceDesc, srv)
}

func _HealthInfoService_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthInfoServiceServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_info.HealthInfoService/CheckHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthInfoServiceServer).CheckHealth(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HealthInfoService_GetServiceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthInfoServiceServer).GetServiceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/health_info.HealthInfoService/GetServiceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthInfoServiceServer).GetServiceInfo(ctx, req.(*ServiceInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthInfoService_ServiceDesc is the grpc.ServiceDesc for HealthInfoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthInfoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "health_info.HealthInfoService",
	HandlerType: (*HealthInfoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _HealthInfoService_CheckHealth_Handler,
		},
		{
			MethodName: "GetServiceInfo",
			Handler:    _HealthInfoService_GetServiceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "health_info/health_info.proto",
}

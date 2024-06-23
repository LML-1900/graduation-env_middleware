// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: env_service.proto

package grpc_env_service

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

// EnvironmentDataClient is the client API for EnvironmentData service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EnvironmentDataClient interface {
	GetStaticData(ctx context.Context, in *GetStaticDataRequest, opts ...grpc.CallOption) (EnvironmentData_GetStaticDataClient, error)
	UpdateCrater(ctx context.Context, in *Crater, opts ...grpc.CallOption) (*UpdateCraterResponse, error)
	UpdateObstacle(ctx context.Context, in *Obstacle, opts ...grpc.CallOption) (*UpdateObstacleResponse, error)
	GetRoutePoints(ctx context.Context, in *StartStopPoints, opts ...grpc.CallOption) (*RoutePoints, error)
}

type environmentDataClient struct {
	cc grpc.ClientConnInterface
}

func NewEnvironmentDataClient(cc grpc.ClientConnInterface) EnvironmentDataClient {
	return &environmentDataClient{cc}
}

func (c *environmentDataClient) GetStaticData(ctx context.Context, in *GetStaticDataRequest, opts ...grpc.CallOption) (EnvironmentData_GetStaticDataClient, error) {
	stream, err := c.cc.NewStream(ctx, &EnvironmentData_ServiceDesc.Streams[0], "/env_data_service.EnvironmentData/GetStaticData", opts...)
	if err != nil {
		return nil, err
	}
	x := &environmentDataGetStaticDataClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EnvironmentData_GetStaticDataClient interface {
	Recv() (*GetStaticDataResponse, error)
	grpc.ClientStream
}

type environmentDataGetStaticDataClient struct {
	grpc.ClientStream
}

func (x *environmentDataGetStaticDataClient) Recv() (*GetStaticDataResponse, error) {
	m := new(GetStaticDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *environmentDataClient) UpdateCrater(ctx context.Context, in *Crater, opts ...grpc.CallOption) (*UpdateCraterResponse, error) {
	out := new(UpdateCraterResponse)
	err := c.cc.Invoke(ctx, "/env_data_service.EnvironmentData/UpdateCrater", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentDataClient) UpdateObstacle(ctx context.Context, in *Obstacle, opts ...grpc.CallOption) (*UpdateObstacleResponse, error) {
	out := new(UpdateObstacleResponse)
	err := c.cc.Invoke(ctx, "/env_data_service.EnvironmentData/UpdateObstacle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentDataClient) GetRoutePoints(ctx context.Context, in *StartStopPoints, opts ...grpc.CallOption) (*RoutePoints, error) {
	out := new(RoutePoints)
	err := c.cc.Invoke(ctx, "/env_data_service.EnvironmentData/GetRoutePoints", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EnvironmentDataServer is the server API for EnvironmentData service.
// All implementations must embed UnimplementedEnvironmentDataServer
// for forward compatibility
type EnvironmentDataServer interface {
	GetStaticData(*GetStaticDataRequest, EnvironmentData_GetStaticDataServer) error
	UpdateCrater(context.Context, *Crater) (*UpdateCraterResponse, error)
	UpdateObstacle(context.Context, *Obstacle) (*UpdateObstacleResponse, error)
	GetRoutePoints(context.Context, *StartStopPoints) (*RoutePoints, error)
	mustEmbedUnimplementedEnvironmentDataServer()
}

// UnimplementedEnvironmentDataServer must be embedded to have forward compatible implementations.
type UnimplementedEnvironmentDataServer struct {
}

func (UnimplementedEnvironmentDataServer) GetStaticData(*GetStaticDataRequest, EnvironmentData_GetStaticDataServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStaticData not implemented")
}
func (UnimplementedEnvironmentDataServer) UpdateCrater(context.Context, *Crater) (*UpdateCraterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCrater not implemented")
}
func (UnimplementedEnvironmentDataServer) UpdateObstacle(context.Context, *Obstacle) (*UpdateObstacleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateObstacle not implemented")
}
func (UnimplementedEnvironmentDataServer) GetRoutePoints(context.Context, *StartStopPoints) (*RoutePoints, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRoutePoints not implemented")
}
func (UnimplementedEnvironmentDataServer) mustEmbedUnimplementedEnvironmentDataServer() {}

// UnsafeEnvironmentDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EnvironmentDataServer will
// result in compilation errors.
type UnsafeEnvironmentDataServer interface {
	mustEmbedUnimplementedEnvironmentDataServer()
}

func RegisterEnvironmentDataServer(s grpc.ServiceRegistrar, srv EnvironmentDataServer) {
	s.RegisterService(&EnvironmentData_ServiceDesc, srv)
}

func _EnvironmentData_GetStaticData_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStaticDataRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EnvironmentDataServer).GetStaticData(m, &environmentDataGetStaticDataServer{stream})
}

type EnvironmentData_GetStaticDataServer interface {
	Send(*GetStaticDataResponse) error
	grpc.ServerStream
}

type environmentDataGetStaticDataServer struct {
	grpc.ServerStream
}

func (x *environmentDataGetStaticDataServer) Send(m *GetStaticDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EnvironmentData_UpdateCrater_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crater)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentDataServer).UpdateCrater(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/env_data_service.EnvironmentData/UpdateCrater",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentDataServer).UpdateCrater(ctx, req.(*Crater))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentData_UpdateObstacle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Obstacle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentDataServer).UpdateObstacle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/env_data_service.EnvironmentData/UpdateObstacle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentDataServer).UpdateObstacle(ctx, req.(*Obstacle))
	}
	return interceptor(ctx, in, info, handler)
}

func _EnvironmentData_GetRoutePoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartStopPoints)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EnvironmentDataServer).GetRoutePoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/env_data_service.EnvironmentData/GetRoutePoints",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EnvironmentDataServer).GetRoutePoints(ctx, req.(*StartStopPoints))
	}
	return interceptor(ctx, in, info, handler)
}

// EnvironmentData_ServiceDesc is the grpc.ServiceDesc for EnvironmentData service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EnvironmentData_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "env_data_service.EnvironmentData",
	HandlerType: (*EnvironmentDataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateCrater",
			Handler:    _EnvironmentData_UpdateCrater_Handler,
		},
		{
			MethodName: "UpdateObstacle",
			Handler:    _EnvironmentData_UpdateObstacle_Handler,
		},
		{
			MethodName: "GetRoutePoints",
			Handler:    _EnvironmentData_GetRoutePoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStaticData",
			Handler:       _EnvironmentData_GetStaticData_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "env_service.proto",
}

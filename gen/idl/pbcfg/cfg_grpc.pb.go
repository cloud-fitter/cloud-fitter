// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbcfg

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

// StatisticServiceClient is the client API for StatisticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatisticServiceClient interface {
	Statistic(ctx context.Context, in *StatisticReq, opts ...grpc.CallOption) (*StatisticResp, error)
}

type statisticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatisticServiceClient(cc grpc.ClientConnInterface) StatisticServiceClient {
	return &statisticServiceClient{cc}
}

func (c *statisticServiceClient) Statistic(ctx context.Context, in *StatisticReq, opts ...grpc.CallOption) (*StatisticResp, error) {
	out := new(StatisticResp)
	err := c.cc.Invoke(ctx, "/pbcfg.StatisticService/Statistic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatisticServiceServer is the server API for StatisticService service.
// All implementations must embed UnimplementedStatisticServiceServer
// for forward compatibility
type StatisticServiceServer interface {
	Statistic(context.Context, *StatisticReq) (*StatisticResp, error)
	mustEmbedUnimplementedStatisticServiceServer()
}

// UnimplementedStatisticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStatisticServiceServer struct {
}

func (UnimplementedStatisticServiceServer) Statistic(context.Context, *StatisticReq) (*StatisticResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Statistic not implemented")
}
func (UnimplementedStatisticServiceServer) mustEmbedUnimplementedStatisticServiceServer() {}

// UnsafeStatisticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatisticServiceServer will
// result in compilation errors.
type UnsafeStatisticServiceServer interface {
	mustEmbedUnimplementedStatisticServiceServer()
}

func RegisterStatisticServiceServer(s grpc.ServiceRegistrar, srv StatisticServiceServer) {
	s.RegisterService(&StatisticService_ServiceDesc, srv)
}

func _StatisticService_Statistic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatisticReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatisticServiceServer).Statistic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbcfg.StatisticService/Statistic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatisticServiceServer).Statistic(ctx, req.(*StatisticReq))
	}
	return interceptor(ctx, in, info, handler)
}

// StatisticService_ServiceDesc is the grpc.ServiceDesc for StatisticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StatisticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbcfg.StatisticService",
	HandlerType: (*StatisticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Statistic",
			Handler:    _StatisticService_Statistic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/pbcfg/cfg.proto",
}
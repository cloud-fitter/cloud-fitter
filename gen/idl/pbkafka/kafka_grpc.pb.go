// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pbkafka

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

// KafkaServiceClient is the client API for KafkaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KafkaServiceClient interface {
	// 查询Kafka明细 - 支持云类型、区域、账户、分页等过滤条件
	ListKafkaDetail(ctx context.Context, in *ListDetailReq, opts ...grpc.CallOption) (*ListDetailResp, error)
	// 全量查询Kafka - 根据云类型
	ListKafka(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error)
}

type kafkaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewKafkaServiceClient(cc grpc.ClientConnInterface) KafkaServiceClient {
	return &kafkaServiceClient{cc}
}

func (c *kafkaServiceClient) ListKafkaDetail(ctx context.Context, in *ListDetailReq, opts ...grpc.CallOption) (*ListDetailResp, error) {
	out := new(ListDetailResp)
	err := c.cc.Invoke(ctx, "/pbkafka.KafkaService/ListKafkaDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *kafkaServiceClient) ListKafka(ctx context.Context, in *ListReq, opts ...grpc.CallOption) (*ListResp, error) {
	out := new(ListResp)
	err := c.cc.Invoke(ctx, "/pbkafka.KafkaService/ListKafka", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KafkaServiceServer is the server API for KafkaService service.
// All implementations must embed UnimplementedKafkaServiceServer
// for forward compatibility
type KafkaServiceServer interface {
	// 查询Kafka明细 - 支持云类型、区域、账户、分页等过滤条件
	ListKafkaDetail(context.Context, *ListDetailReq) (*ListDetailResp, error)
	// 全量查询Kafka - 根据云类型
	ListKafka(context.Context, *ListReq) (*ListResp, error)
	mustEmbedUnimplementedKafkaServiceServer()
}

// UnimplementedKafkaServiceServer must be embedded to have forward compatible implementations.
type UnimplementedKafkaServiceServer struct {
}

func (UnimplementedKafkaServiceServer) ListKafkaDetail(context.Context, *ListDetailReq) (*ListDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKafkaDetail not implemented")
}
func (UnimplementedKafkaServiceServer) ListKafka(context.Context, *ListReq) (*ListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKafka not implemented")
}
func (UnimplementedKafkaServiceServer) mustEmbedUnimplementedKafkaServiceServer() {}

// UnsafeKafkaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KafkaServiceServer will
// result in compilation errors.
type UnsafeKafkaServiceServer interface {
	mustEmbedUnimplementedKafkaServiceServer()
}

func RegisterKafkaServiceServer(s grpc.ServiceRegistrar, srv KafkaServiceServer) {
	s.RegisterService(&KafkaService_ServiceDesc, srv)
}

func _KafkaService_ListKafkaDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServiceServer).ListKafkaDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkafka.KafkaService/ListKafkaDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServiceServer).ListKafkaDetail(ctx, req.(*ListDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _KafkaService_ListKafka_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KafkaServiceServer).ListKafka(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pbkafka.KafkaService/ListKafka",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KafkaServiceServer).ListKafka(ctx, req.(*ListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// KafkaService_ServiceDesc is the grpc.ServiceDesc for KafkaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KafkaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbkafka.KafkaService",
	HandlerType: (*KafkaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKafkaDetail",
			Handler:    _KafkaService_ListKafkaDetail_Handler,
		},
		{
			MethodName: "ListKafka",
			Handler:    _KafkaService_ListKafka_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/pbkafka/kafka.proto",
}

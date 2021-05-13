package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/demo"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbdomain"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbkafka"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pboss"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbrds"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbstatistic"
)

type Server struct {
	// 使用unsafe可以强制让编译器检查是否实现了相关方法
	demo.UnsafeDemoServiceServer
	pbecs.UnsafeEcsServiceServer
	pbstatistic.UnsafeStatisticServiceServer
	pbrds.UnsafeRdsServiceServer
	pbdomain.UnsafeDomainServiceServer
	pboss.UnsafeOssServiceServer
	pbkafka.UnsafeKafkaServiceServer
}

func (s *Server) Echo(ctx context.Context, req *demo.StringMessage) (*demo.StringMessage, error) {
	return &demo.StringMessage{
		Value: "Hello, My Friend! Welcome to Cloud Fitter",
	}, nil
}

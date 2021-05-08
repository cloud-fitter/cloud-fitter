package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/demo"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbdomain"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pboss"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbrds"
)

type Server struct {
	// pb.go中自动生成的，是个空结构体
	demo.YourServiceServer
	pbecs.ECSServiceServer
	pbcfg.StatisticServiceServer
	pbrds.RDSServiceServer
	pbdomain.DomainServiceServer
	pboss.OssServiceServer
}

func (s *Server) Echo(ctx context.Context, req *demo.StringMessage) (*demo.StringMessage, error) {
	return &demo.StringMessage{
		Value: "Hello, My Friend! Welcome to Cloud Fitter",
	}, nil
}

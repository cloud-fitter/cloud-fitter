package server

import (
	"context"

	gw "github.com/cloud-fitter/cloud-fitter/gen/idl/demo"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
)

type Server struct {
	// pb.go中自动生成的，是个空结构体
	gw.YourServiceServer
	pbecs.ECSServiceServer
}

func (s *Server) Echo(ctx context.Context, req *gw.StringMessage) (*gw.StringMessage, error) {
	return &gw.StringMessage{
		Value: "Hello, My Friend! Welcome to Cloud Fitter",
	}, nil
}

func (s *Server) ECSStatistic(ctx context.Context, req *pbecs.ECSStatisticReq) (*pbecs.ECSStatisticResp, error) {

	return &pbecs.ECSStatisticResp{}, nil
}

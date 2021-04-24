package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/demo"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/server/ecs"
)

type Server struct {
	// pb.go中自动生成的，是个空结构体
	demo.YourServiceServer
	pbecs.ECSServiceServer
}

func (s *Server) Echo(ctx context.Context, req *demo.StringMessage) (*demo.StringMessage, error) {
	return &demo.StringMessage{
		Value: "Hello, My Friend! Welcome to Cloud Fitter",
	}, nil
}

func (s *Server) ECSStatistic(ctx context.Context, req *pbecs.ECSStatisticReq) (*pbecs.ECSStatisticResp, error) {
	ecses, err := ecs.Statistic(req.Provider)
	return &pbecs.ECSStatisticResp{Ecses: ecses}, err
}

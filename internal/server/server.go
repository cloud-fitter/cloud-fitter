package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/demo"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/server/ecs"
	"github.com/cloud-fitter/cloud-fitter/internal/server/statistic"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

type Server struct {
	// pb.go中自动生成的，是个空结构体
	demo.YourServiceServer
	pbecs.ECSServiceServer
	pbcfg.StatisticServiceServer
}

func (s *Server) Echo(ctx context.Context, req *demo.StringMessage) (*demo.StringMessage, error) {
	return &demo.StringMessage{
		Value: "Hello, My Friend! Welcome to Cloud Fitter",
	}, nil
}

func (s *Server) List(ctx context.Context, req *pbecs.ListReq) (*pbecs.ListResp, error) {
	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}
	resp, err := ecs.List(req, tenanters)
	if err != nil {
		return nil, errors.WithMessage(err, "ecs list error")
	}
	return resp, nil
}

func (s *Server) Statistic(ctx context.Context, req *pbcfg.StatisticReq) (*pbcfg.StatisticResp, error) {
	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}
	cfgs, err := statistic.Statistic(req.Provider, tenanters)
	return &pbcfg.StatisticResp{Cfgs: cfgs}, err
}

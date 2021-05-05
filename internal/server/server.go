package server

import (
	"context"
	"sync"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/demo"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/server/ecs"
	"github.com/cloud-fitter/cloud-fitter/internal/server/statistic"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
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

func (s *Server) ListDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	resp, err := ecs.ListDetail(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "ecs list detail error")
	}
	return resp, nil
}

func (s *Server) List(ctx context.Context, req *pbecs.ListReq) (*pbecs.ListResp, error) {
	resp, err := ecs.List(ctx, req)
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
	cfgs, err := statistic.Statistic(ctx, req.Provider, tenanters)
	return &pbcfg.StatisticResp{Cfgs: cfgs}, err
}

func (s *Server) StatisticAll(ctx context.Context, req *pbcfg.StatisticAllReq) (*pbcfg.StatisticResp, error) {
	var (
		wg     sync.WaitGroup
		mutex  sync.Mutex
		result []*pbcfg.StatisticRespList
	)
	wg.Add(len(pbtenant.CloudProvider_name))

	for i := range pbtenant.CloudProvider_name {
		go func(provider pbtenant.CloudProvider) {
			defer wg.Done()
			tenanters, err := tenanter.GetTenanters(provider)
			if err != nil {
				glog.Errorf("Provider %v GetTenanters error %v", provider, err)
				return
			}
			cfgs, err := statistic.Statistic(ctx, provider, tenanters)
			mutex.Lock()
			result = append(result, cfgs...)
			mutex.Unlock()
		}(pbtenant.CloudProvider(i))
	}

	wg.Wait()
	return &pbcfg.StatisticResp{Cfgs: result}, nil
}

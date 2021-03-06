package server

import (
	"context"
	"sync"

	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbstatistic"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/server/statistic"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func (s *Server) Statistic(ctx context.Context, req *pbstatistic.StatisticReq) (*pbstatistic.StatisticResp, error) {
	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		glog.Errorf("GetTenanters error %+v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	cfgs, err := statistic.Statistic(ctx, req.Provider, tenanters)
	if err != nil {
		glog.Errorf("Statistic error %+v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pbstatistic.StatisticResp{Cfgs: cfgs}, nil
}

func (s *Server) StatisticAll(ctx context.Context, req *pbstatistic.StatisticAllReq) (*pbstatistic.StatisticResp, error) {
	var (
		wg     sync.WaitGroup
		mutex  sync.Mutex
		result []*pbstatistic.StatisticInfo
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
	return &pbstatistic.StatisticResp{Cfgs: result}, nil
}

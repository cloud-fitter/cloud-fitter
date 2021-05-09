package ecs

import (
	"sync"

	"github.com/golang/glog"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/service/ecser"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"context"

	"github.com/pkg/errors"
)

func ListDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	var (
		ecs ecser.Ecser
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	region, err := tenanter.NewRegion(req.Provider, req.RegionId)
	if err != nil {
		return nil, errors.WithMessagef(err, "provider %v regionId %v", req.Provider, req.RegionId)
	}

	for _, tenanter := range tenanters {
		if req.AccountName == "" || tenanter.AccountName() == req.AccountName {
			if ecs, err = ecser.NewEcsClient(req.Provider, region, tenanter); err != nil {
				return nil, errors.WithMessage(err, "NewEcsClient error")
			}
		}
	}

	return ecs.ListDetail(ctx, req)
}

func List(ctx context.Context, req *pbecs.ListReq) (*pbecs.ListResp, error) {
	var (
		wg    sync.WaitGroup
		mutex sync.Mutex
		ecses []*pbecs.EcsInstance
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	regions := tenanter.GetAllRegionIds(req.Provider)

	wg.Add(len(tenanters) * len(regions))
	for _, t := range tenanters {
		for _, region := range regions {
			go func(tenant tenanter.Tenanter, region tenanter.Region) {
				defer wg.Done()
				ecs, err := ecser.NewEcsClient(req.Provider, region, tenant)
				if err != nil {
					glog.Errorf("New Ecs Client error %v", err)
					return
				}

				request := &pbecs.ListDetailReq{
					Provider:    req.Provider,
					AccountName: tenant.AccountName(),
					RegionId:    region.GetId(),
					PageNumber:  1,
					PageSize:    100,
					NextToken:   "",
				}
				for {
					resp, err := ecs.ListDetail(ctx, request)
					if err != nil {
						glog.Errorf("ListDetail error %v", err)
						return
					}
					mutex.Lock()
					ecses = append(ecses, resp.Ecses...)
					mutex.Unlock()
					if resp.Finished {
						break
					}
					request.PageNumber, request.PageSize, request.NextToken = resp.PageNumber, resp.PageSize, resp.NextToken
				}
			}(t, region)

		}
	}
	wg.Wait()

	return &pbecs.ListResp{Ecses: ecses}, nil
}

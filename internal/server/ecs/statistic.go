package ecs

import (
	"fmt"
	"sync"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/ecser"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

var (
	ErrStatisticNotSupported = errors.New("cloud not supported statistic ecs")
	ErrNoTenantsInfo         = errors.New("no tenants info")
)

func Statistic(provider pbtenant.CloudProvider) ([]*pbecs.ECSStatisticRespList, error) {
	tenanters := tenanter.GetTenantersMap(provider)

	if len(tenanters) == 0 {
		return nil, errors.WithMessage(ErrNoTenantsInfo, fmt.Sprintf("cloud is %d", provider))
	}

	var (
		results []*pbecs.ECSStatisticRespList
		mutex   sync.Mutex
	)

	// TODO: 这一块应该继续抽象
	switch provider {
	case pbtenant.CloudProvider_ali_cloud:
		for i := range tenanters {
			var wg sync.WaitGroup
			wg.Add(len(pbtenant.AliRegionId_name))
			for rId := range pbtenant.AliRegionId_name {
				go func(rId int32, t tenanter.Tenanter) {
					defer wg.Done()
					if rId != int32(pbtenant.AliRegionId_ali_all) {
						if ali, err := ecser.NewAliEcsClient(rId, t); err == nil {
							if res, err := ali.ECSStatistic(); err == nil {
								mutex.Lock()
								results = append(results, res)
								mutex.Unlock()
							}
						}
					}
				}(rId, tenanters[i].Clone())
			}
			wg.Wait()
		}
	case pbtenant.CloudProvider_tencent_cloud:
		for i := range tenanters {
			var wg sync.WaitGroup
			wg.Add(len(pbtenant.TencentRegionId_name))
			for rId := range pbtenant.TencentRegionId_name {
				go func(rId int32, t tenanter.Tenanter) {
					defer wg.Done()
					if rId != int32(pbtenant.TencentRegionId_tc_all) {
						if tc, err := ecser.NewTencentCvmClient(rId, t); err == nil {
							if res, err := tc.ECSStatistic(); err == nil {
								mutex.Lock()
								results = append(results, res)
								mutex.Unlock()
							}
						}
					}
				}(rId, tenanters[i].Clone())
			}
			wg.Wait()
		}
	case pbtenant.CloudProvider_aws_cloud:
		for i := range tenanters {
			var wg sync.WaitGroup
			wg.Add(len(pbtenant.AwsRegionId_name))
			for rId := range pbtenant.AwsRegionId_name {
				go func(rId int32, t tenanter.Tenanter) {
					defer wg.Done()
					if rId != int32(pbtenant.AwsRegionId_aws_all) {
						if aws, err := ecser.NewAwsEcsClient(rId, t); err == nil {
							if res, err := aws.ECSStatistic(); err == nil {
								mutex.Lock()
								results = append(results, res)
								mutex.Unlock()
							}
						}
					}
				}(rId, tenanters[i].Clone())
			}
			wg.Wait()
		}
	default:
		return nil, errors.WithMessage(ErrStatisticNotSupported, fmt.Sprintf("cloud is %d", provider))
	}

	return results, nil
}

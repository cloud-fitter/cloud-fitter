package statistic

import (
	"fmt"
	"sync"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/configger"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

var (
	ErrStatisticNotSupported = errors.New("cloud not supported statistic ecs")
)

func Statistic(provider pbtenant.CloudProvider, tenanters []tenanter.Tenanter) ([]*pbcfg.StatisticRespList, error) {
	var (
		results []*pbcfg.StatisticRespList
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
						if ali, err := configger.NewAliConfigClient(rId, t); err == nil {
							if res, err := ali.Statistic(); err == nil && res.Count > 0 {
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
						if tc, err := configger.NewTencentCfgClient(rId, t); err == nil {
							if res, err := tc.Statistic(); err == nil && res.Count > 0 {
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

	case pbtenant.CloudProvider_huawei_cloud:
		for i := range tenanters {
			var wg sync.WaitGroup
			wg.Add(len(pbtenant.HuaweiRegionId_name))
			for rId := range pbtenant.HuaweiRegionId_name {
				go func(rId int32, t tenanter.Tenanter) {
					defer wg.Done()
					if rId != int32(pbtenant.HuaweiRegionId_hw_all) {
						if hw, err := configger.NewHuaweiCfgClient(rId, t); err == nil {
							if res, err := hw.Statistic(); err == nil && res.Count > 0 {
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
						if aws, err := configger.NewAwsCfgClient(rId, t); err == nil {
							if res, err := aws.Statistic(); err == nil && res.Count > 0 {
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

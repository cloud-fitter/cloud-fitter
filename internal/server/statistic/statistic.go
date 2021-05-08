package statistic

import (
	"context"
	"sync"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/service/configger"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func Statistic(ctx context.Context, provider pbtenant.CloudProvider, tenanters []tenanter.Tenanter) ([]*pbcfg.StatisticRespList, error) {
	var (
		results []*pbcfg.StatisticRespList
		mutex   sync.Mutex
	)

	regions := tenanter.GetAllRegionIds(provider)

	for i := range tenanters {
		var wg sync.WaitGroup
		wg.Add(len(regions))
		for _, region := range regions {
			go func(region tenanter.Region, t tenanter.Tenanter) {
				defer wg.Done()
				if cfg, err := configger.NewConfigClient(provider, region, t); err == nil {
					if res, err := cfg.Statistic(ctx); err == nil && res.Count > 0 {
						mutex.Lock()
						results = append(results, res)
						mutex.Unlock()
					}
				}
			}(region, tenanters[i].Clone())
		}
		wg.Wait()
	}

	return results, nil
}

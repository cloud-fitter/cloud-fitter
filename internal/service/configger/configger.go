package configger

import (
	"context"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbstatistic"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrStatisticNotSupported = errors.New("cloud not supported statistic ecs")
	ErrCfgPanic              = errors.New("configger init panic")
)

type Configger interface {
	// 统计接口
	Statistic(ctx context.Context) (*pbstatistic.StatisticInfo, error)
}

func NewConfigClient(provider pbtenant.CloudProvider, region tenanter.Region, tenant tenanter.Tenanter) (cfg Configger, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewConfigClient panic %v", err1)
			err = errors.WithMessagef(ErrCfgPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_ali:
		return newAliCfgClient(region, tenant)
	case pbtenant.CloudProvider_tencent:
		return newTencentCfgClient(region, tenant)
	case pbtenant.CloudProvider_huawei:
		return newHuaweiCfgClient(region, tenant)
	case pbtenant.CloudProvider_aws:
		return newAwsCfgClient(region, tenant)
	}

	err = errors.WithMessagef(ErrStatisticNotSupported, "cloud provider %v", provider)
	return
}

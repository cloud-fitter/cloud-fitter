package ecser

import (
	"context"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrEcsListNotSupported = errors.New("cloud not supported ecs list")
	ErrEcserPanic          = errors.New("ecs init panic")
)

type Ecser interface {
	ListDetail(ctx context.Context, req *pbecs.ListDetailReq) (resp *pbecs.ListDetailResp, err error)
}

func NewEcsClient(provider pbtenant.CloudProvider, region tenanter.Region, tenant tenanter.Tenanter) (ecser Ecser, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewEcsClient panic %v", err1)
			err = errors.WithMessagef(ErrEcserPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_ali_cloud:
		return NewAliEcsClient(region, tenant)
	case pbtenant.CloudProvider_tencent_cloud:
		return NewTencentCvmClient(region, tenant)
	case pbtenant.CloudProvider_huawei_cloud:
		return NewHuaweiEcsClient(region, tenant)
	case pbtenant.CloudProvider_aws_cloud:
		return NewAwsEcsClient(region, tenant)
	}

	err = errors.WithMessagef(ErrEcsListNotSupported, "cloud provider %v region %v", provider, region)
	return
}

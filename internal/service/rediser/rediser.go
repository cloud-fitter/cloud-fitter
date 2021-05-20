package rediser

import (
	"context"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbredis"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrRedisListNotSupported = errors.New("cloud not supported redis list")
	ErrRediserPanic          = errors.New("redis init panic")
)

type Rediser interface {
	ListDetail(ctx context.Context, req *pbredis.ListDetailReq) (resp *pbredis.ListDetailResp, err error)
}

func NewRedisClient(provider pbtenant.CloudProvider, region tenanter.Region, tenant tenanter.Tenanter) (redis Rediser, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewRedisClient panic %v", err1)
			err = errors.WithMessagef(ErrRediserPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_tencent:
		return newTencentRedisClient(region, tenant)
	}

	err = errors.WithMessagef(ErrRedisListNotSupported, "cloud provider %v region %v", provider, region)
	return
}

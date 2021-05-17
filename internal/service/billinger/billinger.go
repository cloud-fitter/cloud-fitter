package billinger

import (
	"context"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbbilling"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrBillingListNotSupported = errors.New("cloud not supported billing list")
	ErrBillingerPanic          = errors.New("billing init panic")
)

type Billinger interface {
	ListDetail(ctx context.Context, req *pbbilling.ListDetailReq) (resp *pbbilling.ListDetailResp, err error)
}

func NewBillingClient(provider pbtenant.CloudProvider, tenant tenanter.Tenanter) (billing Billinger, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewBillingClient panic %v", err1)
			err = errors.WithMessagef(ErrBillingerPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_ali:
		return newAliBillingClient(tenant)
	case pbtenant.CloudProvider_tencent:
		return newTencentBillingClient(tenant)
	}

	err = errors.WithMessagef(ErrBillingListNotSupported, "cloud provider %v ", provider)
	return
}

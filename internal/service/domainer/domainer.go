package domainer

import (
	"context"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbdomain"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrDomainListNotSupported = errors.New("cloud not supported domain list")
	ErrDomainerPanic          = errors.New("domain init panic")
)

type Rdser interface {
	ListDetail(ctx context.Context, req *pbdomain.ListDetailReq) (resp *pbdomain.ListDetailResp, err error)
}

func NewDomainClient(provider pbtenant.CloudProvider, tenant tenanter.Tenanter) (rds Rdser, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewDomainClient panic %v", err1)
			err = errors.WithMessagef(ErrDomainerPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_ali:
		return newAliDomainClient(tenant)
	case pbtenant.CloudProvider_tencent:
		return newTencentDomainClient(tenant)
	}

	err = errors.WithMessagef(ErrDomainListNotSupported, "cloud provider %v ", provider)
	return
}

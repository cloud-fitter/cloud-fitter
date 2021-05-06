package rdser

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbrds"
	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrRdsListNotSupported = errors.New("cloud not supported rds list")
	ErrRdserPanic          = errors.New("rds init panic")
)

type Rdser interface {
	ListDetail(ctx context.Context, req *pbrds.ListDetailReq) (resp *pbrds.ListDetailResp, err error)
}

func NewRdsClient(provider pbtenant.CloudProvider, region tenanter.Region, tenant tenanter.Tenanter) (rds Rdser, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewRdsClient panic %v", err1)
			err = errors.WithMessagef(ErrRdserPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_ali_cloud:
		return NewAliRdsClient(region, tenant)
	}

	err = errors.WithMessagef(ErrRdsListNotSupported, "cloud provider %v region %v", provider, region)
	return
}

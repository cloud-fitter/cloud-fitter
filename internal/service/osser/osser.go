package osser

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pboss"
	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrOssListNotSupported = errors.New("cloud not supported oss list")
	ErrOsserPanic          = errors.New("oss init panic")
)

type Osser interface {
	ListDetail(ctx context.Context, req *pboss.ListDetailReq) (resp *pboss.ListDetailResp, err error)
}

func NewOssClient(provider pbtenant.CloudProvider, tenant tenanter.Tenanter) (oss Osser, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewOssClient panic %v", err1)
			err = errors.WithMessagef(ErrOsserPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_ali:
		return newAliOssClient(tenant)
	}

	err = errors.WithMessagef(ErrOssListNotSupported, "cloud provider %v ", provider)
	return
}

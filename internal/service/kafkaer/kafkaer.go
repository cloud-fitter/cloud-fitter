package kafkaer

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbkafka"
	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	ErrKafkaListNotSupported = errors.New("cloud not supported kafka list")
	ErrKafkaerPanic          = errors.New("kafka init panic")
)

type Kafkaer interface {
	ListDetail(ctx context.Context, req *pbkafka.ListDetailReq) (resp *pbkafka.ListDetailResp, err error)
}

func NewKafkaClient(provider pbtenant.CloudProvider, region tenanter.Region, tenant tenanter.Tenanter) (kafkaer Kafkaer, err error) {
	// 部分sdk会在内部panic
	defer func() {
		if err1 := recover(); err1 != nil {
			glog.Errorf("NewKafkaClient panic %v", err1)
			err = errors.WithMessagef(ErrKafkaerPanic, "%v", err1)
		}
	}()

	switch provider {
	case pbtenant.CloudProvider_ali:
		return NewAliKafkaerClient(region, tenant)
	case pbtenant.CloudProvider_tencent:
		return NewTencentKafkaerClient(region, tenant)
	}

	err = errors.WithMessagef(ErrKafkaListNotSupported, "cloud provider %v region %v", provider, region)
	return
}

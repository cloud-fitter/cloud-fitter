package ecs

import (
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/ecser"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

var (
	ErrEcsListNotSupported = errors.New("cloud not supported ecs list")
)

func List(req *pbecs.ListReq, tenanters []tenanter.Tenanter) (*pbecs.ListResp, error) {
	var (
		ecs ecser.Ecser
		err error
	)

	switch req.Provider {
	case pbtenant.CloudProvider_ali_cloud:
		for _, tenanter := range tenanters {
			if req.AccountName == "" || tenanter.AccountName() == req.AccountName {
				if ecs, err = ecser.NewAliEcsClient(req.RegionId, tenanter); err != nil {
					return nil, errors.WithMessage(err, "NewAliEcsClient error")
				}
			}
		}

	case pbtenant.CloudProvider_tencent_cloud:
		for _, tenanter := range tenanters {
			if req.AccountName == "" || tenanter.AccountName() == req.AccountName {
				if ecs, err = ecser.NewTencentCvmClient(req.RegionId, tenanter); err != nil {
					return nil, errors.WithMessage(err, "NewTencentCvmClient error")
				}
			}
		}

	case pbtenant.CloudProvider_huawei_cloud:
		for _, tenanter := range tenanters {
			if req.AccountName == "" || tenanter.AccountName() == req.AccountName {
				if ecs, err = ecser.NewHuaweiEcsClient(req.RegionId, tenanter); err != nil {
					return nil, errors.WithMessage(err, "NewHuaweiEcsClient error")
				}
			}
		}

	case pbtenant.CloudProvider_aws_cloud:
		for _, tenanter := range tenanters {
			if req.AccountName == "" || tenanter.AccountName() == req.AccountName {
				if ecs, err = ecser.NewAwsEcsClient(req.RegionId, tenanter); err != nil {
					return nil, errors.WithMessage(err, "NewAwsEcsClient error")
				}
			}
		}

	default:
		return nil, errors.WithMessagef(ErrEcsListNotSupported, "cloud provider %v", req.Provider)
	}

	return ecs.DescribeInstances(req.PageNumber, req.PageSize)
}

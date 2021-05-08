package configger

import (
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"context"

	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TencentCfg struct {
	cli        *cvm.Client
	regionId   pbtenant.TencentRegionId
	regionName string
	tenanter.Tenanter
}

func NewTencentCfgClient(region tenanter.Region, tenant tenanter.Tenanter) (Configger, error) {
	var client *cvm.Client
	var err error

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = cvm.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), region.GetName(), profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent cfg client error")
	}
	return &TencentCfg{
		cli:        client,
		regionId:   pbtenant.TencentRegionId(region.GetId()),
		regionName: region.GetName(),
		Tenanter:   tenant,
	}, nil
}

func (cfg *TencentCfg) Statistic(ctx context.Context) (*pbcfg.StatisticRespList, error) {
	req := cvm.NewDescribeInstancesRequest()
	req.Offset = common.Int64Ptr(1)
	req.Limit = common.Int64Ptr(1)
	resp, err := cfg.cli.DescribeInstances(req)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent ListDetail error")
	}

	return &pbcfg.StatisticRespList{
		Provider:    pbtenant.CloudProvider_tencent_cloud,
		AccountName: cfg.AccountName(),
		Product:     pbtenant.CloudProduct_product_ecs,
		RegionId:    int32(cfg.regionId),
		RegionName:  cfg.regionName,
		Count:       *(resp.Response.TotalCount),
	}, nil
}

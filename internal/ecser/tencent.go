package ecser

import (
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TencentCvm struct {
	cli        *cvm.Client
	regionId   pbtenant.TencentRegionId
	regionName string
}

func NewTencentCvmClient(regionId int32, tenant tenanter.Tenanter) (Ecser, error) {
	var client *cvm.Client

	rName, err := tenanter.GetTencentRegionName(regionId)
	if err != nil {
		return nil, err
	}

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = cvm.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), rName, profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent cvm client error")
	}
	return &TencentCvm{
		cli:        client,
		regionId:   pbtenant.TencentRegionId(regionId),
		regionName: rName,
	}, nil
}

func (ecs *TencentCvm) ECSStatistic() (*pbecs.ECSStatisticRespList, error) {
	req := cvm.NewDescribeInstancesRequest()
	req.Offset = common.Int64Ptr(1)
	req.Limit = common.Int64Ptr(1)
	resp, err := ecs.cli.DescribeInstances(req)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun DescribeInstances error")
	}

	return &pbecs.ECSStatisticRespList{
		Provider:   pbtenant.CloudProvider_tencent_cloud,
		RegionId:   int32(ecs.regionId),
		Count:      *(resp.Response.TotalCount),
		RegionName: ecs.regionName,
	}, nil
}

func (ecs *TencentCvm) DescribeInstances(pageNumber, pageSize int) ([]*pbecs.ECSInstance, error) {
	req := cvm.NewDescribeInstancesRequest()
	req.Offset = common.Int64Ptr(int64((pageNumber - 1) * pageSize))
	req.Limit = common.Int64Ptr(int64(pageSize))
	resp, err := ecs.cli.DescribeInstances(req)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent DescribeInstances error")
	}

	var ecses = make([]*pbecs.ECSInstance, len(resp.Response.InstanceSet))
	for k, v := range resp.Response.InstanceSet {
		ecses[k] = &pbecs.ECSInstance{
			InstanceId: *v.InstanceId,
			PublicIps:  make([]string, len(v.PublicIpAddresses)),
		}
		for k1, v1 := range v.PublicIpAddresses {
			ecses[k].PublicIps[k1] = *v1
		}
	}
	return ecses, nil
}

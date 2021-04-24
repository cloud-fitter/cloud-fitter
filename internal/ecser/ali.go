package ecser

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	aliecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

type AliEcs struct {
	cli        *aliecs.Client
	regionId   pbtenant.AliRegionId
	regionName string
}

func NewAliEcsClient(regionId int32, tenant tenanter.Tenanter) (Ecser, error) {
	var client *aliecs.Client

	rName, err := tenanter.GetAliRegionName(regionId)
	if err != nil {
		return nil, err
	}

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有并发问题，go test 加上-race 能检测出来
		client, err = aliecs.NewClientWithAccessKey(rName, t.GetId(), t.GetSecret())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali ecs client error")
	}

	return &AliEcs{
		cli:        client,
		regionId:   pbtenant.AliRegionId(regionId),
		regionName: rName,
	}, nil
}

func (ecs *AliEcs) ECSStatistic() (*pbecs.ECSStatisticRespList, error) {
	req := aliecs.CreateDescribeInstancesRequest()
	req.PageNumber = requests.NewInteger(1)
	req.PageSize = requests.NewInteger(1)
	resp, err := ecs.cli.DescribeInstances(req)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun DescribeInstances error")
	}

	return &pbecs.ECSStatisticRespList{
		Provider:   pbtenant.CloudProvider_ali_cloud,
		RegionId:   int32(ecs.regionId),
		Count:      int64(resp.TotalCount),
		RegionName: ecs.regionName,
	}, nil
}

func (ecs *AliEcs) DescribeInstances(pageNumber, pageSize int) ([]*pbecs.ECSInstance, error) {
	req := aliecs.CreateDescribeInstancesRequest()
	req.PageNumber = requests.NewInteger(pageNumber)
	req.PageSize = requests.NewInteger(pageSize)
	resp, err := ecs.cli.DescribeInstances(req)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun DescribeInstances error")
	}

	var ecses = make([]*pbecs.ECSInstance, len(resp.Instances.Instance))
	for k, v := range resp.Instances.Instance {
		ecses[k] = &pbecs.ECSInstance{
			InstanceName: v.HostName,
			PublicIps:    v.PublicIpAddress.IpAddress,
		}
	}
	return ecses, nil
}

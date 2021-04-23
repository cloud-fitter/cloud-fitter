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
	cli *aliecs.Client
}

func NewAliEcsClient(regionId pbtenant.AliRegionId, tenant tenanter.Tenanter) (Ecser, error) {
	var (
		client *aliecs.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = aliecs.NewClientWithAccessKey(tenanter.GetAliRegionId(regionId), t.GetId(), t.GetSecret())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali ecs client error")
	}
	return &AliEcs{cli: client}, nil
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
			HostName:  v.HostName,
			PublicIps: v.PublicIpAddress.IpAddress,
		}
	}
	return ecses, nil
}

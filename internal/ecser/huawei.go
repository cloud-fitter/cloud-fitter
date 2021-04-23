package ecser

import (
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	hwecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	"github.com/pkg/errors"
)

type HuaweiEcs struct {
	cli *hwecs.EcsClient
}

func NewHuaweiEcsClient(regionId int32, tenant tenanter.Tenanter) (Ecser, error) {
	var client *hwecs.EcsClient

	rName, err := tenanter.GetHuaweiRegionName(regionId)
	if err != nil {
		return nil, err
	}

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		auth := basic.NewCredentialsBuilder().WithAk(t.GetId()).WithSk(t.GetSecret()).Build()
		hcClient := hwecs.EcsClientBuilder().WithRegion(region.ValueOf(rName)).WithCredential(auth).Build()
		client = hwecs.NewEcsClient(hcClient)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init huawei ecs client error")
	}
	return &HuaweiEcs{cli: client}, nil
}

func (ecs *HuaweiEcs) ECSStatistic() (*pbecs.ECSStatisticRespList, error) {
	return nil, nil
}

func (ecs *HuaweiEcs) DescribeInstances(pageNumber, pageSize int) ([]*pbecs.ECSInstance, error) {
	req := new(model.ListServersDetailsRequest)
	offset := int32((pageNumber - 1) * pageSize)
	req.Offset = &offset
	limit := int32(pageSize)
	req.Limit = &limit

	resp, err := ecs.cli.ListServersDetails(req)
	if err != nil {
		return nil, errors.Wrap(err, "Huawei DescribeInstances error")
	}

	servers := *resp.Servers
	var ecses = make([]*pbecs.ECSInstance, len(servers))
	for k, v := range servers {
		ecses[k] = &pbecs.ECSInstance{
			InstanceName: v.Name,
			PublicIps:    []string{v.AccessIPv4},
		}
	}
	return ecses, nil
}

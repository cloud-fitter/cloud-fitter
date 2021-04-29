package ecser

import (
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	aliecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

var aliClientMutex sync.Mutex

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
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		client, err = aliecs.NewClientWithAccessKey(rName, t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
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

func (ecs *AliEcs) DescribeInstances(pageNumber, pageSize int32) (*pbecs.ListResp, error) {
	req := aliecs.CreateDescribeInstancesRequest()
	req.PageNumber = requests.NewInteger(int(pageNumber))
	req.PageSize = requests.NewInteger(int(pageSize))
	resp, err := ecs.cli.DescribeInstances(req)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun DescribeInstances error")
	}

	var ecses = make([]*pbecs.ECSInstance, len(resp.Instances.Instance))
	for k, v := range resp.Instances.Instance {
		ecses[k] = &pbecs.ECSInstance{
			InstanceId:   v.InstanceId,
			InstanceName: v.InstanceName,
			RegionName:   ecs.regionName,
			InstanceType: v.InstanceType,
			PublicIps:    v.PublicIpAddress.IpAddress,
			Cpu:          int32(v.CPU),
			Memory:       int32(v.Memory),
			Description:  v.Description,
			Status:       v.Status,
			CreationTime: v.CreationTime,
			ExpireTime:   v.ExpiredTime,
		}
	}

	return &pbecs.ListResp{
		Ecses:      ecses,
		NextToken:  resp.NextToken,
		PageNumber: int32(resp.PageNumber),
		PageSize:   int32(resp.PageSize),
		RequestId:  resp.RequestId,
	}, nil
}

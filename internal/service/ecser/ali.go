package ecser

import (
	"context"
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	aliecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var aliClientMutex sync.Mutex

type AliEcs struct {
	cli      *aliecs.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newAliEcsClient(region tenanter.Region, tenant tenanter.Tenanter) (Ecser, error) {
	var (
		client *aliecs.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		client, err = aliecs.NewClientWithAccessKey(region.GetName(), t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali ecs client error")
	}

	return &AliEcs{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (ecs *AliEcs) ListDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	request := aliecs.CreateDescribeInstancesRequest()
	request.PageNumber = requests.NewInteger(int(req.PageNumber))
	request.PageSize = requests.NewInteger(int(req.PageSize))
	request.NextToken = req.NextToken
	resp, err := ecs.cli.DescribeInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun ListDetail error")
	}

	var ecses = make([]*pbecs.EcsInstance, len(resp.Instances.Instance))
	for k, v := range resp.Instances.Instance {
		ecses[k] = &pbecs.EcsInstance{
			Provider:     pbtenant.CloudProvider_ali,
			AccountName:  ecs.tenanter.AccountName(),
			InstanceId:   v.InstanceId,
			InstanceName: v.InstanceName,
			RegionName:   ecs.region.GetName(),
			PublicIps:    v.PublicIpAddress.IpAddress,
			InstanceType: v.InstanceType,
			Cpu:          int32(v.CPU),
			Memory:       int32(v.Memory),
			Description:  v.Description,
			Status:       v.Status,
			CreationTime: v.CreationTime,
			ExpireTime:   v.ExpiredTime,
		}
	}

	isFinished := false
	if len(ecses) < int(req.PageSize) {
		isFinished = true
	}

	return &pbecs.ListDetailResp{
		Ecses:      ecses,
		Finished:   isFinished,
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		NextToken:  resp.NextToken,
		RequestId:  resp.RequestId,
	}, nil
}

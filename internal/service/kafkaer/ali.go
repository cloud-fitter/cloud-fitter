package kafkaer

import (
	"context"
	"sync"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alikafka"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbkafka"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var (
	aliClientMutex sync.Mutex
	aliKafkaStatus = map[int]string{
		0:  "待部署",
		1:  "部署中",
		5:  "服务中",
		15: "已过期",
	}
)

type AliKafka struct {
	cli      *alikafka.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newAliKafkaerClient(region tenanter.Region, tenant tenanter.Tenanter) (Kafkaer, error) {
	var (
		client *alikafka.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		client, err = alikafka.NewClientWithAccessKey(region.GetName(), t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali rds client error")
	}

	return &AliKafka{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (kafka *AliKafka) ListDetail(ctx context.Context, req *pbkafka.ListDetailReq) (*pbkafka.ListDetailResp, error) {
	request := alikafka.CreateGetInstanceListRequest()
	resp, err := kafka.cli.GetInstanceList(request)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun ListDetail error")
	}

	var kafkas = make([]*pbkafka.KafkaInstance, len(resp.InstanceList.InstanceVO))
	for k, v := range resp.InstanceList.InstanceVO {
		status, _ := aliKafkaStatus[v.ServiceStatus]

		kafkas[k] = &pbkafka.KafkaInstance{
			Provider:      pbtenant.CloudProvider_ali,
			AccoutName:    kafka.tenanter.AccountName(),
			InstanceId:    v.InstanceId,
			InstanceName:  v.Name,
			RegionName:    kafka.region.GetName(),
			EndPoint:      v.EndPoint,
			TopicNumLimit: int32(v.TopicNumLimit),
			DistSize:      int32(v.DiskSize),
			Status:        status,
			CreateTime:    time.Unix(v.CreateTime/1000, 0).Format(time.RFC3339),
			ExpiredTime:   time.Unix(v.ExpiredTime/1000, 0).Format(time.RFC3339),
		}
	}

	isFinished := false
	if len(kafkas) < int(req.PageSize) {
		isFinished = true
	}

	return &pbkafka.ListDetailResp{
		Kafkas:     kafkas,
		Finished:   isFinished,
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		NextToken:  "",
		RequestId:  resp.RequestId,
	}, nil
}

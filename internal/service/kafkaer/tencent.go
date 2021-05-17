package kafkaer

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbkafka"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"github.com/pkg/errors"
	ckafka "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ckafka/v20190819"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

var tencentKafkaStatus = map[int64]string{
	0:  "创建中",
	1:  "运行中",
	2:  "删除中",
	5:  "隔离中",
	-1: "创建失败",
}

type TencentCkafka struct {
	cli      *ckafka.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newTencentKafkaerClient(region tenanter.Region, tenant tenanter.Tenanter) (Kafkaer, error) {
	var (
		client *ckafka.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = ckafka.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), region.GetName(), profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent ckafka client error")
	}
	return &TencentCkafka{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (kafka *TencentCkafka) ListDetail(ctx context.Context, req *pbkafka.ListDetailReq) (*pbkafka.ListDetailResp, error) {
	request := ckafka.NewDescribeInstancesRequest()
	request.Offset = common.Int64Ptr(int64((req.PageNumber - 1) * req.PageSize))
	request.Limit = common.Int64Ptr(int64(req.PageSize))
	resp, err := kafka.cli.DescribeInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent ListDetail error")
	}

	var kafkas = make([]*pbkafka.KafkaInstance, len(resp.Response.Result.InstanceList))
	for k, v := range resp.Response.Result.InstanceList {
		status := tencentKafkaStatus[*v.Status]
		kafkas[k] = &pbkafka.KafkaInstance{
			Provider:      pbtenant.CloudProvider_tencent,
			AccoutName:    kafka.tenanter.AccountName(),
			InstanceId:    *v.InstanceId,
			InstanceName:  *v.InstanceName,
			RegionName:    kafka.region.GetName(),
			EndPoint:      "",
			TopicNumLimit: 0,
			DistSize:      0,
			Status:        status,
			CreateTime:    "",
			ExpiredTime:   "",
		}
	}

	isFinished := false
	if len(kafkas) < int(req.PageSize) {
		isFinished = true
	}

	return &pbkafka.ListDetailResp{
		Kafkas:     kafkas,
		Finished:   isFinished,
		NextToken:  "",
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		RequestId:  *resp.Response.RequestId,
	}, nil
}

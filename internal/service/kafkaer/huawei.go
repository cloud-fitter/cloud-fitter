package kafkaer

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	hwregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dms/v2/region"
	hwiam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iammodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	iamregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	hwkafka "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kafka/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/kafka/v2/model"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbkafka"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

type HuaweiKafka struct {
	cli      *hwkafka.KafkaClient
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newHuaweiKafkaClient(region tenanter.Region, tenant tenanter.Tenanter) (Kafkaer, error) {
	var (
		client *hwkafka.KafkaClient
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		auth := basic.NewCredentialsBuilder().WithAk(t.GetId()).WithSk(t.GetSecret()).Build()
		rName := region.GetName()
		cli := hwiam.IamClientBuilder().WithRegion(iamregion.ValueOf(rName)).WithCredential(auth).Build()
		c := hwiam.NewIamClient(cli)
		request := new(iammodel.KeystoneListProjectsRequest)
		request.Name = &rName
		r, err := c.KeystoneListProjects(request)
		if err != nil || len(*r.Projects) == 0 {
			return nil, errors.Wrapf(err, "Huawei KeystoneListProjects regionName %s", rName)
		}
		projectId := (*r.Projects)[0].Id

		auth = basic.NewCredentialsBuilder().WithAk(t.GetId()).WithSk(t.GetSecret()).WithProjectId(projectId).Build()
		hcClient := hwkafka.KafkaClientBuilder().WithRegion(hwregion.ValueOf(rName)).WithCredential(auth).Build()
		client = hwkafka.NewKafkaClient(hcClient)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init huawei kafka client error")
	}
	return &HuaweiKafka{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (kafka *HuaweiKafka) ListDetail(ctx context.Context, req *pbkafka.ListDetailReq) (*pbkafka.ListDetailResp, error) {
	request := new(model.ListInstancesRequest)
	request.Engine = model.GetListInstancesRequestEngineEnum().KAFKA
	// offset := (req.PageNumber - 1) * req.PageSize
	// request.Offset = &offset
	// limit := req.PageSize
	// request.Limit = &limit

	resp, err := kafka.cli.ListInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Huawei ListDetail error")
	}

	instances := *resp.Instances
	var kafkas = make([]*pbkafka.KafkaInstance, len(instances))
	for k, v := range instances {
		kafkas[k] = &pbkafka.KafkaInstance{
			Provider:      pbtenant.CloudProvider_huawei,
			AccoutName:    kafka.tenanter.AccountName(),
			InstanceId:    *v.InstanceId,
			InstanceName:  *v.Name,
			RegionName:    kafka.region.GetName(),
			EndPoint:      "",
			TopicNumLimit: 0,
			DistSize:      0,
			Status:        *v.Status,
			CreateTime:    *v.CreatedAt,
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
		RequestId:  "",
	}, nil
}

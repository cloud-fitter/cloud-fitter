package rediser

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	hwdcs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/model"
	hwregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/dcs/v2/region"
	hwiam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iammodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	iamregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbredis"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

type HuaweiDcs struct {
	cli      *hwdcs.DcsClient
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newHuaweiDcsClient(region tenanter.Region, tenant tenanter.Tenanter) (Rediser, error) {
	var (
		client *hwdcs.DcsClient
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
		hcClient := hwdcs.DcsClientBuilder().WithRegion(hwregion.ValueOf(rName)).WithCredential(auth).Build()
		client = hwdcs.NewDcsClient(hcClient)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init huawei redis client error")
	}
	return &HuaweiDcs{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (redis *HuaweiDcs) ListDetail(ctx context.Context, req *pbredis.ListDetailReq) (*pbredis.ListDetailResp, error) {
	request := new(model.ListInstancesRequest)
	offset := (req.PageNumber - 1) * req.PageSize
	request.Offset = &offset
	limit := req.PageSize
	request.Limit = &limit

	resp, err := redis.cli.ListInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Huawei ListDetail error")
	}

	instances := *resp.Instances
	var redises = make([]*pbredis.RedisInstance, len(instances))
	for k, v := range instances {
		redises[k] = &pbredis.RedisInstance{
			Provider:     pbtenant.CloudProvider_huawei,
			AccoutName:   redis.tenanter.AccountName(),
			InstanceId:   *v.InstanceId,
			InstanceName: *v.Name,
			RegionName:   redis.region.GetName(),
			Size:         *v.MaxMemory,
			Status:       *v.Status,
			CreationTime: *v.CreatedAt,
			ExpireTime:   "",
		}
	}

	isFinished := false
	if len(redises) < int(req.PageSize) {
		isFinished = true
	}

	return &pbredis.ListDetailResp{
		Redises:    redises,
		Finished:   isFinished,
		NextToken:  "",
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		RequestId:  "",
	}, nil
}

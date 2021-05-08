package ecser

import (
	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	hwecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	hwregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	hwiam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iammodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	iamregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

type HuaweiEcs struct {
	cli      *hwecs.EcsClient
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func NewHuaweiEcsClient(region tenanter.Region, tenant tenanter.Tenanter) (Ecser, error) {
	var (
		client *hwecs.EcsClient
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
		hcClient := hwecs.EcsClientBuilder().WithRegion(hwregion.ValueOf(rName)).WithCredential(auth).Build()
		client = hwecs.NewEcsClient(hcClient)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init huawei ecs client error")
	}
	return &HuaweiEcs{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (ecs *HuaweiEcs) ListDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	request := new(model.ListServersDetailsRequest)
	offset := (req.PageNumber - 1) * req.PageSize
	request.Offset = &offset
	limit := req.PageSize
	request.Limit = &limit

	resp, err := ecs.cli.ListServersDetails(request)
	if err != nil {
		return nil, errors.Wrap(err, "Huawei ListDetail error")
	}

	servers := *resp.Servers
	var ecses = make([]*pbecs.ECSInstance, len(servers))
	for k, v := range servers {
		ecses[k] = &pbecs.ECSInstance{
			Provider:     pbtenant.CloudProvider_huawei_cloud,
			AccountName:  ecs.tenanter.AccountName(),
			InstanceId:   v.Id,
			InstanceName: v.Name,
			RegionName:   ecs.region.GetName(),
			InstanceType: v.Flavor.Name,
			PublicIps:    []string{v.AccessIPv4},
			// Cpu:          v.Flavor.Vcpus,
			// Memory:       v.Flavor.Ram,
			Description:  *v.Description,
			Status:       v.Status,
			CreationTime: v.Created,
			ExpireTime:   v.OSSRVUSGterminatedAt,
		}
	}

	isFinished := false
	if len(ecses) < int(req.PageSize) {
		isFinished = true
	}

	return &pbecs.ListDetailResp{
		Ecses:      ecses,
		Finished:   isFinished,
		NextToken:  "",
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		RequestId:  "",
	}, nil
}

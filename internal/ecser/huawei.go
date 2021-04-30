package ecser

import (
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	hwecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	hwiam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iammodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	iamregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	"github.com/pkg/errors"
)

type HuaweiEcs struct {
	cli        *hwecs.EcsClient
	regionId   pbtenant.HuaweiRegionId
	regionName string
}

func NewHuaweiEcsClient(regionId int32, tenant tenanter.Tenanter) (Ecser, error) {
	var client *hwecs.EcsClient

	rName, err := tenanter.GetHuaweiRegionName(regionId)
	if err != nil {
		return nil, errors.WithMessage(err, "GetHuaweiRegionName error")
	}

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		auth := basic.NewCredentialsBuilder().WithAk(t.GetId()).WithSk(t.GetSecret()).Build()
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
		hcClient := hwecs.EcsClientBuilder().WithRegion(region.ValueOf(rName)).WithCredential(auth).Build()
		client = hwecs.NewEcsClient(hcClient)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init huawei ecs client error")
	}
	return &HuaweiEcs{
		cli:        client,
		regionId:   pbtenant.HuaweiRegionId(regionId),
		regionName: rName,
	}, nil
}

func (ecs *HuaweiEcs) DescribeInstances(pageNumber, pageSize int32) (*pbecs.ListResp, error) {
	req := new(model.ListServersDetailsRequest)
	offset := (pageNumber - 1) * pageSize
	req.Offset = &offset
	limit := pageSize
	req.Limit = &limit

	resp, err := ecs.cli.ListServersDetails(req)
	if err != nil {
		return nil, errors.Wrap(err, "Huawei DescribeInstances error")
	}

	servers := *resp.Servers
	var ecses = make([]*pbecs.ECSInstance, len(servers))
	for k, v := range servers {
		ecses[k] = &pbecs.ECSInstance{
			InstanceId:   v.Id,
			InstanceName: v.Name,
			RegionName:   ecs.regionName,
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
	return &pbecs.ListResp{
		Ecses:      ecses,
		NextToken:  "",
		PageNumber: 0,
		PageSize:   0,
		RequestId:  "",
	}, nil
}

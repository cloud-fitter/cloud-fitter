package configger

import (
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbstatistic"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"context"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	hwecs "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/model"
	hwregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ecs/v2/region"
	hwiam "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3"
	iammodel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/model"
	iamregion "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/iam/v3/region"
	"github.com/pkg/errors"
)

type HuaweiCfg struct {
	cli        *hwecs.EcsClient
	regionId   pbtenant.HuaweiRegionId
	regionName string
	tenanter.Tenanter
}

func NewHuaweiCfgClient(region tenanter.Region, tenant tenanter.Tenanter) (cfg Configger, err error) {
	var client *hwecs.EcsClient

	defer func() {
		if e := recover(); e != nil {
			err = errors.Errorf("NewHuaweiCfgClient panic %v", e)
		}
	}()

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
	return &HuaweiCfg{
		cli:        client,
		regionId:   pbtenant.HuaweiRegionId(region.GetId()),
		regionName: region.GetName(),
		Tenanter:   tenant,
	}, err
}

func (cfg *HuaweiCfg) Statistic(ctx context.Context) (*pbstatistic.StatisticInfo, error) {
	req := new(model.ListServersDetailsRequest)
	var offset int32 = 0
	var limit int32 = 1
	req.Offset = &offset
	req.Limit = &limit

	resp, err := cfg.cli.ListServersDetails(req)
	if err != nil {
		return nil, errors.Wrap(err, "Huawei ListServersDetails error")
	}

	return &pbstatistic.StatisticInfo{
		Provider:    pbtenant.CloudProvider_huawei,
		AccountName: cfg.AccountName(),
		Product:     pbtenant.CloudProduct_product_ecs,
		RegionId:    int32(cfg.regionId),
		RegionName:  cfg.regionName,
		Count:       *resp.Count,
	}, nil
}

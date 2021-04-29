package configger

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awscs "github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

type AwsCfg struct {
	cli        *awscs.Client
	regionId   pbtenant.AwsRegionId
	regionName string
	tenanter.Tenanter
}

func NewAwsCfgClient(regionId int32, tenant tenanter.Tenanter) (Configger, error) {
	var client *awscs.Client

	rName, err := tenanter.GetAwsRegionName(regionId)
	if err != nil {
		return nil, err
	}

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(t.GetId(), t.GetSecret(), "")),
			config.WithRegion(rName),
		)
		if err != nil {
			return nil, errors.Wrap(err, "LoadDefaultConfig aws ecs client error")
		}
		client = awscs.NewFromConfig(cfg)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init aws cfg client error")
	}
	return &AwsCfg{
		cli:        client,
		regionId:   pbtenant.AwsRegionId(regionId),
		regionName: rName,
		Tenanter:   tenant,
	}, nil
}

func (cfg *AwsCfg) Statistic() (*pbcfg.StatisticRespList, error) {
	req := new(awscs.GetDiscoveredResourceCountsInput)
	req.ResourceTypes = []string{}

	resp, err := cfg.cli.GetDiscoveredResourceCounts(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "Aws Statistic error")
	}

	result := &pbcfg.StatisticRespList{
		Provider:    pbtenant.CloudProvider_aws_cloud,
		AccountName: cfg.AccountName(),
		Product:     pbtenant.CloudProduct_product_ecs,
		RegionId:    int32(cfg.regionId),
		RegionName:  cfg.regionName,
		Count:       0,
	}

	for _, v := range resp.ResourceCounts {
		if v.ResourceType == types.ResourceTypeInstance {
			result.Count += v.Count
		}
	}
	return result, nil
}

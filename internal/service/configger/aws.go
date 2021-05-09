package configger

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awscs "github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbstatistic"
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

func NewAwsCfgClient(region tenanter.Region, tenant tenanter.Tenanter) (Configger, error) {
	var client *awscs.Client
	var err error

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(t.GetId(), t.GetSecret(), "")),
			config.WithRegion(region.GetName()),
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
		regionId:   pbtenant.AwsRegionId(region.GetId()),
		regionName: region.GetName(),
		Tenanter:   tenant,
	}, nil
}

func (cfg *AwsCfg) Statistic(ctx context.Context) (*pbstatistic.StatisticInfo, error) {
	req := new(awscs.GetDiscoveredResourceCountsInput)
	req.ResourceTypes = []string{}

	resp, err := cfg.cli.GetDiscoveredResourceCounts(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "Aws Statistic error")
	}

	result := &pbstatistic.StatisticInfo{
		Provider:    pbtenant.CloudProvider_aws,
		AccountName: cfg.AccountName(),
		Product:     pbtenant.CloudProduct_product_ecs,
		RegionId:    int32(cfg.regionId),
		RegionName:  cfg.regionName,
		Count:       0,
	}

	for _, v := range resp.ResourceCounts {
		if v.ResourceType == types.ResourceTypeInstance {
			result.Count += int32(v.Count)
		}
	}
	return result, nil
}

package ecser

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awscs "github.com/aws/aws-sdk-go-v2/service/configservice"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

type AwsEcs struct {
	cli *awscs.Client
}

func NewAwsEcsClient(tenant tenanter.Tenanter) (Ecser, error) {
	var client *awscs.Client

	var err error
	// rName, err := tenanter.GetHuaweiRegionName(regionId)
	// if err != nil {
	// 	return nil, err
	// }

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(t.GetId(), t.GetSecret(), "")),
			config.WithRegion("us-west-2"),
		)
		if err != nil {
			return nil, errors.Wrap(err, "LoadDefaultConfig aws ecs client error")
		}
		client = awscs.NewFromConfig(cfg)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init huawei ecs client error")
	}
	return &AwsEcs{cli: client}, nil
}

func (ecs *AwsEcs) ECSStatistic() (*pbecs.ECSStatisticRespList, error) {
	req := new(awscs.GetDiscoveredResourceCountsInput)
	req.ResourceTypes = []string{string(types.ResourceTypeInstance)}

	resp, err := ecs.cli.GetDiscoveredResourceCounts(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "Aws ECSStatistic error")
	}

	result := new(pbecs.ECSStatisticRespList)

	for _, v := range resp.ResourceCounts {
		result.Count += v.Count
	}
	return nil, nil
}

func (ecs *AwsEcs) DescribeInstances(pageNumber, pageSize int) ([]*pbecs.ECSInstance, error) {
	// req := new(awscs.GetDiscoveredResourceCountsInput)
	// req.ResourceTypes = []string{string(types.ResourceTypeInstance)}
	//
	// resp, err := ecs.cli.GetDiscoveredResourceCounts(context.Background(), req)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "Huawei DescribeInstances error")
	// }

	// var ecses = make([]*pbecs.ECSInstance, len(resp.ResourceCounts))
	// for k, v := range resp.ResourceCounts {
	// 	ecses[k] = &pbecs.ECSInstance{
	// 		InstanceName: v.ResourceType,
	// 		PublicIps:    []string{v.AccessIPv4},
	// 	}
	// }
	return nil, nil
}

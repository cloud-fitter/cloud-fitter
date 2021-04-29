package ecser

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
)

type AwsEcs struct {
	cli        *awsec2.Client
	regionId   pbtenant.AwsRegionId
	regionName string
}

func NewAwsEcsClient(regionId int32, tenant tenanter.Tenanter) (Ecser, error) {
	var client *awsec2.Client

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
		client = awsec2.NewFromConfig(cfg)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init aws ec2 client error")
	}
	return &AwsEcs{
		cli:        client,
		regionId:   pbtenant.AwsRegionId(regionId),
		regionName: rName,
	}, nil
}

func (ecs *AwsEcs) DescribeInstances(pageNumber, pageSize int32) (*pbecs.ListResp, error) {
	req := new(awsec2.DescribeInstancesInput)
	req.MaxResults = pageSize

	resp, err := ecs.cli.DescribeInstances(context.Background(), req)
	if err != nil {
		return nil, errors.Wrap(err, "Aws DescribeInstances error")
	}

	var ecses []*pbecs.ECSInstance
	for _, v := range resp.Reservations {
		for _, v2 := range v.Instances {
			ecses = append(ecses, &pbecs.ECSInstance{
				InstanceId:   *v2.InstanceId,
				InstanceName: "",
				RegionName:   "",
				PublicIps:    []string{*v2.PublicIpAddress},
				InstanceType: string(v2.InstanceType),
				Cpu:          v2.CpuOptions.CoreCount,
				Memory:       0,
				Description:  "",
				Status:       string(v2.State.Name),
				CreationTime: "",
				ExpireTime:   "",
			})
		}

	}
	return &pbecs.ListResp{
		Ecses: ecses,
	}, nil
}

package ecser

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

type AwsEcs struct {
	cli      *awsec2.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func NewAwsEcsClient(region tenanter.Region, tenant tenanter.Tenanter) (Ecser, error) {
	var (
		client *awsec2.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(t.GetId(), t.GetSecret(), "")),
			config.WithRegion(region.GetName()),
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
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (ecs *AwsEcs) ListDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	request := new(awsec2.DescribeInstancesInput)
	request.MaxResults = req.PageSize
	request.NextToken = &req.NextToken

	resp, err := ecs.cli.DescribeInstances(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "Aws ListDetail error")
	}

	var ecses []*pbecs.EcsInstance
	for _, v := range resp.Reservations {
		for _, v2 := range v.Instances {
			ecses = append(ecses, &pbecs.EcsInstance{
				Provider:     pbtenant.CloudProvider_aws,
				AccountName:  ecs.tenanter.AccountName(),
				InstanceId:   *v2.InstanceId,
				InstanceName: "",
				RegionName:   ecs.region.GetName(),
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

	if resp.NextToken != nil {
		return &pbecs.ListDetailResp{
			Ecses:     ecses,
			Finished:  false,
			NextToken: *resp.NextToken,
		}, nil
	}
	return &pbecs.ListDetailResp{
		Ecses:     ecses,
		Finished:  true,
		NextToken: "",
	}, nil
}

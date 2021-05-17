package rdser

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	awsrds "github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbrds"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

type AwsRds struct {
	cli      *awsrds.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newAwsRdsClient(region tenanter.Region, tenant tenanter.Tenanter) (Rdser, error) {
	var (
		client *awsrds.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(t.GetId(), t.GetSecret(), "")),
			config.WithRegion(region.GetName()),
		)
		if err != nil {
			return nil, errors.Wrap(err, "LoadDefaultConfig aws rds client error")
		}
		client = awsrds.NewFromConfig(cfg)
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init aws rds client error")
	}
	return &AwsRds{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (rds *AwsRds) ListDetail(ctx context.Context, req *pbrds.ListDetailReq) (*pbrds.ListDetailResp, error) {
	request := new(awsrds.DescribeDBInstancesInput)
	request.MaxRecords = &req.PageSize

	resp, err := rds.cli.DescribeDBInstances(ctx, request)
	if err != nil {
		return nil, errors.Wrap(err, "Aws ListDetail error")
	}

	var rdses = make([]*pbrds.RdsInstance, len(resp.DBInstances))
	for k, v := range resp.DBInstances {
		rdses[k] = &pbrds.RdsInstance{
			Provider:      pbtenant.CloudProvider_aws,
			AccoutName:    rds.tenanter.AccountName(),
			InstanceId:    *v.DBInstanceIdentifier,
			InstanceName:  "",
			RegionName:    rds.region.GetName(),
			InstanceType:  "",
			Engine:        *v.Engine,
			EngineVersion: *v.EngineVersion,
			InstanceClass: *v.DBInstanceClass,
			Status:        *v.DBInstanceStatus,
			CreationTime:  v.InstanceCreateTime.Format(time.RFC3339),
			ExpireTime:    "",
		}
	}

	if resp.Marker != nil {
		return &pbrds.ListDetailResp{
			Rdses:     rdses,
			Finished:  false,
			NextToken: *resp.Marker,
		}, nil
	}
	return &pbrds.ListDetailResp{
		Rdses:     rdses,
		Finished:  true,
		NextToken: "",
	}, nil
}

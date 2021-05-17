package ecser

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TencentCvm struct {
	cli      *cvm.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newTencentCvmClient(region tenanter.Region, tenant tenanter.Tenanter) (Ecser, error) {
	var (
		client *cvm.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = cvm.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), region.GetName(), profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent cvm client error")
	}
	return &TencentCvm{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (ecs *TencentCvm) ListDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	request := cvm.NewDescribeInstancesRequest()
	request.Offset = common.Int64Ptr(int64((req.PageNumber - 1) * req.PageSize))
	request.Limit = common.Int64Ptr(int64(req.PageSize))
	resp, err := ecs.cli.DescribeInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent ListDetail error")
	}

	var ecses = make([]*pbecs.EcsInstance, len(resp.Response.InstanceSet))
	for k, v := range resp.Response.InstanceSet {
		ecses[k] = &pbecs.EcsInstance{
			Provider:     pbtenant.CloudProvider_tencent,
			AccountName:  ecs.tenanter.AccountName(),
			InstanceId:   *v.InstanceId,
			InstanceName: *v.InstanceName,
			RegionName:   ecs.region.GetName(),
			InstanceType: *v.InstanceType,
			PublicIps:    make([]string, len(v.PublicIpAddresses)),
			Cpu:          int32(*v.CPU),
			Memory:       int32(*v.Memory),
			Description:  "",
			Status:       *v.InstanceState,
			CreationTime: *v.CreatedTime,
			ExpireTime:   *v.ExpiredTime,
		}
		for k1, v1 := range v.PublicIpAddresses {
			ecses[k].PublicIps[k1] = *v1
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
		RequestId:  *resp.Response.RequestId,
	}, nil
}

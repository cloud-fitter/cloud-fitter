package ecser

import (
	pbecs "github.com/cloud-fitter/cloud-fitter/gen/idl/ecs"
	pbtenant "github.com/cloud-fitter/cloud-fitter/gen/idl/tenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

type TencentCvm struct {
	cli *cvm.Client
}

func NewTencentCvmClient(regionId pbtenant.TencentRegionId, tenant tenanter.Tenanter) (Ecser, error) {
	var (
		client *cvm.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = cvm.NewClient(common.NewCredential(t.GetId(), t.GetSecret()),
			tenanter.GetTencentRegionId(regionId),
			profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent cvm client error")
	}
	return &TencentCvm{cli: client}, nil
}

func (ecs *TencentCvm) DescribeInstances(pageNumber, pageSize int) ([]*pbecs.ECSInstance, error) {
	req := cvm.NewDescribeInstancesRequest()
	req.Offset = common.Int64Ptr(int64((pageNumber - 1) * pageSize))
	req.Limit = common.Int64Ptr(int64(pageSize))
	resp, err := ecs.cli.DescribeInstances(req)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent DescribeInstances error")
	}

	var ecses = make([]*pbecs.ECSInstance, len(resp.Response.InstanceSet))
	for k, v := range resp.Response.InstanceSet {
		ecses[k] = &pbecs.ECSInstance{
			HostName:  *v.InstanceId,
			PublicIps: make([]string, len(v.PublicIpAddresses)),
		}
		for k1, v1 := range v.PublicIpAddresses {
			ecses[k].PublicIps[k1] = *v1
		}
	}
	return ecses, nil
}

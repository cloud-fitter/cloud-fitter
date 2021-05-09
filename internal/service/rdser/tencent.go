package rdser

import (
	"context"
	"fmt"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbrds"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"github.com/pkg/errors"
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

type TencentCdb struct {
	cli      *cdb.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func NewTencentCdbClient(region tenanter.Region, tenant tenanter.Tenanter) (Rdser, error) {
	var (
		client *cdb.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = cdb.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), region.GetName(), profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent cdb client error")
	}
	return &TencentCdb{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (rds *TencentCdb) ListDetail(ctx context.Context, req *pbrds.ListDetailReq) (*pbrds.ListDetailResp, error) {
	request := cdb.NewDescribeDBInstancesRequest()
	request.Offset = common.Uint64Ptr(uint64((req.PageNumber - 1) * req.PageSize))
	request.Limit = common.Uint64Ptr(uint64(req.PageSize))
	resp, err := rds.cli.DescribeDBInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent ListDetail error")
	}

	var rdses = make([]*pbrds.RdsInstance, len(resp.Response.Items))
	for k, v := range resp.Response.Items {
		rdses[k] = &pbrds.RdsInstance{
			Provider:      pbtenant.CloudProvider_tencent,
			AccoutName:    rds.tenanter.AccountName(),
			InstanceId:    *v.InstanceId,
			InstanceName:  *v.InstanceName,
			RegionName:    rds.region.GetName(),
			InstanceType:  fmt.Sprint(*v.InstanceType),
			Engine:        "",
			EngineVersion: *v.EngineVersion,
			InstanceClass: *v.DeviceType,
			Status:        fmt.Sprint(*v.Status),
			CreationTime:  *v.CreateTime,
			ExpireTime:    *v.DeadlineTime,
		}
	}

	isFinished := false
	if len(rdses) < int(req.PageSize) {
		isFinished = true
	}

	return &pbrds.ListDetailResp{
		Rdses:      rdses,
		Finished:   isFinished,
		NextToken:  "",
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		RequestId:  *resp.Response.RequestId,
	}, nil
}

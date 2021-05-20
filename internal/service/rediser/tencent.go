package rediser

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbredis"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"
)

// 0：待初始化；1：实例在流程中；2：实例运行中；-2：实例已隔离；-3：实例待删除
var tencentRedisStatus = map[int64]string{
	0:  "待初始化",
	1:  "实例在流程中",
	2:  "实例运行中",
	-2: "实例已隔离",
	-3: "实例待删除",
}

type TencentRedis struct {
	cli      *redis.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newTencentRedisClient(region tenanter.Region, tenant tenanter.Tenanter) (Rediser, error) {
	var (
		client *redis.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = redis.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), region.GetName(), profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent redis client error")
	}
	return &TencentRedis{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (rds *TencentRedis) ListDetail(ctx context.Context, req *pbredis.ListDetailReq) (*pbredis.ListDetailResp, error) {
	request := redis.NewDescribeInstancesRequest()
	request.Offset = common.Uint64Ptr(uint64((req.PageNumber - 1) * req.PageSize))
	request.Limit = common.Uint64Ptr(uint64(req.PageSize))
	resp, err := rds.cli.DescribeInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent ListDetail error")
	}

	var rdses = make([]*pbredis.RedisInstance, len(resp.Response.InstanceSet))
	for k, v := range resp.Response.InstanceSet {
		status, _ := tencentRedisStatus[*v.Status]
		rdses[k] = &pbredis.RedisInstance{
			Provider:     pbtenant.CloudProvider_tencent,
			AccoutName:   rds.tenanter.AccountName(),
			InstanceId:   *v.InstanceId,
			InstanceName: *v.InstanceName,
			RegionName:   rds.region.GetName(),
			Size:         int32(*v.Size),
			Status:       status,
			CreationTime: *v.Createtime,
			ExpireTime:   *v.DeadlineTime,
		}
	}

	isFinished := false
	if len(rdses) < int(req.PageSize) {
		isFinished = true
	}

	return &pbredis.ListDetailResp{
		Redises:    rdses,
		Finished:   isFinished,
		NextToken:  "",
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		RequestId:  *resp.Response.RequestId,
	}, nil
}

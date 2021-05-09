package rdser

import (
	"context"
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	alirds "github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbrds"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var aliClientMutex sync.Mutex

type AliRds struct {
	cli      *alirds.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func NewAliRdsClient(region tenanter.Region, tenant tenanter.Tenanter) (Rdser, error) {
	var (
		client *alirds.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		client, err = alirds.NewClientWithAccessKey(region.GetName(), t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali rds client error")
	}

	return &AliRds{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (rds *AliRds) ListDetail(ctx context.Context, req *pbrds.ListDetailReq) (*pbrds.ListDetailResp, error) {
	request := alirds.CreateDescribeDBInstancesRequest()
	request.PageNumber = requests.NewInteger(int(req.PageNumber))
	request.PageSize = requests.NewInteger(int(req.PageSize))
	resp, err := rds.cli.DescribeDBInstances(request)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun ListDetail error")
	}

	var rdses = make([]*pbrds.RdsInstance, len(resp.Items.DBInstance))
	for k, v := range resp.Items.DBInstance {
		rdses[k] = &pbrds.RdsInstance{
			Provider:      pbtenant.CloudProvider_ali,
			AccoutName:    rds.tenanter.AccountName(),
			InstanceId:    v.DBInstanceId,
			InstanceName:  v.DBInstanceDescription,
			RegionName:    rds.region.GetName(),
			InstanceType:  v.DBInstanceType,
			Engine:        v.Engine,
			EngineVersion: v.EngineVersion,
			InstanceClass: v.DBInstanceClass,
			Status:        v.DBInstanceStatus,
			CreationTime:  v.CreateTime,
			ExpireTime:    v.ExpireTime,
		}
	}

	isFinished := false
	if len(rdses) < int(req.PageSize) {
		isFinished = true
	}

	return &pbrds.ListDetailResp{
		Rdses:      rdses,
		Finished:   isFinished,
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		NextToken:  "",
		RequestId:  resp.RequestId,
	}, nil
}

package osser

import (
	"context"
	"sync"

	alioss "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pboss"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var aliClientMutex sync.Mutex

type AliOss struct {
	cli      *alioss.Client
	tenanter tenanter.Tenanter
}

func newAliOssClient(tenant tenanter.Tenanter) (Osser, error) {
	var (
		client *alioss.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		client, err = alioss.New("oss-cn-hangzhou.aliyuncs.com", t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali oss client error")
	}

	return &AliOss{
		cli:      client,
		tenanter: tenant,
	}, nil
}

func (oss *AliOss) ListDetail(ctx context.Context, req *pboss.ListDetailReq) (*pboss.ListDetailResp, error) {
	option := alioss.Range(int64(req.PageNumber*req.PageSize), int64((req.PageNumber+1)*req.PageSize))
	resp, err := oss.cli.ListBuckets(option)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun ListDetail error")
	}

	var osses = make([]*pboss.OssInstance, len(resp.Buckets))
	for k, v := range resp.Buckets {
		osses[k] = &pboss.OssInstance{
			Provider:    pbtenant.CloudProvider_ali,
			AccountName: oss.tenanter.AccountName(),
			BucketName:  v.Name,
			Location:    v.Location,
		}
	}

	isFinished := false
	if len(osses) < int(req.PageSize) {
		isFinished = true
	}

	return &pboss.ListDetailResp{
		Osses:      osses,
		Finished:   isFinished,
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		NextToken:  "",
		RequestId:  "",
	}, nil
}

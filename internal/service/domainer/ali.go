package domainer

import (
	"context"
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	alidomain "github.com/aliyun/alibaba-cloud-sdk-go/services/domain"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbdomain"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

var aliClientMutex sync.Mutex

type AliDomain struct {
	cli      *alidomain.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newAliDomainClient(tenant tenanter.Tenanter) (Rdser, error) {
	var (
		client *alidomain.Client
		err    error
	)

	hzRegion, _ := tenanter.NewRegion(pbtenant.CloudProvider_ali, int32(pbtenant.AliRegionId_ali_cn_hangzhou))

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		client, err = alidomain.NewClientWithAccessKey(hzRegion.GetName(), t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali domain client error")
	}

	return &AliDomain{
		cli:      client,
		region:   hzRegion,
		tenanter: tenant,
	}, nil
}

func (domain *AliDomain) ListDetail(ctx context.Context, req *pbdomain.ListDetailReq) (*pbdomain.ListDetailResp, error) {
	request := alidomain.CreateQueryDomainListRequest()
	request.PageNum = requests.NewInteger(int(req.PageNumber))
	request.PageSize = requests.NewInteger(int(req.PageSize))
	resp, err := domain.cli.QueryDomainList(request)
	if err != nil {
		return nil, errors.Wrap(err, "Aliyun ListDetail error")
	}

	var domains = make([]*pbdomain.DomainInstance, len(resp.Data.Domain))
	for k, v := range resp.Data.Domain {
		domains[k] = &pbdomain.DomainInstance{
			Provider:         pbtenant.CloudProvider_ali,
			AccountName:      domain.tenanter.AccountName(),
			DomainName:       v.DomainName,
			DomainStatus:     v.DomainStatus,
			DomainType:       v.DomainType,
			ExpirationDate:   v.ExpirationDate,
			RegistrationDate: v.RegistrationDate,
		}
	}

	isFinished := false
	if len(domains) < int(req.PageSize) {
		isFinished = true
	}

	return &pbdomain.ListDetailResp{
		Domains:    domains,
		Finished:   isFinished,
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		NextToken:  "",
		RequestId:  resp.RequestId,
	}, nil
}

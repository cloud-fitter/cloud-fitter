package domainer

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbdomain"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"github.com/pkg/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tcdomain "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/domain/v20180808"
)

type TencentDomain struct {
	cli      *tcdomain.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newTencentDomainClient(tenant tenanter.Tenanter) (Rdser, error) {
	var (
		client *tcdomain.Client
		err    error
	)

	region, _ := tenanter.NewRegion(pbtenant.CloudProvider_tencent, int32(pbtenant.TencentRegionId_tc_ap_guangzhou))

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = tcdomain.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), region.GetName(), profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent tcdomain client error")
	}
	return &TencentDomain{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (domain *TencentDomain) ListDetail(ctx context.Context, req *pbdomain.ListDetailReq) (*pbdomain.ListDetailResp, error) {
	request := tcdomain.NewDescribeDomainNameListRequest()
	request.Offset = common.Uint64Ptr(uint64((req.PageNumber - 1) * req.PageSize))
	request.Limit = common.Uint64Ptr(uint64(req.PageSize))
	resp, err := domain.cli.DescribeDomainNameList(request)
	if err != nil {
		return nil, errors.Wrap(err, "Tencent ListDetail error")
	}

	var domains = make([]*pbdomain.DomainInstance, len(resp.Response.DomainSet))
	for k, v := range resp.Response.DomainSet {
		domains[k] = &pbdomain.DomainInstance{
			Provider:         pbtenant.CloudProvider_tencent,
			AccountName:      domain.tenanter.AccountName(),
			DomainName:       *v.DomainName,
			DomainStatus:     *v.BuyStatus,
			DomainType:       "",
			ExpirationDate:   *v.ExpirationDate,
			RegistrationDate: *v.CreationDate,
		}
	}

	isFinished := false
	if len(domains) < int(req.PageSize) {
		isFinished = true
	}

	return &pbdomain.ListDetailResp{
		Domains:    domains,
		Finished:   isFinished,
		NextToken:  "",
		PageNumber: req.PageNumber + 1,
		PageSize:   req.PageSize,
		RequestId:  *resp.Response.RequestId,
	}, nil
}

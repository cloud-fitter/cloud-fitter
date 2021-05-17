package billinger

import (
	"context"
	"strconv"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbbilling"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"

	"github.com/pkg/errors"
	tcbilling "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const (
	tencentOffset = 100
)

type TencentBilling struct {
	cli      *tcbilling.Client
	region   tenanter.Region
	tenanter tenanter.Tenanter
}

func newTencentBillingClient(tenant tenanter.Tenanter) (Billinger, error) {
	var (
		client *tcbilling.Client
		err    error
	)

	region, _ := tenanter.NewRegion(pbtenant.CloudProvider_tencent, int32(pbtenant.TencentRegionId_tc_ap_guangzhou))

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		client, err = tcbilling.NewClient(common.NewCredential(t.GetId(), t.GetSecret()), region.GetName(), profile.NewClientProfile())
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init tencent billing client error")
	}
	return &TencentBilling{
		cli:      client,
		region:   region,
		tenanter: tenant,
	}, nil
}

func (domain *TencentBilling) ListDetail(ctx context.Context, req *pbbilling.ListDetailReq) (*pbbilling.ListDetailResp, error) {
	request := tcbilling.NewDescribeBillDetailRequest()
	request.Limit = common.Uint64Ptr(tencentOffset)
	request.PeriodType = common.StringPtr("byUsedTime")
	request.Month = common.StringPtr(req.BillingCycle)

	var result []*pbbilling.BillingInstance

	offset := 0
	for {
		request.Offset = common.Uint64Ptr(uint64(offset))
		resp, err := domain.cli.DescribeBillDetail(request)
		if err != nil {
			return nil, errors.Wrap(err, "Tencent ListDetail error")
		}

		var billings = make([]*pbbilling.BillingInstance, len(resp.Response.DetailSet))
		for k, v := range resp.Response.DetailSet {
			var costs float64
			for _, v2 := range v.ComponentSet {
				pay, _ := strconv.ParseFloat(*v2.RealCost, 64)
				costs += pay
			}
			billings[k] = &pbbilling.BillingInstance{
				Provider:         pbtenant.CloudProvider_tencent,
				BillingCycle:     req.BillingCycle,
				AccountName:      domain.tenanter.AccountName(),
				RegionName:       *v.RegionName,
				ProductCode:      *v.ProductCode,
				ProductType:      *v.ProductCodeName,
				InstanceId:       *v.ResourceId,
				Fee:              costs,
				Item:             *v.ActionTypeName,
				InstanceName:     *v.ResourceName,
				SubscriptionType: *v.PayModeName,
			}
		}
		result = append(result, billings...)
		if len(resp.Response.DetailSet) < tencentOffset {
			break
		}
		offset += tencentOffset
	}

	return &pbbilling.ListDetailResp{Billings: result}, nil
}

package billinger

import (
	"context"
	"sync"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	alibss "github.com/aliyun/alibaba-cloud-sdk-go/services/bssopenapi"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbbilling"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbtenant"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

const (
	aliyunBillMaxCount    = 300
	aliyunProductMaxCount = 100
)

var aliClientMutex sync.Mutex

type AliBss struct {
	cli      *alibss.Client
	tenanter tenanter.Tenanter
}

func newAliBillingClient(tenant tenanter.Tenanter) (Billinger, error) {
	var (
		client *alibss.Client
		err    error
	)

	switch t := tenant.(type) {
	case *tenanter.AccessKeyTenant:
		// 阿里云的sdk有一个 map 的并发问题，go test 加上-race 能检测出来，所以这里加一个锁
		aliClientMutex.Lock()
		client, err = alibss.NewClientWithAccessKey("cn-hangzhou", t.GetId(), t.GetSecret())
		aliClientMutex.Unlock()
	default:
	}

	if err != nil {
		return nil, errors.Wrap(err, "init ali billing client error")
	}

	return &AliBss{
		cli:      client,
		tenanter: tenant,
	}, nil
}

func (bss *AliBss) ListDetail(ctx context.Context, req *pbbilling.ListDetailReq) (*pbbilling.ListDetailResp, error) {
	var result []*pbbilling.BillingInstance

	request := alibss.CreateDescribeInstanceBillRequest()
	request.BillingCycle = req.BillingCycle
	request.MaxResults = requests.NewInteger(aliyunBillMaxCount)

	for {
		resp, err := bss.cli.DescribeInstanceBill(request)
		if err != nil {
			return nil, errors.Wrap(err, "Aliyun ListDetail error")
		}
		var billings = make([]*pbbilling.BillingInstance, len(resp.Data.Items))
		for k, v := range resp.Data.Items {
			billings[k] = &pbbilling.BillingInstance{
				Provider:         pbtenant.CloudProvider_ali,
				BillingCycle:     req.BillingCycle,
				AccountName:      bss.tenanter.AccountName(),
				RegionName:       v.Region,
				ProductCode:      v.ProductCode,
				ProductType:      v.ProductType,
				InstanceId:       v.InstanceID,
				Fee:              v.PretaxAmount,
				Item:             v.Item,
				InstanceName:     v.NickName,
				SubscriptionType: v.SubscriptionType,
			}
		}

		result = append(result, billings...)

		if len(resp.Data.Items) < aliyunBillMaxCount {
			break
		}

		request.NextToken = resp.Data.NextToken
	}

	return &pbbilling.ListDetailResp{
		Billings: result,
	}, nil
}

func (bss *AliBss) ListAllProducts(ctx context.Context) ([]*pbbilling.AliProductInfo, error) {
	var result []*pbbilling.AliProductInfo

	request := alibss.CreateQueryProductListRequest()
	page := 1
	request.PageSize = requests.NewInteger(aliyunProductMaxCount)

	for {
		request.PageNum = requests.NewInteger(page)
		resp, err := bss.cli.QueryProductList(request)
		if err != nil {
			return nil, errors.Wrap(err, "Aliyun ListDetail error")
		}

		var products = make([]*pbbilling.AliProductInfo, len(resp.Data.ProductList.Product))
		for k, v := range resp.Data.ProductList.Product {
			products[k] = &pbbilling.AliProductInfo{
				ProductCode:      v.ProductCode,
				ProductName:      v.ProductName,
				ProductType:      v.ProductType,
				SubscriptionType: v.SubscriptionType,
			}
		}
		result = append(result, products...)

		if len(products) < aliyunProductMaxCount {
			break
		}
		page++
	}

	return result, nil
}

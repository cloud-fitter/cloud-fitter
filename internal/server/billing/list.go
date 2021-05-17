package billing

import (
	"context"
	"sync"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbbilling"
	"github.com/cloud-fitter/cloud-fitter/internal/service/billinger"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
)

func ListDetail(ctx context.Context, req *pbbilling.ListDetailReq) (*pbbilling.ListDetailResp, error) {
	var (
		billing billinger.Billinger
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	for _, tenanter := range tenanters {
		if req.AccountName == "" || tenanter.AccountName() == req.AccountName {
			if billing, err = billinger.NewBillingClient(req.Provider, tenanter); err != nil {
				return nil, errors.WithMessage(err, "NewBillingClient error")
			}
			break
		}
	}

	return billing.ListDetail(ctx, req)
}

func List(ctx context.Context, req *pbbilling.ListReq) (*pbbilling.ListResp, error) {
	var (
		wg       sync.WaitGroup
		mutex    sync.Mutex
		billings []*pbbilling.BillingInstance
	)

	tenanters, err := tenanter.GetTenanters(req.Provider)
	if err != nil {
		return nil, errors.WithMessage(err, "getTenanters error")
	}

	wg.Add(len(tenanters))
	for _, t := range tenanters {
		go func(tenant tenanter.Tenanter) {
			defer wg.Done()
			billing, err := billinger.NewBillingClient(req.Provider, tenant)
			if err != nil {
				glog.Errorf("NewBillingClient error %v", err)
				return
			}

			request := &pbbilling.ListDetailReq{
				Provider:     req.Provider,
				BillingCycle: req.BillingCycle,
				AccountName:  tenant.AccountName(),
			}
			resp, err := billing.ListDetail(ctx, request)
			if err != nil {
				glog.Errorf("ListDetail error %v", err)
				return
			}
			mutex.Lock()
			billings = append(billings, resp.Billings...)
			mutex.Unlock()
		}(t)
	}
	wg.Wait()

	return &pbbilling.ListResp{
		Billings: billings,
	}, nil
}

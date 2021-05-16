package billing

import (
	"context"

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

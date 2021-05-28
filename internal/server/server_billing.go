package server

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbbilling"
	"github.com/cloud-fitter/cloud-fitter/internal/server/billing"
)

func (s *Server) ListBillingDetail(ctx context.Context, req *pbbilling.ListDetailReq) (*pbbilling.ListDetailResp, error) {
	resp, err := billing.ListDetail(ctx, req)
	if err != nil {
		glog.Errorf("ListBillingDetail error %+v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) ListBilling(ctx context.Context, req *pbbilling.ListReq) (*pbbilling.ListResp, error) {
	resp, err := billing.List(ctx, req)
	if err != nil {
		glog.Errorf("ListBilling error %+v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return resp, nil
}

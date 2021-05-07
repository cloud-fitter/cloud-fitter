package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbdomain"
)

func (s *Server) ListDomainDetail(ctx context.Context, req *pbdomain.ListDetailReq) (*pbdomain.ListDetailResp, error) {
	return &pbdomain.ListDetailResp{}, nil
}

func (s *Server) ListDomain(ctx context.Context, req *pbdomain.ListReq) (*pbdomain.ListResp, error) {
	return &pbdomain.ListResp{}, nil
}

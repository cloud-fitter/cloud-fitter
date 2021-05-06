package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbrds"
)

func (s *Server) ListRDSDetail(ctx context.Context, req *pbrds.ListDetailReq) (*pbrds.ListDetailResp, error) {
	return &pbrds.ListDetailResp{}, nil
}

func (s *Server) ListRDS(ctx context.Context, req *pbrds.ListReq) (*pbrds.ListResp, error) {
	return &pbrds.ListResp{}, nil
}

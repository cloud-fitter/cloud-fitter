package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pboss"
)

func (s *Server) ListOssDetail(ctx context.Context, req *pboss.ListDetailReq) (*pboss.ListDetailResp, error) {
	return &pboss.ListDetailResp{}, nil
}

func (s *Server) ListOss(ctx context.Context, req *pboss.ListReq) (*pboss.ListResp, error) {
	return &pboss.ListResp{}, nil
}

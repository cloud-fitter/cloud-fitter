package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/server/ecs"
	"github.com/pkg/errors"
)

func (s *Server) ListECSDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	resp, err := ecs.ListDetail(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "ecs list detail error")
	}
	return resp, nil
}

func (s *Server) ListECS(ctx context.Context, req *pbecs.ListReq) (*pbecs.ListResp, error) {
	resp, err := ecs.List(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "ecs list error")
	}
	return resp, nil
}

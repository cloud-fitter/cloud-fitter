package server

import (
	"context"

	"github.com/golang/glog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/server/ecs"
)

func (s *Server) ListEcsDetail(ctx context.Context, req *pbecs.ListDetailReq) (*pbecs.ListDetailResp, error) {
	resp, err := ecs.ListDetail(ctx, req)
	if err != nil {
		glog.Errorf("ListEcsDetail error %+v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return resp, nil
}

func (s *Server) ListEcs(ctx context.Context, req *pbecs.ListReq) (*pbecs.ListResp, error) {
	resp, err := ecs.List(ctx, req)
	if err != nil {
		glog.Errorf("ListEcs error %+v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return resp, nil
}

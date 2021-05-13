package server

import (
	"context"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbkafka"
)

func (s *Server) ListKafkaDetail(ctx context.Context, req *pbkafka.ListDetailReq) (*pbkafka.ListDetailResp, error) {
	return &pbkafka.ListDetailResp{}, nil
}

func (s *Server) ListKafka(ctx context.Context, req *pbkafka.ListReq) (*pbkafka.ListResp, error) {
	return &pbkafka.ListResp{}, nil
}

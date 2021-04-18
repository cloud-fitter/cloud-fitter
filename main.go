package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	gw "github.com/cloud-fitter/cloud-fitter/gen/idl/demo" // Update
)

type server struct {
	// pb.go中自动生成的，是个空结构体
	gw.YourServiceServer
}

func (s *server) Echo(ctx context.Context, req *gw.StringMessage) (*gw.StringMessage, error) {
	return &gw.StringMessage{
		Value: "Hello, My Friend! Welcome to Cloud Fitter",
	}, nil
}

var (
	// command-line options:
	// gRPC server endpoint
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			glog.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		gw.RegisterYourServiceServer(s, &server{})
		if err := s.Serve(lis); err != nil {
			glog.Fatalf("failed to serve: %v", err)
		}
	}()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

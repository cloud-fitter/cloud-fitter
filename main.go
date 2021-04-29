package main

import (
	"context"
	"flag"
	"net"
	"net/http"

	"github.com/cloud-fitter/cloud-fitter/gen/idl/demo" // Update
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbcfg"
	"github.com/cloud-fitter/cloud-fitter/gen/idl/pbecs"
	"github.com/cloud-fitter/cloud-fitter/internal/server"
	"github.com/cloud-fitter/cloud-fitter/internal/tenanter"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

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
	if err := demo.RegisterYourServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return errors.Wrap(err, "RegisterYourServiceHandlerFromEndpoint error")
	}

	if err := pbecs.RegisterECSServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return errors.Wrap(err, "RegisterECSServiceHandlerFromEndpoint error")
	}

	if err := pbcfg.RegisterStatisticServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return errors.Wrap(err, "RegisterStatisticServiceHandlerFromEndpoint error")
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8081", mux)
}

func main() {
	var configFile string
	flag.StringVar(&configFile, "conf", "config.yaml", "config file path")
	flag.Parse()
	defer glog.Flush()

	if err := tenanter.LoadCloudConfigsFromFile(configFile); err != nil {
		if !errors.Is(err, tenanter.ErrLoadTenanterFileEmpty) {
			glog.Fatalf("LoadCloudConfigsFromFile error %+v", err)
		}
		glog.Warningf("LoadCloudConfigsFromFile empty file path %s", configFile)
	}

	glog.Infof("load tenant from file finished")

	go func() {
		lis, err := net.Listen("tcp", ":9090")
		if err != nil {
			glog.Fatalf("failed to listen: %v", err)
		}

		s := grpc.NewServer()
		demo.RegisterYourServiceServer(s, &server.Server{})
		pbecs.RegisterECSServiceServer(s, &server.Server{})
		pbcfg.RegisterStatisticServiceServer(s, &server.Server{})

		if err = s.Serve(lis); err != nil {
			glog.Fatalf("failed to serve: %v", err)
		}
	}()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}

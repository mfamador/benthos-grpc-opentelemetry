package main

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"log"
	"net"

	_ "github.com/benthosdev/benthos/v4/public/components/all"
	"github.com/benthosdev/benthos/v4/public/service"
	"github.com/mfamador/go-opentelemetry/grpcservice"
	_ "github.com/mfamador/go-opentelemetry/processor"
	"github.com/mfamador/go-opentelemetry/servicev1"
	"google.golang.org/grpc"
)

func main() {
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 8181))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		grpcServer := grpc.NewServer(grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()))
		server := grpcservice.ServiceImpl{}
		servicev1.RegisterServiceServer(grpcServer, server)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
	service.RunCLI(context.Background())
}

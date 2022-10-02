package grpcservice

import (
	"context"
	"github.com/mfamador/go-opentelemetry/servicev1"
)

// ServiceImpl is the gRPC server implementation
type ServiceImpl struct{}

func (s ServiceImpl) Ping(context.Context, *servicev1.PingRequest) (*servicev1.PingResponse, error) {
	return &servicev1.PingResponse{Message: "bar"}, nil
}

package server

import (
	"context"
	ttrace "grpcwrapper/trace"
	"testing"

	"google.golang.org/grpc"
)

func TextTraceUnaryServerInterceptor(t *testing.T) {
	ttrace.StartAgent()
	defer ttrace.StopAgent()

	TraceUnaryServerInterceptor()(context.Background(), nil, &grpc.UnaryServerInfo{
		FullMethod: "/helloworld.Greeter/SayHello",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})

	TraceUnaryServerInterceptor()(context.Background(), nil, &grpc.UnaryServerInfo{
		FullMethod: "/helloworld.Greeter/SayBay",
	}, func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, nil
	})
}

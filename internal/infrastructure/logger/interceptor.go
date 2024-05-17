package logger

import (
	"context"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx = stackTrace.Init(ctx)
	return handler(ctx, req)
}

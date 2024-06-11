package logger

import (
	"context"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx = stackTrace.Init(ctx, &stackTrace.Option{
		IsLock: false,
	})
	stackTrace.Add(ctx, "handler: GRPC, method: "+info.FullMethod)
	return handler(ctx, req)
}

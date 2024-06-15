package testCtx

import (
	"context"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
)

func New() context.Context {
	ctx := context.Background()
	ctx = stacktrace.Init(ctx, &stacktrace.Option{Capacity: 5})
	return ctx
}

package testCtx

import (
	"context"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
)

const (
	Trace1 = "first"
	Trace2 = "second"
)

func New() context.Context {
	ctx := context.Background()
	ctx = stacktrace.Init(ctx, &stacktrace.Option{Capacity: 5})
	stacktrace.Add(ctx, Trace1)
	stacktrace.Add(ctx, Trace2)
	return ctx
}

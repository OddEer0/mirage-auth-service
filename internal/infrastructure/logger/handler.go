package logger

import (
	"context"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

type AppLogHandler struct {
	slog.Handler
}

func (a AppLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level == slog.LevelError {
		sTrace, ok := ctx.Value(stackTrace.Key).(*stackTrace.StackTrace)
		if ok && sTrace.IsLock == 0 {
			r.AddAttrs(slog.Any("stack_trace", stackTrace.GetStack(ctx)))
		}
		stackTrace.Lock(ctx)
	}
	return a.Handler.Handle(ctx, r)
}

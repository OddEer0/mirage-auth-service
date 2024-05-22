package logger

import (
	"context"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"log/slog"
)

type AppLogHandler struct {
	slog.Handler
}

func (a AppLogHandler) Handle(ctx context.Context, r slog.Record) error {
	if r.Level == slog.LevelError {
		if !stackTrace.IsLock(ctx) {
			r.AddAttrs(slog.String("stack_trace", stackTrace.Get(ctx)))
		}
		stackTrace.Lock(ctx)
	}
	return a.Handler.Handle(ctx, r)
}

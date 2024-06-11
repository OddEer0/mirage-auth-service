package domain

import (
	"context"
)

type (
	Logger interface {
		ErrorContext(ctx context.Context, message string, attrs ...any)
		InfoContext(ctx context.Context, message string, attrs ...any)
		DebugContext(ctx context.Context, message string, attrs ...any)
		WarnContext(ctx context.Context, message string, attrs ...any)
		Error(message string, attrs ...any)
		Info(message string, attrs ...any)
		Warn(message string, attrs ...any)
		Debug(message string, attrs ...any)
	}
)

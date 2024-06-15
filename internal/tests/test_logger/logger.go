package testLogger

import (
	"context"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
)

type Logger struct {
	Message string
	Stack   []any
}

func (l *Logger) ErrorContext(ctx context.Context, message string, attrs ...any) {
	l.Message = message
	l.Stack = stacktrace.GetStack(ctx)
}

func (l *Logger) InfoContext(ctx context.Context, message string, attrs ...any) {
	l.Message = message
	l.Stack = stacktrace.GetStack(ctx)
}

func (l *Logger) DebugContext(ctx context.Context, message string, attrs ...any) {
	l.Message = message
	l.Stack = stacktrace.GetStack(ctx)
}

func (l *Logger) WarnContext(ctx context.Context, message string, attrs ...any) {
	l.Message = message
	l.Stack = stacktrace.GetStack(ctx)
}

func (l *Logger) Error(message string, attrs ...any) {
	l.Message = message
}

func (l *Logger) Info(message string, attrs ...any) {
	l.Message = message
}

func (l *Logger) Warn(message string, attrs ...any) {
	l.Message = message
}

func (l *Logger) Debug(message string, attrs ...any) {
	l.Message = message
}

func (l *Logger) Clean() {
	l.Message = ""
	l.Stack = nil
}

func New() *Logger {
	return &Logger{}
}

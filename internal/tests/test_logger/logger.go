package testLogger

import (
	"context"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
)

type Logger struct {
	Message string
	Stack   []any
	Lock    bool
}

func (l *Logger) ErrorContext(ctx context.Context, message string, attrs ...any) {
	if l.Lock {
		return
	}
	l.Message = message
	stack := stacktrace.GetStack(ctx)
	copyStack := make([]any, 0, len(stack))
	for _, trace := range stack {
		copyStack = append(copyStack, trace)
	}
	l.Stack = copyStack
	l.Lock = true
}

func (l *Logger) InfoContext(ctx context.Context, message string, attrs ...any) {
	l.Message = message
	stack := stacktrace.GetStack(ctx)
	copyStack := make([]any, 0, len(stack))
	for _, trace := range stack {
		copyStack = append(copyStack, trace)
	}
	l.Stack = copyStack
}

func (l *Logger) DebugContext(ctx context.Context, message string, attrs ...any) {
	l.Message = message
	stack := stacktrace.GetStack(ctx)
	copyStack := make([]any, 0, len(stack))
	for _, trace := range stack {
		copyStack = append(copyStack, trace)
	}
	l.Stack = copyStack
}

func (l *Logger) WarnContext(ctx context.Context, message string, attrs ...any) {
	l.Message = message
	stack := stacktrace.GetStack(ctx)
	copyStack := make([]any, 0, len(stack))
	for _, trace := range stack {
		copyStack = append(copyStack, trace)
	}
	l.Stack = copyStack
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
	l.Lock = false
}

func New() *Logger {
	return &Logger{}
}

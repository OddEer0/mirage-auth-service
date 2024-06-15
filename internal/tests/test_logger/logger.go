package testLogger

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
)

type logger struct {
}

func (l logger) ErrorContext(ctx context.Context, message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func (l logger) InfoContext(ctx context.Context, message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func (l logger) DebugContext(ctx context.Context, message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func (l logger) WarnContext(ctx context.Context, message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Error(message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Info(message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Warn(message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func (l logger) Debug(message string, attrs ...any) {
	//TODO implement me
	panic("implement me")
}

func New() domain.Logger {
	return &logger{}
}

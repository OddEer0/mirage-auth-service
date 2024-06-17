package pgUserRepository

import (
	"context"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (p *userRepository) Delete(ctx context.Context, id string) error {
	stacktrace.Add(ctx, TraceDelete)
	defer stacktrace.Done(ctx)

	has, err := p.HasUserById(ctx, id)
	if err != nil {
		p.log.ErrorContext(ctx, "Has user by id method error", slog.Any("cause", err), slog.String("id", id))
		return ErrInternal
	}
	if !has {
		p.log.ErrorContext(ctx, "User not found", slog.Any("cause", ErrUserNotFound), slog.String("id", id))
		return ErrUserNotFound
	}

	_, err = p.db.Exec(ctx, DeleteUserByIdQuery, id)
	if err != nil {
		p.log.ErrorContext(ctx, "Delete query error", slog.Any("cause", err), slog.String("id", id))
		return ErrInternal
	}

	return nil
}

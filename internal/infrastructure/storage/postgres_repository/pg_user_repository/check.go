package pgUserRepository

import (
	"context"
	"database/sql"
	"errors"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (p *userRepository) CheckUserRole(ctx context.Context, id, role string) (bool, error) {
	stacktrace.Add(ctx, TraceCheckUserRole)
	defer stacktrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, CheckUserRoleQuery, id, role).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "User not found", slog.Any("cause", err), slog.String("id", id))
			return exists, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return exists, ErrInternal
	}
	return exists, nil
}

func (p *userRepository) HasUserById(ctx context.Context, id string) (bool, error) {
	stacktrace.Add(ctx, TraceHasUserById)
	defer stacktrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, HasUserByIdQuery, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return exists, ErrInternal
	}
	return exists, nil
}

func (p *userRepository) HasUserByLogin(ctx context.Context, login string) (bool, error) {
	stacktrace.Add(ctx, TraceHasUserByLogin)
	defer stacktrace.Done(ctx)
	var exists bool
	err := p.db.QueryRow(ctx, HasUserByLoginQuery, login).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("login", login))
		return exists, ErrInternal
	}
	return exists, nil
}

func (p *userRepository) HasUserByEmail(ctx context.Context, email string) (bool, error) {
	stacktrace.Add(ctx, TraceHasUserByEmail)
	defer stacktrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, HasUserByEmailQuery, email).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("email", email))
		return exists, ErrInternal
	}
	return exists, nil
}

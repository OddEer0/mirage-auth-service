package pgUserRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (p *userRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	stacktrace.Add(ctx, TraceGetByLogin)
	defer stacktrace.Done(ctx)

	row := p.db.QueryRow(ctx, GetUserByLoginQuery, login)
	var user model.User
	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Role, &user.IsBanned, &user.BanReason, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "user by login not found", slog.Any("cause", err), slog.String("login", login))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "Error db query", slog.Any("cause", err), slog.String("login", login))
		return nil, ErrInternal
	}
	return &user, nil
}

func (p *userRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	stacktrace.Add(ctx, TraceGetById)
	defer stacktrace.Done(ctx)

	row := p.db.QueryRow(ctx, GetUserByIdQuery, id)
	var user model.User
	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Role, &user.IsBanned, &user.BanReason, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "user by id not found", slog.Any("cause", err), slog.String("id", id))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "error db query", slog.Any("cause", err), slog.String("id", id))
		return nil, ErrInternal
	}
	return &user, nil
}

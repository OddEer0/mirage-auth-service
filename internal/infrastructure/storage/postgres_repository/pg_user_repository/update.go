package pgUserRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (p *userRepository) UpdateById(ctx context.Context, user *model.User) (*model.User, error) {
	stacktrace.Add(ctx, TraceUpdateById)
	defer stacktrace.Done(ctx)

	row := p.db.QueryRow(ctx, UpdateUserById, user.Id, user.Login, user.Email, user.Password, user.Role, user.IsBanned, user.BanReason, user.UpdatedAt)
	updatedUser := &model.User{}
	err := row.Scan(
		&updatedUser.Id,
		&updatedUser.Login,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.Role,
		&updatedUser.IsBanned,
		&updatedUser.BanReason,
		&updatedUser.UpdatedAt,
		&updatedUser.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "user not found", slog.Any("cause", err), slog.String("id", user.Id))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "update query and scan error", slog.Any("cause", err), "update_data", user)
		return nil, ErrInternal
	}

	return updatedUser, nil
}

func (p *userRepository) UpdateRoleById(ctx context.Context, id string, role string) (*model.User, error) {
	stacktrace.Add(ctx, TraceUpdateRoleById)
	defer stacktrace.Done(ctx)

	row := p.db.QueryRow(ctx, UpdateUserRoleById, id, role)
	updatedUser := &model.User{}
	err := row.Scan(
		&updatedUser.Id,
		&updatedUser.Login,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.Role,
		&updatedUser.IsBanned,
		&updatedUser.BanReason,
		&updatedUser.UpdatedAt,
		&updatedUser.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "user not found", slog.Any("cause", err), slog.String("id", id), slog.String("role", role))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "update role query and scan error", slog.Any("cause", err), slog.String("id", id), slog.String("role", role))
		return nil, ErrInternal
	}

	return updatedUser, nil
}

func (p *userRepository) UpdatePasswordById(ctx context.Context, id string, password string) (*model.User, error) {
	stacktrace.Add(ctx, TraceUpdatePasswordById)
	defer stacktrace.Done(ctx)

	row := p.db.QueryRow(ctx, UpdateUserPasswordById, id, password)
	updatedUser := &model.User{}
	err := row.Scan(
		&updatedUser.Id,
		&updatedUser.Login,
		&updatedUser.Email,
		&updatedUser.Password,
		&updatedUser.Role,
		&updatedUser.IsBanned,
		&updatedUser.BanReason,
		&updatedUser.UpdatedAt,
		&updatedUser.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "user not found", slog.Any("cause", err), slog.String("id", id))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "update password query and scan error", slog.Any("cause", err), slog.String("id", id))
		return nil, ErrInternal
	}

	return updatedUser, nil
}

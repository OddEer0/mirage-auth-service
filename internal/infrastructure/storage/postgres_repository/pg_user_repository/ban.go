package pgUserRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (p *userRepository) BanUserById(ctx context.Context, id string, banReason string) (*model.User, error) {
	stacktrace.Add(ctx, TraceBanUserById)
	defer stacktrace.Done(ctx)

	row := p.db.QueryRow(ctx, UpdateUserBanById, id, true, banReason)
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
			p.log.ErrorContext(ctx, "user not found", slog.Any("cause", err), slog.String("id", id), slog.String("ban_reason", banReason))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "ban user query and scan error", slog.Any("cause", err), slog.String("id", id), slog.String("ban_reason", banReason))
		return nil, ErrInternal
	}

	return updatedUser, nil

}

func (p *userRepository) UnbanUserById(ctx context.Context, id string) (*model.User, error) {
	stacktrace.Add(ctx, TraceUnbanUserById)
	defer stacktrace.Done(ctx)

	row := p.db.QueryRow(ctx, UpdateUserBanById, id, false, nil)
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
		p.log.ErrorContext(ctx, "unban user query and scan error", slog.Any("cause", err), slog.String("id", id))
		return nil, ErrInternal
	}

	return updatedUser, nil
}

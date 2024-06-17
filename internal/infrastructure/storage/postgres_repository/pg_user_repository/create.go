package pgUserRepository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (p *userRepository) Create(ctx context.Context, data *model.User) (*model.User, error) {
	stacktrace.Add(ctx, TraceCreate)
	defer stacktrace.Done(ctx)

	var createdUser model.User
	row := p.db.QueryRow(
		ctx,
		CreateUserQuery,
		data.Id,
		data.Login,
		data.Email,
		data.Password,
		data.Role,
		data.IsBanned,
		data.BanReason,
		data.UpdatedAt,
		data.CreatedAt,
	)
	err := row.Scan(
		&createdUser.Id,
		&createdUser.Login,
		&createdUser.Email,
		&createdUser.Password,
		&createdUser.Role,
		&createdUser.IsBanned,
		&createdUser.BanReason,
		&createdUser.UpdatedAt,
		&createdUser.CreatedAt,
	)
	if err != nil {
		p.log.ErrorContext(ctx, "error create new user", slog.Any("cause", err))
		return nil, ErrInternal
	}
	return &createdUser, nil
}

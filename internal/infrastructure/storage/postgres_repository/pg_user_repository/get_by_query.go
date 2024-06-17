package pgUserRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	stacktrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (p *userRepository) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*model.User, uint, error) {
	stacktrace.Add(ctx, TraceGetByQuery)
	defer stacktrace.Done(ctx)

	offset := query.PaginationQuery.PageCount * (query.PaginationQuery.CurrentPage - 1)
	limit := query.PaginationQuery.PageCount
	pageCount, err := postgresRepository.GetPageCount(ctx, p.db, "users", query.PaginationQuery.PageCount)
	if err != nil {
		p.log.ErrorContext(ctx, "getPageCount fn error", slog.Any("cause", err))
		return nil, pageCount, err
	}

	// TODO - аллоциируется память на каждый вызов функций, попытаться сделать константой
	queryStr := `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM Users
		ORDER BY ` + query.OrderQuery.OrderBy + " " + query.OrderQuery.OrderDirection + `
		LIMIT $1 OFFSET $2;
	`

	rows, err := p.db.Query(ctx, queryStr, limit, offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "not found users by query", slog.Any("cause", err), "query", query)
			return nil, 0, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "db query error", slog.Any("cause", err), "query", query)
		return nil, 0, ErrInternal
	}

	users := make([]*model.User, 0, limit)
	for rows.Next() {
		data := model.User{}
		err := rows.Scan(&data.Id, &data.Login, &data.Email, &data.Password, &data.Role, &data.IsBanned, &data.BanReason, &data.UpdatedAt, &data.CreatedAt)
		if err != nil {
			p.log.ErrorContext(ctx, "rows scan error", slog.Any("cause", err), "query", query)
			return nil, 0, ErrInternal
		}
		users = append(users, &data)
	}

	if err = rows.Err(); err != nil {
		p.log.ErrorContext(ctx, "rows error", slog.Any("cause", err))
		return nil, 0, ErrInternal
	}

	return users, pageCount, nil
}

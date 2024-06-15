package postgresRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

type userRepository struct {
	log domain.Logger
	db  postgres.Query
}

func (p *userRepository) CheckUserRole(ctx context.Context, id, role string) (bool, error) {
	stackTrace.Add(ctx, TraceUserRepoCheckUserRole)
	defer stackTrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, CheckUserRoleQuery, id, role).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return exists, domain.NewErr(domain.ErrInternalCode, domain.ErrNotFoundMessage)
	}
	return exists, nil
}

func (p *userRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	stackTrace.Add(ctx, TraceUserRepoGetByLogin)
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, GetUserByLoginQuery, login)
	var user model.User
	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Role, &user.IsBanned, &user.BanReason, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, LogNoRowMessage, slog.Any("cause", err), slog.String("login", login))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, LogDbQueryMessage, slog.Any("cause", err), slog.String("login", login))
		return nil, ErrInternal
	}
	return &user, nil
}

func (p *userRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	stackTrace.Add(ctx, TraceUserRepoGetById)
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, GetUserByIdQuery, id)
	var user model.User
	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Role, &user.IsBanned, &user.BanReason, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, LogNoRowMessage, slog.Any("cause", err), slog.String("id", id))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, LogDbQueryMessage, slog.Any("cause", err), slog.String("id", id))
		return nil, ErrInternal
	}
	return &user, nil
}

func (p *userRepository) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*model.User, uint, error) {
	stackTrace.Add(ctx, TraceUserRepoGetByQuery)
	defer stackTrace.Done(ctx)

	offset := query.PaginationQuery.PageCount * (query.PaginationQuery.CurrentPage - 1)
	limit := query.PaginationQuery.PageCount
	pageCount, err := getPageCount(ctx, p.db, "users", query.PaginationQuery.PageCount)
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

func (p *userRepository) Delete(ctx context.Context, id string) error {
	stackTrace.Add(ctx, TraceUserRepoDelete)
	defer stackTrace.Done(ctx)

	has, err := p.HasUserById(ctx, id)
	if err != nil {
		p.log.ErrorContext(ctx, "Has user by id method error", slog.Any("cause", err), slog.String("id", id))
		return ErrInternal
	}
	if !has {
		p.log.ErrorContext(ctx, "User not found", slog.Any("cause", err), slog.String("id", id))
		return ErrUserNotFound
	}

	_, err = p.db.Exec(ctx, DeleteUserById, id)
	if err != nil {
		p.log.ErrorContext(ctx, "Delete query error", slog.Any("cause", err), slog.String("id", id))
		return ErrInternal
	}

	return nil
}

func (p *userRepository) UpdateById(ctx context.Context, user *model.User) (*model.User, error) {
	stackTrace.Add(ctx, TraceUserRepoUpdateById)
	defer stackTrace.Done(ctx)

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
	stackTrace.Add(ctx, TraceUserRepoUpdateRoleById)
	defer stackTrace.Done(ctx)

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
	stackTrace.Add(ctx, TraceUserRepoUpdatePasswordById)
	defer stackTrace.Done(ctx)

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

func (p *userRepository) BanUserById(ctx context.Context, id string, banReason string) (*model.User, error) {
	stackTrace.Add(ctx, TraceUserRepoBanUserById)
	defer stackTrace.Done(ctx)

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
	stackTrace.Add(ctx, TraceUserRepoUnbanUserById)
	defer stackTrace.Done(ctx)

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

func (p *userRepository) HasUserById(ctx context.Context, id string) (bool, error) {
	stackTrace.Add(ctx, TraceUserRepoHasUserById)
	defer stackTrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, HasUserByIdQuery, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return exists, domain.NewErr(domain.ErrInternalCode, domain.ErrNotFoundMessage)
	}
	return exists, nil
}

func (p *userRepository) HasUserByLogin(ctx context.Context, login string) (bool, error) {
	stackTrace.Add(ctx, TraceUserRepoHasUserByLogin)
	defer stackTrace.Done(ctx)
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
	stackTrace.Add(ctx, TraceUserRepoHasUserByEmail)
	defer stackTrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, HasUserByEmailQuery, "email", email).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("email", email))
		return exists, ErrInternal
	}
	return exists, nil
}

func (p *userRepository) Create(ctx context.Context, data *model.User) (*model.User, error) {
	stackTrace.Add(ctx, TraceUserRepoCreate)
	defer stackTrace.Done(ctx)

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

func NewUserRepository(logger domain.Logger, db postgres.Query) repository.UserRepository {
	return &userRepository{db: db, log: logger}
}

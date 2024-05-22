package postgresRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

const (
	createUserQuery = ` 
		INSERT INTO users (id, login, email, password, role, isBanned, banReason, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	getUserByIdQuery = `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM users
		WHERE id = $1;
	`
	hasUserByLoginQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)
	`
	hasUserByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)
	`
	hasUserByEmailQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)
	`
	deleteUserById = `
		DELETE FROM users WHERE id = $1
	`
)

type postgresRepository struct {
	log *slog.Logger
	db  *pgx.Conn
}

func (p *postgresRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: GetById")
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, getUserByIdQuery, id)
	var user model.User
	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Role, &user.IsBanned, &user.BanReason, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "not found user by id", slog.Any("cause", err), slog.String("id", id))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return nil, ErrInternal
	}
	return &user, nil
}

func (p *postgresRepository) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*model.User, uint, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: GetById")
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
			return nil, 0, domain.NewErr(domain.ErrNotFoundCode, domain.ErrNotFoundMessage)
		}
		p.log.ErrorContext(ctx, "db query error", slog.Any("cause", err), "query", query)
		return nil, 0, domain.NewErr(domain.ErrInternalCode, domain.ErrInternalMessage)
	}

	users := make([]*model.User, 0, limit)
	for rows.Next() {
		data := model.User{}
		err := rows.Scan(&data.Id, &data.Login, &data.Email, &data.Password, &data.Role, &data.IsBanned, &data.BanReason, &data.UpdatedAt, &data.CreatedAt)
		if err != nil {
			p.log.ErrorContext(ctx, "rows scan error", slog.Any("cause", err), "query", query)
			return nil, 0, err
		}
		users = append(users, &data)
	}

	if err = rows.Err(); err != nil {
		p.log.ErrorContext(ctx, "rows error", slog.Any("cause", err))
		return nil, 0, err
	}

	return users, pageCount, nil
}

func (p *postgresRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (p *postgresRepository) UpdateById(ctx context.Context, user *model.CreateUser) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) UpdateRoleById(ctx context.Context, id string, role string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) UpdatePasswordById(ctx context.Context, id string, password string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) BanUserById(ctx context.Context, id string, banReason string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) UnbanUserById(ctx context.Context, id string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) HasUserById(ctx context.Context, id string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: HasUserByLogin")
	defer stackTrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, hasUserByIdQuery, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "Error user by id not found", slog.Any("cause", err), slog.String("id", id))
			return false, domain.NewErr(domain.ErrNotFoundCode, domain.ErrInternalMessage)
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return false, domain.NewErr(domain.ErrInternalCode, domain.ErrNotFoundMessage)
	}
	return exists, nil
}

func (p *postgresRepository) HasUserByLogin(ctx context.Context, login string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: HasUserByLogin")
	defer stackTrace.Done(ctx)
	var exists bool
	err := p.db.QueryRow(ctx, hasUserByLoginQuery, login).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "Error user by login not found", slog.Any("cause", err), slog.String("login", login))
			return false, domain.NewErr(domain.ErrNotFoundCode, domain.ErrInternalMessage)
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("login", login))
		return false, domain.NewErr(domain.ErrInternalCode, domain.ErrNotFoundMessage)
	}
	return exists, nil
}

func (p *postgresRepository) HasUserByEmail(ctx context.Context, email string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: HasUserByLogin")
	defer stackTrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, hasUserByEmailQuery, "email", email).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "Error user by email not found", slog.Any("cause", err), slog.String("email", email))
			return false, domain.NewErr(domain.ErrNotFoundCode, domain.ErrInternalMessage)
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("email", email))
		return false, domain.NewErr(domain.ErrInternalCode, domain.ErrNotFoundMessage)
	}
	return exists, nil
}

func (p *postgresRepository) Create(ctx context.Context, data *model.User) (*model.User, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: Create")
	defer stackTrace.Done(ctx)

	var createdUser model.User
	row := p.db.QueryRow(
		ctx,
		createUserQuery,
		data.Id,
		data.Login,
		data.Email,
		data.Password,
		data.Role,
		data.IsBanned,
		data.BanReason, data.UpdatedAt,
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
		return nil, domain.NewErr(domain.ErrInternalCode, domain.ErrInternalMessage)
	}
	return &createdUser, nil
}

func NewUserRepository(logger *slog.Logger, db *pgx.Conn) repository.UserRepository {
	return &postgresRepository{db: db, log: logger}
}

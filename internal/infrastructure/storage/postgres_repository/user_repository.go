package postgresRepository

import (
	"context"
	"database/sql"
	"errors"
	domainError "github.com/OddEer0/mirage-auth-service/internal/domain/domain_error"
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
	hasUserByFieldQuery = `
		SELECT 1 FROM users WHERE $1 = $2;
	`
	deleteUserByField = `
		DELETE FROM users WHERE $1 = $2
	`
)

type postgresRepository struct {
	log *slog.Logger
	db  *pgx.Conn
}

func (p *postgresRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	row := p.db.QueryRow(ctx, getUserByIdQuery, id)
	var user model.User
	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Role, &user.IsBanned, &user.BanReason, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domainError.NotFound
		}
		return nil, err
	}
	return &user, nil
}

func (p *postgresRepository) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*model.User, uint, error) {
	offset := query.PaginationQuery.PageCount * (query.PaginationQuery.CurrentPage - 1)
	limit := query.PaginationQuery.PageCount
	pageCount, err := getPageCount(ctx, p.db, "users", query.PaginationQuery.PageCount)
	if err != nil {
		return nil, pageCount, err
	}

	queryStr := `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM Users
		ORDER BY ` + query.OrderQuery.OrderBy + " " + query.OrderQuery.OrderDirection + `
		LIMIT $1 OFFSET $2;
	`

	rows, err := p.db.Query(ctx, queryStr, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	users := make([]*model.User, 0, limit)
	for rows.Next() {
		data := model.User{}
		err := rows.Scan(&data.Id, &data.Login, &data.Email, &data.Password, &data.Role, &data.IsBanned, &data.BanReason, &data.UpdatedAt, &data.CreatedAt)
		if err != nil {
			return nil, 0, err
		}
		users = append(users, &data)
	}

	if err = rows.Err(); err != nil {
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
	err := p.db.QueryRow(ctx, hasUserByFieldQuery, "id", id).Scan(&exists)
	if err != nil {
		p.log.Error("Error database query", "stack_trace", stackTrace.Get(ctx), "Cause", err.Error())
		return false, err
	}
	return exists, nil
}

func (p *postgresRepository) HasUserByLogin(ctx context.Context, login string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: HasUserByLogin")
	defer stackTrace.Done(ctx)
	var exists bool
	err := p.db.QueryRow(ctx, hasUserByFieldQuery, "login", login).Scan(&exists)
	if err != nil {
		p.log.Error("Error database query", "stack_trace", stackTrace.Get(ctx), "Cause", err.Error())
		return false, err
	}
	return exists, nil
}

func (p *postgresRepository) HasUserByEmail(ctx context.Context, email string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: HasUserByLogin")
	defer stackTrace.Done(ctx)
	var exists bool
	err := p.db.QueryRow(ctx, hasUserByFieldQuery, "email", email).Scan(&exists)
	if err != nil {
		p.log.Error("Error database query", "stack_trace", stackTrace.Get(ctx), "Cause", err.Error())
		return false, err
	}
	return exists, nil
}

func (p *postgresRepository) Create(ctx context.Context, data *model.User) (*model.User, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: Create")
	defer stackTrace.Done(ctx)

	var createdUser model.User
	row := p.db.QueryRow(ctx, createUserQuery, data.Id, data.Login, data.Email, data.Password, data.Role, data.IsBanned, data.BanReason, data.UpdatedAt, data.CreatedAt)
	err := row.Scan(&createdUser.Id, &createdUser.Login, &createdUser.Email, &createdUser.Password, &createdUser.Role, &createdUser.IsBanned, &createdUser.BanReason, &createdUser.UpdatedAt, &createdUser.CreatedAt)
	if err != nil {
		p.log.Error("error create new user", "stackTrace:", stackTrace.Get(ctx), "Cause", err.Error())
		return nil, err
	}
	return &createdUser, nil
}

func NewUserRepository(logger *slog.Logger, db *pgx.Conn) repository.UserRepository {
	return &postgresRepository{db: db, log: logger}
}

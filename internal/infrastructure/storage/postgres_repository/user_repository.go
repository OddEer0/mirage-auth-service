package postgresRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
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
	getUserByLoginQuery = `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM users
		WHERE login = $1;
	`
	hasUserByLoginQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE login = $1);
	`
	hasUserByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1);
	`
	hasUserByEmailQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE email = $1);
	`
	deleteUserById = `
		DELETE FROM users WHERE id = $1;
	`
	updateUserById = `
		UPDATE users SET login = $2, email = $3, password = $4, role = $5, isBanned = $6, banReason = $7, updatedAt = $8
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	updateUserRoleById = `
		UPDATE users SET role = $2
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	updateUserPasswordById = `
		UPDATE users SET password = $2
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	updateUserBanById = `
		UPDATE users SET isBanned = $2, banReason = $3
		WHERE id = $1
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	checkUserRoleQuery = `
		SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND role = $2);
	`
)

type postgresRepository struct {
	log *slog.Logger
	db  *pgx.Conn
}

func (p *postgresRepository) CheckUserRole(ctx context.Context, id, role string) (bool, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: CheckUserRole")
	defer stackTrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, checkUserRoleQuery, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return exists, domain.NewErr(domain.ErrInternalCode, domain.ErrNotFoundMessage)
	}
	return exists, nil
}

func (p *postgresRepository) GetByLogin(ctx context.Context, login string) (*model.User, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: GetByLogin")
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, getUserByLoginQuery, login)
	var user model.User
	err := row.Scan(&user.Id, &user.Login, &user.Email, &user.Password, &user.Role, &user.IsBanned, &user.BanReason, &user.UpdatedAt, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			p.log.ErrorContext(ctx, "not found user by id", slog.Any("cause", err), slog.String("login", login))
			return nil, ErrUserNotFound
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("login", login))
		return nil, ErrInternal
	}
	return &user, nil
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
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: GetByQuery")
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

func (p *postgresRepository) Delete(ctx context.Context, id string) error {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: Delete")
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

	_, err = p.db.Exec(ctx, deleteUserById, id)
	if err != nil {
		p.log.ErrorContext(ctx, "Delete query error", slog.Any("cause", err), slog.String("id", id))
		return ErrInternal
	}

	return nil
}

func (p *postgresRepository) UpdateById(ctx context.Context, user *model.User) (*model.User, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: UpdateById")
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, updateUserById, user.Id, user.Login, user.Email, user.Password, user.Role, user.IsBanned, user.BanReason, user.UpdatedAt)
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

func (p *postgresRepository) UpdateRoleById(ctx context.Context, id string, role string) (*model.User, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: UpdateRoleById")
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, updateUserRoleById, id, role)
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

func (p *postgresRepository) UpdatePasswordById(ctx context.Context, id string, password string) (*model.User, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: UpdatePasswordById")
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, updateUserPasswordById, id, password)
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

func (p *postgresRepository) BanUserById(ctx context.Context, id string, banReason string) (*model.User, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: BanUserById")
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, updateUserBanById, id, true, banReason)
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

func (p *postgresRepository) UnbanUserById(ctx context.Context, id string) (*model.User, error) {
	stackTrace.Add(ctx, "package: userRepository, type: postgresRepository, method: UnbanUserById")
	defer stackTrace.Done(ctx)

	row := p.db.QueryRow(ctx, updateUserBanById, id, false, nil)
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

func (p *postgresRepository) HasUserById(ctx context.Context, id string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: postgresRepository, method: HasUserByLogin")
	defer stackTrace.Done(ctx)

	var exists bool
	err := p.db.QueryRow(ctx, hasUserByIdQuery, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("id", id))
		return exists, domain.NewErr(domain.ErrInternalCode, domain.ErrNotFoundMessage)
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
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("login", login))
		return exists, ErrInternal
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
			return exists, nil
		}
		p.log.ErrorContext(ctx, "Error database query", slog.Any("cause", err), slog.String("email", email))
		return exists, ErrInternal
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
		return nil, ErrInternal
	}
	return &createdUser, nil
}

func NewUserRepository(logger *slog.Logger, db *pgx.Conn) repository.UserRepository {
	return &postgresRepository{db: db, log: logger}
}

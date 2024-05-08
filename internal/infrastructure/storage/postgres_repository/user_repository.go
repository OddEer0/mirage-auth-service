package postgresRepository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	"github.com/jackc/pgx/v5"
)

const (
	createUserQuery = ` 
		INSERT INTO users (id, login, email, password, role, isBanned, banReason, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, login, email, password, role, isBanned, banReason, updatedAt, createdAt;
	`
	selectUsers = `
		SELECT id, login, email, password, role, isBanned, banReason, updatedAt, createdAt FROM Users
	`
)

type postgresRepository struct {
	db *pgx.Conn
}

func (p *postgresRepository) GetById(ctx context.Context, id string) (*model.User, error) {
	//TODO implement me
	panic("implement me")
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

	return users, 0, nil
}

func (p *postgresRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
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
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) HasUserByLogin(ctx context.Context, id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) HasUserByEmail(ctx context.Context, id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (p *postgresRepository) Create(ctx context.Context, data *model.User) (*model.User, error) {
	var createdUser model.User
	row := p.db.QueryRow(ctx, createUserQuery, data.Id, data.Login, data.Email, data.Password, data.Role, data.IsBanned, data.BanReason, data.UpdatedAt, data.CreatedAt)
	err := row.Scan(&createdUser.Id, &createdUser.Login, &createdUser.Email, &createdUser.Password, &createdUser.Role, &createdUser.IsBanned, &createdUser.BanReason, &createdUser.UpdatedAt, &createdUser.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &createdUser, nil
}

func NewUserRepository(db *pgx.Conn) repository.UserRepository {
	return &postgresRepository{db: db}
}

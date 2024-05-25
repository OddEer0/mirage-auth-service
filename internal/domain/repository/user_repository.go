package repository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (*model.User, error)
	GetById(ctx context.Context, id string) (*model.User, error)
	GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*model.User, uint, error)
	Delete(ctx context.Context, id string) error
	UpdateById(ctx context.Context, user *model.User) (*model.User, error)
	UpdateRoleById(ctx context.Context, id string, role string) (*model.User, error)
	UpdatePasswordById(ctx context.Context, id string, password string) (*model.User, error)
	BanUserById(ctx context.Context, id string, banReason string) (*model.User, error)
	UnbanUserById(ctx context.Context, id string) (*model.User, error)
	HasUserById(ctx context.Context, id string) (bool, error)
	HasUserByLogin(ctx context.Context, login string) (bool, error)
	HasUserByEmail(ctx context.Context, email string) (bool, error)
}

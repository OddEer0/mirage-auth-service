package userUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	userService "github.com/OddEer0/mirage-auth-service/internal/app/services/user_service"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
)

type (
	UseCase interface {
		GetById(ctx context.Context, id string) (*appDto.PureUser, error)
		GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, uint, error)
		DeleteById(ctx context.Context, id string) error
		UpdateUserRole(ctx context.Context, id, role string) (*appDto.PureUser, error)
		CheckUserRole(ctx context.Context, userId, role string) (bool, error)
	}

	useCase struct {
		userService    userService.Service
		userRepository repository.UserRepository
	}
)

func New() UseCase {
	return &useCase{}
}

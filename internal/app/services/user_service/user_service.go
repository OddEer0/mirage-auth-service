package userService

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	"log/slog"
)

type (
	Service interface {
		GetById(ctx context.Context, id string) (*appDto.PureUser, error)
		GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, error)
		DeleteById(ctx context.Context, id string) error
		Create(ctx context.Context, data *appDto.RegistrationData) (*model.User, error)
	}

	service struct {
		log            *slog.Logger
		userRepository repository.UserRepository
	}
)

func (s *service) GetById(ctx context.Context, id string) (*appDto.PureUser, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteById(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func New(logger *slog.Logger, userRepository repository.UserRepository) Service {
	return &service{
		log:            logger,
		userRepository: userRepository,
	}
}

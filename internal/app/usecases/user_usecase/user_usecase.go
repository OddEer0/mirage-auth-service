package userUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
)

type (
	UseCase interface {
		GetById(ctx context.Context, id string) (*appDto.PureUser, error)
		GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, error)
		DeleteById(ctx context.Context, id string) error
	}

	useCase struct {
	}
)

func New() UseCase {
	return &useCase{}
}

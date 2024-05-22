package userUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
)

func (u *useCase) GetById(ctx context.Context, id string) (*appDto.PureUser, error) {
	//TODO implement me
	panic("implement me")
}

func (u *useCase) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, error) {
	//TODO implement me
	panic("implement me")
}

package userUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
)

func (u *useCase) GetById(ctx context.Context, id string) (*appDto.PureUser, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: GetById")
	defer stackTrace.Done(ctx)

	return u.userService.GetById(ctx, id)
}

func (u *useCase) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: GetByQuery")
	defer stackTrace.Done(ctx)

	return u.userService.GetByQuery(ctx, query)
}

package userUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	appMapper "github.com/OddEer0/mirage-auth-service/internal/app/app_mapper"
	domainQuery "github.com/OddEer0/mirage-auth-service/internal/domain/repository/domain_query"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
)

func (u *useCase) GetById(ctx context.Context, id string) (*appDto.PureUser, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: GetById")
	defer stackTrace.Done(ctx)

	user, err := u.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	mapper := appMapper.UserMapper{}
	return mapper.ToPureUser(user), nil
}

func (u *useCase) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, uint, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: GetByQuery")
	defer stackTrace.Done(ctx)

	users, pageCount, err := u.userRepository.GetByQuery(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	mappedUser := make([]*appDto.PureUser, 0, len(users))
	mapper := appMapper.UserMapper{}
	for _, user := range users {
		mappedUser = append(mappedUser, mapper.ToPureUser(user))
	}

	return mappedUser, pageCount, nil
}

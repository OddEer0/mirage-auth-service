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

	user, err := u.userRepository.GetById(ctx, id)
	if err != nil {

	}
	return &appDto.PureUser{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		IsBanned:  user.IsBanned,
		BanReason: user.BanReason,
		Role:      user.Role,
	}, nil
}

func (u *useCase) GetByQuery(ctx context.Context, query *domainQuery.UserQueryRequest) ([]*appDto.PureUser, uint, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: GetByQuery")
	defer stackTrace.Done(ctx)

	users, pageCount, err := u.userRepository.GetByQuery(ctx, query)
	if err != nil {
		return nil, 0, err
	}

	mappedUser := make([]*appDto.PureUser, 0, len(users))
	for _, user := range users {
		mappedUser = append(mappedUser, &appDto.PureUser{
			Id:        user.Id,
			Login:     user.Login,
			Email:     user.Email,
			IsBanned:  user.IsBanned,
			BanReason: user.BanReason,
			Role:      user.Role,
		})
	}

	return mappedUser, pageCount, nil
}

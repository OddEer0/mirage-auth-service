package userUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	appMapper "github.com/OddEer0/mirage-auth-service/internal/app/app_mapper"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
)

func (u *useCase) UpdateUserRole(ctx context.Context, id, role string) (*appDto.PureUser, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: UpdateUserRole")
	defer stackTrace.Done(ctx)

	user, err := u.userRepository.UpdateRoleById(ctx, id, role)
	if err != nil {
		return nil, err
	}

	mapper := appMapper.UserMapper{}
	return mapper.ToPureUser(user), nil
}

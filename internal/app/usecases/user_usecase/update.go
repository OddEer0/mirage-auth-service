package userUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
)

func (u *useCase) UpdateUserRole(ctx context.Context, id, role string) (*appDto.PureUser, error) {
	stackTrace.Add(ctx, "package: userUseCase, type: useCase, method: UpdateUserRole")
	defer stackTrace.Done(ctx)

	user, err := u.userRepository.UpdateRoleById(ctx, id, role)
	if err != nil {
		return nil, err
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

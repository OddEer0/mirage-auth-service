package authUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	tokenService "github.com/OddEer0/mirage-auth-service/internal/app/services/token_service"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
)

func (u *useCase) Registration(ctx context.Context, data *appDto.RegistrationData) (*AuthResult, error) {
	stackTrace.Add(ctx, "package: authUseCase, type: useCase, method: Registration")
	defer stackTrace.Done(ctx)

	user, err := u.userService.Create(ctx, data)
	if err != nil {
		return nil, err
	}

	tokens, err := u.tokenService.Generate(ctx, tokenService.JwtUserData{
		Id:   user.Id,
		Role: user.Role,
	})
	if err != nil {
		return nil, err
	}

	return &AuthResult{User: &appDto.PureUser{
		Id:        user.Id,
		Login:     user.Login,
		Email:     user.Email,
		IsBanned:  user.IsBanned,
		BanReason: user.BanReason,
		Role:      user.Role,
	}, Tokens: tokens}, nil
}

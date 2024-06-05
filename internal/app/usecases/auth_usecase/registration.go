package authUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	appMapper "github.com/OddEer0/mirage-auth-service/internal/app/app_mapper"
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
	_, err = u.tokenService.Save(ctx, appDto.SaveTokenServiceDto{Id: user.Id, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, err
	}

	mapper := appMapper.UserMapper{}
	pureUser := mapper.ToPureUser(user)
	return &AuthResult{User: pureUser, Tokens: tokens}, nil
}

package authUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	tokenService "github.com/OddEer0/mirage-auth-service/internal/app/services/token_service"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
)

const userExist = "user login exist"

func (u *useCase) Login(ctx context.Context, data *appDto.LoginData) (*AuthResult, error) {
	stackTrace.Add(ctx, "package: authUseCase, type: useCase, method: Login")
	defer stackTrace.Done(ctx)

	has, err := u.userRepository.HasUserByLogin(ctx, data.Login)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, domain.NewErr(domain.ErrConflictCode, userExist)
	}

	userDb, err := u.userRepository.GetByLogin(ctx, data.Login)
	if err != nil {
		return nil, err
	}
	tokens, err := u.tokenService.Generate(ctx, tokenService.JwtUserData{Id: userDb.Id, Role: userDb.Role})
	if err != nil {
		return nil, err
	}
	_, err = u.tokenService.Save(ctx, appDto.SaveTokenServiceDto{Id: userDb.Id, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		User: &appDto.PureUser{
			Id:        userDb.Id,
			Login:     userDb.Login,
			Email:     userDb.Email,
			IsBanned:  userDb.IsBanned,
			BanReason: userDb.BanReason,
			Role:      userDb.Role,
		},
		Tokens: tokens,
	}, nil
}

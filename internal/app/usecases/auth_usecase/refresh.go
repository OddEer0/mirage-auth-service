package authUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	appMapper "github.com/OddEer0/mirage-auth-service/internal/app/app_mapper"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
	"log/slog"
)

func (u *useCase) Refresh(ctx context.Context, refreshToken string) (*AuthResult, error) {
	stackTrace.Add(ctx, "package: authUseCase, type: useCase, method: Refresh")
	defer stackTrace.Done(ctx)

	if refreshToken == "" {
		return nil, domain.NewErr(domain.ErrUnauthorizedCode, domain.ErrUnauthorizedMessage)
	}

	jwtUserData, err := u.tokenService.ValidateRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	has, err := u.tokenService.HasByValue(ctx, refreshToken)
	if err != nil {
		return nil, err
	}
	if !has {
		u.log.ErrorContext(ctx, "refresh token not found", slog.String("refresh token", refreshToken))
		return nil, domain.NewErr(domain.ErrUnauthorizedCode, domain.ErrUnauthorizedMessage)
	}
	userDb, err := u.userRepository.GetById(ctx, jwtUserData.Id)
	if err != nil {
		return nil, err
	}

	tokens, err := u.tokenService.Generate(ctx, *jwtUserData)
	if err != nil {
		return nil, err
	}
	_, err = u.tokenService.Save(ctx, appDto.SaveTokenServiceDto{Id: userDb.Id, RefreshToken: tokens.RefreshToken})
	if err != nil {
		return nil, err
	}

	mapper := &appMapper.UserMapper{}
	pureUser := mapper.ToPureUser(userDb)
	return &AuthResult{
		User:   pureUser,
		Tokens: tokens,
	}, nil
}

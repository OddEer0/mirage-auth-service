package authUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	appMapper "github.com/OddEer0/mirage-auth-service/internal/app/app_mapper"
	tokenService "github.com/OddEer0/mirage-auth-service/internal/app/services/token_service"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"golang.org/x/crypto/bcrypt"
)

const LoginOrPasswordIncorrect = "incorrect login or password"

func (u *useCase) Login(ctx context.Context, data *appDto.LoginData) (*AuthResult, error) {
	stackTrace.Add(ctx, "package: authUseCase, type: useCase, method: Login")
	defer stackTrace.Done(ctx)

	has, err := u.userRepository.HasUserByLogin(ctx, data.Login)
	if err != nil {
		return nil, err
	}
	if !has {
		u.log.ErrorContext(ctx, "login error", "login_input_data", data)
		return nil, domain.NewErr(domain.ErrForbiddenCode, LoginOrPasswordIncorrect)
	}

	userDb, err := u.userRepository.GetByLogin(ctx, data.Login)
	if err != nil {
		return nil, err
	}
	isEqualPassword := bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(data.Password))
	if isEqualPassword != nil {
		u.log.ErrorContext(ctx, "login error", "login_input_data", data)
		return nil, domain.NewErr(domain.ErrForbiddenCode, LoginOrPasswordIncorrect)
	}
	tokens, err := u.tokenService.Generate(ctx, tokenService.JwtUserData{Id: userDb.Id, Role: userDb.Role})
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

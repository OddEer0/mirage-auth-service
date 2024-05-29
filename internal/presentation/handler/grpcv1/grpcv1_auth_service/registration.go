package grpcv1AuthService

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	errorHandler "github.com/OddEer0/mirage-auth-service/internal/presentation/error_handler"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) Registration(ctx context.Context, data *authv1.RegistrationRequest) (*authv1.AuthResponse, error) {
	authRes, err := a.authUseCase.Registration(ctx, &appDto.RegistrationData{
		Login:    data.Login,
		Password: data.Password,
		Email:    data.Email,
	})
	if err != nil {
		return nil, errorHandler.Catch(err)
	}
	banReason := ""
	if authRes.User.BanReason != nil {
		banReason = *authRes.User.BanReason
	}
	return &authv1.AuthResponse{
		User:   &authv1.ResponseUser{Id: authRes.User.Id, Login: authRes.User.Login, Email: authRes.User.Email, Role: authRes.User.Role, IsBanned: authRes.User.IsBanned, BanReason: banReason},
		Tokens: &authv1.JwtTokens{AccessToken: authRes.Tokens.AccessToken, RefreshToken: authRes.Tokens.RefreshToken},
	}, nil
}

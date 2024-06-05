package grpcMapper

import (
	authUseCase "github.com/OddEer0/mirage-auth-service/internal/app/usecases/auth_usecase"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

type AuthMapper struct {
}

func (a *AuthMapper) AuthResultToAuthResponseV1(authResult *authUseCase.AuthResult) *authv1.AuthResponse {
	banReason := ""
	if authResult.User.BanReason != nil {
		banReason = *authResult.User.BanReason
	}
	return &authv1.AuthResponse{
		User:   &authv1.ResponseUser{Id: authResult.User.Id, Login: authResult.User.Login, Email: authResult.User.Email, Role: authResult.User.Role, IsBanned: authResult.User.IsBanned, BanReason: banReason},
		Tokens: &authv1.JwtTokens{AccessToken: authResult.Tokens.AccessToken, RefreshToken: authResult.Tokens.RefreshToken},
	}
}

package grpcv1

import (
	"context"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) Registration(ctx context.Context, request *authv1.RegistrationRequest) (*authv1.AuthResponse, error) {
	return &authv1.AuthResponse{
		User: &authv1.ResponseUser{
			Id:    "kekw",
			Login: "Lol",
			Email: "sosipisyu@gmail.com",
			Roles: []string{"ADMIN", "MILFHUNTER"},
		},
		Tokens: &authv1.JwtTokens{
			AccessToken:  "invalid",
			RefreshToken: "real-invalid",
		},
	}, nil
}

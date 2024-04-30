package grpcv1

import (
	"context"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.AuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

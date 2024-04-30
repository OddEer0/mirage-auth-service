package grpcv1

import (
	"context"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) ChangePasswordWithAuth(ctx context.Context, auth *authv1.ChangePasswordRequestWithAuth) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AuthServiceServer) ChangePasswordWithoutAuth(ctx context.Context, auth *authv1.ChangePasswordRequestWithoutAuth) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}

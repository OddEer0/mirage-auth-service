package grpcv1

import (
	"context"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) Logout(ctx context.Context, token *authv1.RefreshToken) (*authv1.Empty, error) {
	//TODO implement me
	panic("implement me")
}

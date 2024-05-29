package grpcv1AuthService

import (
	"context"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
)

func (a *AuthServiceServer) Logout(ctx context.Context, token *authv1.RefreshToken) (*authv1.Empty, error) {
	err := a.authUseCase.Logout(ctx, token.RefreshToken)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}
	return nil, nil
}

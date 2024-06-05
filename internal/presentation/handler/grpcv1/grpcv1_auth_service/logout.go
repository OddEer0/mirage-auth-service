package grpcv1AuthService

import (
	"context"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (a *AuthServiceServer) Logout(ctx context.Context, token *authv1.RefreshToken) (*authv1.Empty, error) {
	err := a.authUseCase.Logout(ctx, token.RefreshToken)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}
	return nil, nil
}

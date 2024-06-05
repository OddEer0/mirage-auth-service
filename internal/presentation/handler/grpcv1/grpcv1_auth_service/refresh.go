package grpcv1AuthService

import (
	"context"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	grpcMapper "github.com/OddEer0/mirage-auth-service/internal/presentation/mapper/grpc_mapper"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (a *AuthServiceServer) Refresh(ctx context.Context, token *authv1.RefreshToken) (*authv1.AuthResponse, error) {
	authRes, err := a.authUseCase.Refresh(ctx, token.RefreshToken)
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}
	mapper := grpcMapper.AuthMapper{}
	return mapper.AuthResultToAuthResponseV1(authRes), nil
}

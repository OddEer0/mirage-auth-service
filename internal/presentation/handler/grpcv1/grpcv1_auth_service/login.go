package grpcv1AuthService

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	errorgrpc "github.com/OddEer0/mirage-auth-service/internal/presentation/errors/error_grpc"
	grpcMapper "github.com/OddEer0/mirage-auth-service/internal/presentation/mapper/grpc_mapper"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

func (a *AuthServiceServer) Login(ctx context.Context, data *authv1.LoginRequest) (*authv1.AuthResponse, error) {
	authRes, err := a.authUseCase.Login(ctx, &appDto.LoginData{
		Login:    data.Login,
		Password: data.Password,
	})
	if err != nil {
		return nil, errorgrpc.Catch(err)
	}
	mapper := grpcMapper.AuthMapper{}
	return mapper.AuthResultToAuthResponseV1(authRes), nil
}

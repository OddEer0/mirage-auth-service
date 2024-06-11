package grpcv1AuthService

import (
	authUseCase "github.com/OddEer0/mirage-auth-service/internal/app/usecases/auth_usecase"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
)

type (
	Dependencies struct {
		UserRepository repository.UserRepository
		AuthUseCase    authUseCase.UseCase
		Log            domain.Logger
	}

	AuthServiceServer struct {
		authv1.UnimplementedAuthServiceServer
		authUseCase    authUseCase.UseCase
		userRepository repository.UserRepository
		log            domain.Logger
	}
)

func New(dependencies *Dependencies) authv1.AuthServiceServer {
	return &AuthServiceServer{authUseCase: dependencies.AuthUseCase, userRepository: dependencies.UserRepository, log: dependencies.Log}
}

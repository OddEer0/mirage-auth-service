package grpcv1UserService

import (
	userUseCase "github.com/OddEer0/mirage-auth-service/internal/app/usecases/user_usecase"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	authv1 "github.com/OddEer0/mirage-src/protogen/mirage_auth_service"
	"log/slog"
)

type (
	Dependencies struct {
		UserRepository repository.UserRepository
		Log            *slog.Logger
		UserUseCase    userUseCase.UseCase
	}

	UserServiceServer struct {
		authv1.UnimplementedUserServiceServer
		userRepository repository.UserRepository
		userUseCase    userUseCase.UseCase
		log            *slog.Logger
	}
)

func New(dependencies *Dependencies) authv1.UserServiceServer {
	return &UserServiceServer{userUseCase: dependencies.UserUseCase, log: dependencies.Log, userRepository: dependencies.UserRepository}
}

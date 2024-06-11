package interactor

import (
	tokenService "github.com/OddEer0/mirage-auth-service/internal/app/services/token_service"
	userService "github.com/OddEer0/mirage-auth-service/internal/app/services/user_service"
	authUseCase "github.com/OddEer0/mirage-auth-service/internal/app/usecases/auth_usecase"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
)

type Dependencies struct {
	UserRepository         repository.UserRepository
	JwtTokenRepository     repository.JwtTokenRepository
	UserActivateRepository repository.UserActivateRepository
	UserService            userService.Service
	TokenService           tokenService.Service
	AuthUseCase            authUseCase.UseCase
}

func New(cfg *config.Config, log domain.Logger, db postgres.Query) *Dependencies {
	// postgres Repository initialize
	pgUserRepo := postgresRepository.NewUserRepository(log, db)
	pgJwtTokenRepo := postgresRepository.NewTokenRepository(log, db)
	pgUserActivateRepo := postgresRepository.NewUserActivateRepository(log, db)

	// app services initialize
	userServ := userService.New(log, pgUserRepo)
	tokenServ := tokenService.New(log, cfg, pgJwtTokenRepo)

	// app use case initialize
	authUCase := authUseCase.New(log, userServ, pgUserRepo, tokenServ)

	return &Dependencies{
		UserRepository:         pgUserRepo,
		JwtTokenRepository:     pgJwtTokenRepo,
		UserActivateRepository: pgUserActivateRepo,
		UserService:            userServ,
		TokenService:           tokenServ,
		AuthUseCase:            authUCase,
	}
}

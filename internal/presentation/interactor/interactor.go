package interactor

import (
	tokenService "github.com/OddEer0/mirage-auth-service/internal/app/services/token_service"
	userService "github.com/OddEer0/mirage-auth-service/internal/app/services/user_service"
	authUseCase "github.com/OddEer0/mirage-auth-service/internal/app/usecases/auth_usecase"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	postgresRepository "github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres_repository"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type Dependencies struct {
	UserRepository         repository.UserRepository
	JwtTokenRepository     repository.JwtTokenRepository
	UserActivateRepository repository.UserActivateRepository
	UserService            userService.Service
	TokenService           tokenService.Service
	AuthUseCase            authUseCase.UseCase
}

func New(cfg *config.Config, log *slog.Logger, db *pgx.Conn) *Dependencies {
	// postgres Repository initialize
	pgUserRepo := postgresRepository.NewUserRepository(log, db)
	pgJwtTokenRepo := postgresRepository.NewTokenRepository(log, db)
	pgUserActivateRepo := postgresRepository.NewUserActivateRepository()

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

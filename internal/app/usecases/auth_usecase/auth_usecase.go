package authUseCase

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	tokenService "github.com/OddEer0/mirage-auth-service/internal/app/services/token_service"
	userService "github.com/OddEer0/mirage-auth-service/internal/app/services/user_service"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
)

type (
	AuthResult struct {
		User   *appDto.PureUser
		Tokens *tokenService.JwtTokens
	}

	UseCase interface {
		Registration(ctx context.Context, data *appDto.RegistrationData) (*AuthResult, error)
		Login(ctx context.Context, data *appDto.LoginData) (*AuthResult, error)
		Logout(ctx context.Context, refreshToken string) error
		Refresh(ctx context.Context, refreshToken string) (*AuthResult, error)
	}

	useCase struct {
		log            domain.Logger
		userService    userService.Service
		userRepository repository.UserRepository
		tokenService   tokenService.Service
	}
)

func New(logger domain.Logger, userServ userService.Service, userRepo repository.UserRepository, tokenServ tokenService.Service) UseCase {
	return &useCase{
		log:            logger,
		userService:    userServ,
		userRepository: userRepo,
		tokenService:   tokenServ,
	}
}

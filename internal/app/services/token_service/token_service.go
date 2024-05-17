package tokenService

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/golang-jwt/jwt"
	"log/slog"
)

type (
	JwtTokens struct {
		AccessToken, RefreshToken string
	}

	JwtUserData struct {
		Id   string
		Role string
	}

	CustomClaims struct {
		JwtUserData `json:"jwtUserData"`
		jwt.StandardClaims
	}

	Service interface {
		HasByValue(ctx context.Context, refreshToken string) (bool, error)
		Generate(data JwtUserData) (*JwtTokens, error)
		ValidateRefreshToken(refreshToken string) (*JwtUserData, error)
		Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*model.JwtToken, error)
		DeleteByValue(ctx context.Context, value string) error
	}

	service struct {
		log             *slog.Logger
		cfg             *config.Config
		tokenRepository repository.JwtTokenRepository
	}
)

func (s *service) HasByValue(ctx context.Context, refreshToken string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) ValidateRefreshToken(refreshToken string) (*JwtUserData, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) Save(ctx context.Context, data appDto.SaveTokenServiceDto) (*model.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (s *service) DeleteByValue(ctx context.Context, value string) error {
	//TODO implement me
	panic("implement me")
}

func New(logger *slog.Logger, cfg *config.Config, tokenRepo repository.JwtTokenRepository) Service {
	return &service{
		log:             logger,
		cfg:             cfg,
		tokenRepository: tokenRepo,
	}
}

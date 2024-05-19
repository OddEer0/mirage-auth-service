package tokenService

import (
	"context"
	appError "github.com/OddEer0/mirage-auth-service/internal/app/app_error"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"github.com/golang-jwt/jwt"
	"log/slog"
	"time"
)

func (s *service) Generate(ctx context.Context, data JwtUserData) (*JwtTokens, error) {
	stackTrace.Add(ctx, "package: tokenService, type: service, method: Generate")
	defer stackTrace.Done(ctx)

	cfg := s.cfg
	accessDuration, err := time.ParseDuration(cfg.Secret.AccessTokenTime)
	if err != nil {
		s.log.ErrorContext(ctx, "parse access token duration from cfg error", slog.Any("cause", err))
		return nil, appError.Internal
	}
	refreshDuration, err := time.ParseDuration(cfg.Secret.RefreshTokenTime)
	if err != nil {
		s.log.ErrorContext(ctx, "parse refresh token duration from cfg error", slog.Any("cause", err))
		return nil, appError.Internal
	}
	accessClaims := CustomClaims{
		JwtUserData:    data,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(accessDuration).Unix()},
	}
	refreshClaims := CustomClaims{
		JwtUserData:    data,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(refreshDuration).Unix()},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	accessTokenString, err := accessToken.SignedString([]byte(cfg.Secret.ApiKey))
	if err != nil {
		s.log.ErrorContext(ctx, "access token signed token string error", slog.Any("cause", err))
		return nil, appError.Internal
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.Secret.ApiKey))
	if err != nil {
		s.log.ErrorContext(ctx, "refresh token signed token string error", slog.Any("cause", err))
		return nil, appError.Internal
	}
	return &JwtTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

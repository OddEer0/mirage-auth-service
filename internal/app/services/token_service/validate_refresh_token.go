package tokenService

import (
	"context"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
	"github.com/golang-jwt/jwt"
	"log/slog"
)

func (s *service) ValidateRefreshToken(ctx context.Context, refreshToken string) (*JwtUserData, error) {
	stackTrace.Add(ctx, "package: tokenService, type: service, method: ValidateRefreshToken")
	defer stackTrace.Done(ctx)

	cfg := s.cfg
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret.ApiKey), nil
	})
	if err != nil {
		var jwtErr *jwt.ValidationError
		if errors.As(err, &jwtErr) {
			return nil, jwtErrHandle(ctx, jwtErr, s.log)
		}
		s.log.ErrorContext(ctx, "validate jwt token error", slog.Any("cause", err), slog.String("token", refreshToken))
		return nil, domain.NewErr(domain.ErrInternalCode, domain.ErrInternalMessage)
	}
	if !token.Valid {
		s.log.ErrorContext(ctx, "invalid token", slog.Any("cause", err), slog.String("token", refreshToken))
		return nil, domain.NewErr(domain.ErrUnauthorizedCode, domain.ErrUnauthorizedMessage)
	}
	claim := token.Claims.(*CustomClaims)
	return &claim.JwtUserData, nil
}

func jwtErrHandle(ctx context.Context, jwtErr *jwt.ValidationError, log *slog.Logger) error {
	if jwtErr.Errors&jwt.ValidationErrorMalformed != 0 {
		log.ErrorContext(ctx, "uncorrected jwt token")
	} else if jwtErr.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
		log.ErrorContext(ctx, "token does not work or time")
	} else {
		log.ErrorContext(ctx, "token validate error")
	}
	return domain.NewErr(domain.ErrUnauthorizedCode, domain.ErrUnauthorizedMessage)
}

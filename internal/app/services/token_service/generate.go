package tokenService

import (
	appError "github.com/OddEer0/mirage-auth-service/internal/app/app_error"
	"github.com/golang-jwt/jwt"
	"time"
)

func (s *service) Generate(data JwtUserData) (*JwtTokens, error) {
	cfg := s.cfg
	accessDuration, err := time.ParseDuration(cfg.Secret.AccessTokenTime)
	if err != nil {
		s.log.Error("Parse duration from cfg error")
		return nil, appError.Internal
	}
	refreshDuration, err := time.ParseDuration(cfg.Secret.RefreshTokenTime)
	if err != nil {
		s.log.Error("Parse duration from cfg error")
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
		s.log.Error("access token signed token string error")
		return nil, appError.Internal
	}
	refreshTokenString, err := refreshToken.SignedString([]byte(cfg.Secret.ApiKey))
	if err != nil {
		s.log.Error("refresh token signed token string error")
		return nil, appError.Internal
	}
	return &JwtTokens{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

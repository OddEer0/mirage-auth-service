package postgresRepository

import "github.com/OddEer0/mirage-auth-service/internal/domain"

var (
	ErrTokenNotFound = domain.NewErr(domain.ErrNotFoundCode, "token not found")
	ErrUserNotFound  = domain.NewErr(domain.ErrNotFoundCode, "user not found")
	ErrInternal      = domain.NewErr(domain.ErrInternalCode, "internal error")
)

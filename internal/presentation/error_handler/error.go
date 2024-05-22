package errorHandler

import (
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Catch(err error) error {
	var e *domain.Error
	if errors.As(err, &e) {
		switch e.Code {
		case domain.ErrInternalCode:
			return status.Error(codes.Internal, e.Message)
		case domain.ErrNotFoundCode:
			return status.Error(codes.NotFound, e.Message)
		case domain.ErrConflictCode:
			return status.Error(codes.AlreadyExists, e.Message)
		}
	}
	return status.Error(codes.Internal, "internal error")
}

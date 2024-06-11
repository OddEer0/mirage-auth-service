package errorgrpc

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
		case domain.ErrUnauthorizedCode:
			return status.Error(codes.Unauthenticated, e.Message)
		case domain.ErrForbiddenCode:
			return status.Error(codes.PermissionDenied, e.Message)
		case domain.ErrRequestTimeoutCode:
			return status.Error(codes.DeadlineExceeded, e.Message)
		}
	}
	return status.Error(codes.Internal, "internal error")
}

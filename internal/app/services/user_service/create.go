package userService

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	appError "github.com/OddEer0/mirage-auth-service/internal/app/app_error"
	domainError "github.com/OddEer0/mirage-auth-service/internal/domain/domain_error"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/shared/constants"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

func (s *service) Create(ctx context.Context, data *appDto.RegistrationData) (*model.User, error) {
	stackTrace.Add(ctx, "package: userService, type: service, method: Create")
	defer stackTrace.Done(ctx)

	candidate, err := s.userRepository.HasUserByLogin(ctx, data.Login)
	if err != nil {
		return nil, err
	}
	if candidate {
		s.log.ErrorContext(ctx, "has user by current login", slog.String("cause", "has candidate"))
		return nil, domainError.NotFound
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.ErrorContext(ctx, "bcrypt hash password error", slog.Any("cause", err))
		return nil, appError.Internal
	}

	newUser, err := s.userRepository.Create(ctx, &model.User{
		Id:        uuid.New().String(),
		Login:     data.Login,
		Email:     data.Email,
		Password:  string(hashPassword),
		Role:      constants.User,
		IsBanned:  false,
		BanReason: nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

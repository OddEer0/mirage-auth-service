package userService

import (
	"context"
	appDto "github.com/OddEer0/mirage-auth-service/internal/app/app_dto"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	domainConstants "github.com/OddEer0/mirage-auth-service/internal/domain/domain_constants"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	stackTrace "github.com/OddEer0/stack-trace/stack_trace"
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
		s.log.ErrorContext(ctx, "error checking user existence", slog.Any("cause", err))
		return nil, err
	}
	if candidate {
		s.log.ErrorContext(ctx, "user already exists", slog.String("login", data.Login))
		return nil, domain.NewErr(domain.ErrConflictCode, "user login exist")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.ErrorContext(ctx, "bcrypt hash password error", slog.Any("cause", err))
		return nil, domain.NewErr(domain.ErrInternalCode, "internal error")
	}

	newUser, err := s.userRepository.Create(ctx, &model.User{
		Id:        uuid.New().String(),
		Login:     data.Login,
		Email:     data.Email,
		Password:  string(hashPassword),
		Role:      domainConstants.RoleUser,
		IsBanned:  false,
		BanReason: nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		s.log.ErrorContext(ctx, "error creating user", slog.Any("cause", err))
		return nil, err
	}
	return newUser, nil
}

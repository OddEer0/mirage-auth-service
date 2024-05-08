package grpcv1

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/shared/constants"
	authv1 "github.com/OddEer0/mirage-auth-service/pkg/gen/auth_v1"
	"github.com/google/uuid"
	"time"
)

func (a *AuthServiceServer) Registration(ctx context.Context, data *authv1.RegistrationRequest) (*authv1.AuthResponse, error) {
	user := &model.User{
		Id:        uuid.New().String(),
		Login:     data.Login,
		Email:     data.Email,
		Password:  data.Password,
		Role:      constants.User,
		IsBanned:  false,
		BanReason: nil,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newUser, err := a.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &authv1.AuthResponse{
		User:   &authv1.ResponseUser{Id: newUser.Id, Login: newUser.Login, Email: newUser.Role, IsBanned: newUser.IsBanned, BanReason: ""},
		Tokens: nil,
	}, nil
}

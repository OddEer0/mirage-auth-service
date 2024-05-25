package repository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
)

type UserActivateRepository interface {
	Create(ctx context.Context, userId string) (model.UserActivate, error)
	Delete(ctx context.Context, userId string) error
	Update(ctx context.Context, activate model.UserActivate) (model.UserActivate, error)
	GetByUserId(ctx context.Context, userId string) (model.UserActivate, error)
	ActivateUser(ctx context.Context, userId string) (model.UserActivate, error)
	IsActivateById(ctx context.Context, userId string) (bool, error)
}

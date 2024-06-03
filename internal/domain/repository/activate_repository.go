package repository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
)

type UserActivateRepository interface {
	Create(ctx context.Context, userId string) (*model.UserActivate, error)
	Delete(ctx context.Context, userId string) error
	Update(ctx context.Context, activate *model.UserActivate) (*model.UserActivate, error)
	GetByUserId(ctx context.Context, userId string) (*model.UserActivate, error)
	ActivateUserById(ctx context.Context, userId string) (*model.UserActivate, error)
	ActivateUserByLink(ctx context.Context, link string) (*model.UserActivate, error)
	IsActivateById(ctx context.Context, userId string) (bool, error)
	HasById(ctx context.Context, userId string) (bool, error)
}

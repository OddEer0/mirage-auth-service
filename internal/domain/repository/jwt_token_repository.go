package repository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
)

type JwtTokenRepository interface {
	Create(ctx context.Context, id, token string) (*model.JwtToken, error)
	Delete(ctx context.Context, id string) error
	DeleteByValue(ctx context.Context, value string) error
	UpdateById(ctx context.Context, id, token string) (*model.JwtToken, error)
	Save(ctx context.Context, id, token string) (*model.JwtToken, error)
	GetById(ctx context.Context, id string) (*model.JwtToken, error)
	HasById(ctx context.Context, id string) (bool, error)
	HasByValue(ctx context.Context, token string) (bool, error)
}

package postgresRepository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type tokenRepository struct {
	log *slog.Logger
	db  *pgx.Conn
}

func (t tokenRepository) Create(ctx context.Context, id, token string) (*model.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) UpdateById(ctx context.Context, id, token string) (*model.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) Save(ctx context.Context, id, token string) (*model.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) GetById(ctx context.Context, id string) (*model.JwtToken, error) {
	//TODO implement me
	panic("implement me")
}

func (t tokenRepository) HasById(ctx context.Context, id string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewTokenRepository(logger *slog.Logger, db *pgx.Conn) repository.JwtTokenRepository {
	return &tokenRepository{db: db, log: logger}
}

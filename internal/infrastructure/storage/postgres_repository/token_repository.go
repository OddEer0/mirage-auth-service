package postgresRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

const (
	createTokenQuery = `
		INSERT INTO tokens (id, value, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4)
		RETURNING id, value, updatedAt, createdAt;
	`
	deleteTokenQuery = `
		DELETE FROM tokens WHERE id = $1
	`
	updateTokenQuery = `
		UPDATE tokens SET value = $2, updatedAt = $3 WHERE id = $1
		RETURNING id, value, updatedAt, createdAt;
	`
	hasTokenById = `
		SELECT EXISTS(SELECT 1 FROM tokens WHERE id = $1);
	`
	hasTokenByValue = `
		SELECT EXISTS(SELECT 1 FROM tokens WHERE value = $1);
	`
	getTokenById = `
		SELECT id, value, updatedAt, createdAt FROM tokens
		WHERE id = $1;
	`
)

type tokenRepository struct {
	log *slog.Logger
	db  *pgx.Conn
}

func (t tokenRepository) Create(ctx context.Context, id, token string) (*model.JwtToken, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: structRepository, method: Create")
	defer stackTrace.Done(ctx)

	newToken := &model.JwtToken{}
	row := t.db.QueryRow(ctx, createTokenQuery, id, token, time.Now(), time.Now())
	err := row.Scan(&newToken.UserId, &newToken.Value, &newToken.UpdatedAt, &newToken.CreatedAt)
	if err != nil {
		t.log.ErrorContext(ctx, "create token query error", slog.Any("cause", err), slog.String("id", id), slog.String("token", token))
		return nil, ErrInternal
	}

	return newToken, nil
}

func (t tokenRepository) Delete(ctx context.Context, id string) error {
	stackTrace.Add(ctx, "package: postgresRepository, type: structRepository, method: Delete")
	defer stackTrace.Done(ctx)

	_, err := t.db.Exec(ctx, deleteTokenQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.log.ErrorContext(ctx, "token not found", slog.Any("cause", err), slog.String("id", id))
			return ErrTokenNotFound
		}
		t.log.ErrorContext(ctx, "delete token query error", slog.Any("cause", err), slog.String("id", id))
		return ErrInternal
	}

	return nil
}

func (t tokenRepository) UpdateById(ctx context.Context, id, token string) (*model.JwtToken, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: structRepository, method: UpdateById")
	defer stackTrace.Done(ctx)

	updatedToken := &model.JwtToken{}
	row := t.db.QueryRow(ctx, updateTokenQuery, id, token, time.Now())
	err := row.Scan(&updatedToken.UserId, &updatedToken.Value, &updatedToken.UpdatedAt, &updatedToken.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.log.ErrorContext(ctx, "token by id not found", slog.Any("cause", err), slog.String("id", id), slog.String("token", token))
			return nil, ErrTokenNotFound
		}
		t.log.ErrorContext(ctx, "update token query or scan error", slog.Any("cause", err), slog.String("id", id), slog.String("token", token))
		return nil, ErrInternal
	}

	return updatedToken, nil
}

func (t tokenRepository) Save(ctx context.Context, id, token string) (*model.JwtToken, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: structRepository, method: Save")
	defer stackTrace.Done(ctx)

	has, err := t.HasById(ctx, id)
	if err != nil {
		t.log.ErrorContext(ctx, "HasById method error")
		return nil, err
	}
	if has {
		updated, err := t.UpdateById(ctx, id, token)
		if err != nil {
			t.log.ErrorContext(ctx, "UpdateById method error")
			return nil, err
		}
		return updated, nil
	}
	created, err := t.Create(ctx, id, token)
	if err != nil {
		t.log.ErrorContext(ctx, "Create method error")
		return nil, err
	}

	return created, nil
}

func (t tokenRepository) GetById(ctx context.Context, id string) (*model.JwtToken, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: structRepository, method: GetById")
	defer stackTrace.Done(ctx)

	token := &model.JwtToken{}
	err := t.db.QueryRow(ctx, getTokenById, id).Scan(&token.UserId, &token.Value, &token.UpdatedAt, &token.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.log.ErrorContext(ctx, "token not found", slog.Any("cause", err), slog.String("id", id))
			return nil, ErrTokenNotFound
		}
		t.log.ErrorContext(ctx, "GetTokenById query error", slog.Any("cause", err), slog.String("id", id))
		return nil, ErrInternal
	}

	return token, nil
}

func (t tokenRepository) HasById(ctx context.Context, id string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: structRepository, method: HasById")
	defer stackTrace.Done(ctx)

	var exists bool
	err := t.db.QueryRow(ctx, hasTokenById, id).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.log.ErrorContext(ctx, "token not found", slog.Any("cause", err), slog.String("id", id))
			return exists, ErrTokenNotFound
		}
		t.log.ErrorContext(ctx, "HasTokenById query error", slog.Any("cause", err), slog.String("id", id))
		return exists, ErrInternal
	}

	return exists, nil
}

func (t tokenRepository) HasByValue(ctx context.Context, token string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: structRepository, method: HasById")
	defer stackTrace.Done(ctx)

	var exists bool
	err := t.db.QueryRow(ctx, hasTokenByValue, token).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.log.ErrorContext(ctx, "token not found", slog.Any("cause", err), slog.String("token", token))
			return exists, ErrTokenNotFound
		}
		t.log.ErrorContext(ctx, "HasTokenById query error", slog.Any("cause", err), slog.String("token", token))
		return exists, ErrInternal
	}

	return exists, nil
}

func NewTokenRepository(logger *slog.Logger, db *pgx.Conn) repository.JwtTokenRepository {
	return &tokenRepository{db: db, log: logger}
}

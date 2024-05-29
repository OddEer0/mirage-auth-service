package postgresRepository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/OddEer0/mirage-auth-service/internal/domain/model"
	"github.com/OddEer0/mirage-auth-service/internal/domain/repository"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"time"
)

const (
	createUserActivateQuery = `
		INSERT INTO user_activate (id, isActivate, link, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, isActivate, link, updatedAt, createdAt;
	`
	deleteUserActivateQuery = `
		DELETE FROM user_activate WHERE id = $1;
	`
	getUserActivateByIdQuery = `
		SELECT id, isActivate, link, updatedAt, createdAt FROM user_activate WHERE id = $1;
	`
	hasUserActivateByIdQuery = `
		SELECT EXISTS(SELECT 1 FROM user_activate WHERE id = $1);
	`
	isUserActivateQuery = `
		SELECT EXISTS(SELECT 1 FROM user_activate WHERE id = $1 AND isActivate = true);
	`
	activateUserByIdQuery = `
		UPDATE user_activate SET isActivate = true WHERE id = $1;
	`
	activateUserByLinkQuery = `
		UPDATE user_activate SET isActivate = true WHERE link = $1;
	`
	updateUserActivateQuery = `
		UPDATE user_activate SET isActivate = $2, link = $3, updatedAt = $4 WHERE id = $1
		RETURNING id, isActivate, link, updatedAt, createdAt;
	`
)

type userActivateRepository struct {
	log *slog.Logger
	db  *pgx.Conn
}

func (u *userActivateRepository) Create(ctx context.Context, userId string) (*model.UserActivate, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: userActivateRepository, method: Create")
	defer stackTrace.Done(ctx)

	// TODO - Сделать правильный метод для генераций ссылок
	link := uuid.New().String() + ".com"

	result := &model.UserActivate{}
	err := u.db.QueryRow(ctx, createUserActivateQuery, userId, false, link, time.Now(), time.Now()).
		Scan(&result.UserId, &result.IsActivate, &result.Link, &result.UpdatedAt, &result.CreatedAt)

	if err != nil {
		u.log.ErrorContext(ctx, "", slog.Any("cause", err), slog.String("userId", userId))
		return nil, ErrInternal
	}

	return result, nil
}

func (u *userActivateRepository) Delete(ctx context.Context, userId string) error {
	stackTrace.Add(ctx, "package: postgresRepository, type: userActivateRepository, method: Delete")
	defer stackTrace.Done(ctx)

	_, err := u.db.Exec(ctx, deleteUserActivateQuery, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.log.ErrorContext(ctx, "userActivate table not found", slog.Any("cause", err), slog.String("userId", userId))
			return ErrUserActivateNotFound
		}
		u.log.ErrorContext(ctx, "delete user activate query error", slog.Any("cause", err), slog.String("userId", userId))
		return ErrInternal
	}

	return nil
}

func (u *userActivateRepository) Update(ctx context.Context, activate *model.UserActivate) (*model.UserActivate, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: userActivateRepository, method: Update")
	defer stackTrace.Done(ctx)

	updatedActivate := &model.UserActivate{}
	err := u.db.QueryRow(ctx, updateUserActivateQuery, activate.UserId, activate.IsActivate, activate.Link, activate.UpdatedAt).
		Scan(&updatedActivate.UserId, &updatedActivate.IsActivate, &updatedActivate.Link, &updatedActivate.UpdatedAt, &updatedActivate.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.log.ErrorContext(ctx, "user activate not found", slog.Any("cause", err), slog.String("id", activate.UserId))
			return nil, ErrUserActivateNotFound
		}
		u.log.ErrorContext(ctx, "user activate update query error", slog.Any("cause", err), "data", activate)
		return nil, ErrInternal
	}

	return updatedActivate, nil
}

func (u *userActivateRepository) GetByUserId(ctx context.Context, userId string) (*model.UserActivate, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: userActivateRepository, method: GetByUserId")
	defer stackTrace.Done(ctx)

	userActivate := &model.UserActivate{}
	err := u.db.QueryRow(ctx, getUserActivateByIdQuery, userId).
		Scan(&userActivate.UserId, &userActivate.IsActivate, &userActivate.Link, &userActivate.UpdatedAt, &userActivate.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.log.ErrorContext(ctx, "userActivate table not found", slog.Any("cause", err), slog.String("userId", userId))
			return nil, ErrUserActivateNotFound
		}
		u.log.ErrorContext(ctx, "get user activate by id query error", slog.Any("cause", err), slog.String("userId", userId))
		return nil, ErrInternal
	}

	return userActivate, nil
}

func (u *userActivateRepository) ActivateUser(ctx context.Context, userId string) (*model.UserActivate, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: userActivateRepository, method: ActivateUser")
	defer stackTrace.Done(ctx)

	updatedActivate := &model.UserActivate{}
	err := u.db.QueryRow(ctx, activateUserByIdQuery, userId).
		Scan(&updatedActivate.UserId, &updatedActivate.IsActivate, &updatedActivate.Link, &updatedActivate.UpdatedAt, &updatedActivate.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.log.ErrorContext(ctx, "user activate not found", slog.Any("cause", err), slog.String("id", userId))
			return nil, ErrUserActivateNotFound
		}
		u.log.ErrorContext(ctx, "activate user query error", slog.Any("cause", err), slog.String("id", userId))
		return nil, ErrInternal
	}

	return updatedActivate, nil
}

func (u *userActivateRepository) IsActivateById(ctx context.Context, userId string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: userActivateRepository, method: IsActivateById")
	defer stackTrace.Done(ctx)

	var exists bool
	err := u.db.QueryRow(ctx, isUserActivateQuery, userId).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			u.log.ErrorContext(ctx, "user activate not found", slog.Any("cause", err), slog.String("id", userId))
			return exists, ErrUserActivateNotFound
		}
		u.log.ErrorContext(ctx, "is user activate checking query error", slog.Any("cause", err), slog.String("id", userId))
		return exists, ErrInternal
	}

	return exists, nil
}

func (u *userActivateRepository) HasById(ctx context.Context, userId string) (bool, error) {
	stackTrace.Add(ctx, "package: postgresRepository, type: userActivateRepository, method: HasById")
	defer stackTrace.Done(ctx)

	var exists bool
	err := u.db.QueryRow(ctx, hasUserActivateByIdQuery, userId).Scan(&exists)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exists, nil
		}
		u.log.ErrorContext(ctx, "is user activate checking query error", slog.Any("cause", err), slog.String("id", userId))
		return exists, ErrInternal
	}

	return exists, nil
}

func NewUserActivateRepository(log *slog.Logger, db *pgx.Conn) repository.UserActivateRepository {
	return &userActivateRepository{
		log: log,
		db:  db,
	}
}

package postgres

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/OddEer0/mirage-auth-service/internal/shared/constants"
	stackTrace "github.com/OddEer0/mirage-auth-service/pkg/stack_trace"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"log/slog"
	"time"
)

func InitMigration(ctx context.Context, conn *pgx.Conn, log *slog.Logger) error {
	if _, err := conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		login VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		role VARCHAR(30) NOT NULL,
		isBanned BOOLEAN NOT NULL,
		banReason VARCHAR(255),
		updatedAt DATE NOT NULL,
		createdAt DATE NOT NULL
	)`); err != nil {
		return err
	}
	log.Info("add users table")

	if _, err := conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS tokens (
		id UUID PRIMARY KEY REFERENCES users(id),
    	value VARCHAR(255) NOT NULL
	)`); err != nil {
		return err
	}
	log.Info("add token table")

	return nil
}

func InitSuperAdmin(ctx context.Context, conn *pgx.Conn, cfg *config.Config, log *slog.Logger) error {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)"
	var exists bool
	err := conn.QueryRow(ctx, query, cfg.AdminName).Scan(&exists)
	if err != nil {
		log.Error("Error database query", "stack_trace", stackTrace.Get(ctx), "Cause", err.Error())
		return err
	}

	if exists {
		return nil
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(cfg.AdminPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	row := conn.QueryRow(ctx, `
		INSERT INTO users (id, login, email, password, role, isBanned, banReason, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, login, email;	
	`, uuid.New().String(), cfg.AdminName, "kostyl@gmail.com", hashPassword, constants.SuperAdmin, false, nil, time.Now(), time.Now())
	if err != nil {
		return err
	}
	var login, email, id string
	err = row.Scan(&id, &login, &email)
	if err != nil {
		return err
	}
	log.Info("Add super admin", "id", id, "login", login, "email", email)
	return nil
}

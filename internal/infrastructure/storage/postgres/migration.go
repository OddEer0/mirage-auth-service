package postgres

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	domainConstants "github.com/OddEer0/mirage-auth-service/internal/domain/domain_constants"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func InitMigration(ctx context.Context, conn Query, log domain.Logger) error {
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

	if _, err := conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS tokens (
		id UUID PRIMARY KEY REFERENCES users(id),
    	value VARCHAR(255) NOT NULL,
    	updatedAt DATE NOT NULL,
		createdAt DATE NOT NULL
	)`); err != nil {
		return err
	}

	if _, err := conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS user_activate (
		id UUID PRIMARY KEY REFERENCES users(id),
    	isActivate BOOLEAN NOT NULL,
    	link VARCHAR(255) NOT NULL,
    	updatedAt DATE NOT NULL,
		createdAt DATE NOT NULL
	)`); err != nil {
		return err
	}

	log.Info("success init migration")

	return nil
}

func InitSuperAdmin(ctx context.Context, conn Query, cfg *config.Config, log domain.Logger) error {
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE login = $1)"
	var exists bool
	err := conn.QueryRow(ctx, query, cfg.SuperAdmin.Login).Scan(&exists)
	if err != nil {
		log.Error("init super admin has check query error", "Cause", err.Error())
		return err
	}

	if exists {
		log.Info("super admin exists")
		return nil
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(cfg.SuperAdmin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	row := conn.QueryRow(ctx, `
		INSERT INTO users (id, login, email, password, role, isBanned, banReason, updatedAt, createdAt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, login, email;	
	`, uuid.New().String(), cfg.SuperAdmin.Login, "kostyl@gmail.com", hashPassword, domainConstants.RoleSuperAdmin, false, nil, time.Now(), time.Now())
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

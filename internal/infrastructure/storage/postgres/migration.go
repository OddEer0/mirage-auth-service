package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func InitMigration(ctx context.Context, conn *pgx.Conn) error {
	if _, err := conn.Exec(ctx, `CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		login VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
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
    	value VARCHAR(255) NOT NULL
	)`); err != nil {
		return err
	}

	return nil
}

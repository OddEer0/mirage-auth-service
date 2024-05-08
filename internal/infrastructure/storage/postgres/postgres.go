package postgres

import (
	"context"
	"fmt"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/jackc/pgx/v5"
)

func Connect(cfg *config.Config) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DbName)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	err = InitMigration(ctx, conn)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

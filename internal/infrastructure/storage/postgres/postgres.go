package postgres

import (
	"context"
	"fmt"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

func Connect(cfg *config.Config, log *slog.Logger) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DbName)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	log.Info("Connect Postgres")
	err = InitMigration(ctx, conn, log)
	if err != nil {
		return nil, err
	}
	log.Info("Success init postgres")
	err = InitSuperAdmin(ctx, conn, cfg, log)
	if err != nil {
		return nil, err
	}
	log.Info("Success add super admin")
	return conn, nil
}

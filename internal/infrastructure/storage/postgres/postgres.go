package postgres

import (
	"context"
	"fmt"
	"github.com/OddEer0/mirage-auth-service/internal/domain"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Query interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Close(ctx context.Context) error
	Prepare(ctx context.Context, name string, sql string) (sd *pgconn.StatementDescription, err error)
	Deallocate(ctx context.Context, name string) error
	DeallocateAll(ctx context.Context) error
	WaitForNotification(ctx context.Context) (*pgconn.Notification, error)
	IsClosed() bool
	Ping(ctx context.Context) error
	PgConn() *pgconn.PgConn
	TypeMap() *pgtype.Map
	Config() *pgx.ConnConfig
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	SendBatch(ctx context.Context, b *pgx.Batch) (br pgx.BatchResults)
	LoadType(ctx context.Context, typeName string) (*pgtype.Type, error)
}

func Connect(cfg *config.Config, log domain.Logger) (Query, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.DbName)
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		return nil, err
	}
	log.Info("connect Postgres")
	err = InitMigration(ctx, conn, log)
	if err != nil {
		return nil, err
	}
	log.Info("success init postgres")
	err = InitSuperAdmin(ctx, conn, cfg, log)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

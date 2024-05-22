package postgresRepository

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func getTableCount(ctx context.Context, conn *pgx.Conn, tableName string) (uint, error) {
	var result uint
	err := conn.QueryRow(ctx, `
        SELECT COUNT(*)
        FROM `+tableName).Scan(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func getPageCount(ctx context.Context, conn *pgx.Conn, tableName string, limit uint) (uint, error) {
	var result uint
	err := conn.QueryRow(ctx, `
        SELECT COUNT(*)
        FROM `+tableName).Scan(&result)
	if err != nil {
		return result, err
	}

	remainder := result % limit
	result /= limit
	if remainder != 0 {
		result++
	}

	return result, nil
}

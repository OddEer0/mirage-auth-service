package postgresRepository

import (
	"context"
	"github.com/OddEer0/mirage-auth-service/internal/infrastructure/storage/postgres"
)

func GetTableCount(ctx context.Context, conn postgres.Query, tableName string) (uint, error) {
	var result uint
	err := conn.QueryRow(ctx, `
        SELECT COUNT(*)
        FROM `+tableName).Scan(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func GetPageCount(ctx context.Context, conn postgres.Query, tableName string, limit uint) (uint, error) {
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

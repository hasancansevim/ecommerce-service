package infrastructure

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func TruncateProductTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE products RESTART IDENTITY")

	if truncateResultErr != nil {
		log.Error(truncateResultErr.Error())
	} else {
		log.Info("Products table truncated successfully")
	}
}

func TruncateUserTestData(ctx context.Context, dbPool *pgxpool.Pool) {
	_, truncateResultErr := dbPool.Exec(ctx, "TRUNCATE users RESTART IDENTITY")

	if truncateResultErr != nil {
		log.Error(truncateResultErr.Error())
	} else {
		log.Info("Users table truncated successfully")
	}
}

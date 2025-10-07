package testdata

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type TestDataManager interface {
	Initialize(ctx context.Context, dbPool *pgxpool.Pool) error
	Cleanup(ctx context.Context, dbPool *pgxpool.Pool) error
	GetName() string
}

type TestData struct {
	Name       string
	SQL        string
	CleanupSQL string
}

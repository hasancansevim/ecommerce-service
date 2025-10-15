package postgresql

import (
	"context"
	"fmt"
	"go-ecommerce-service/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetConnectionPool(ctx context.Context, cfg config.DatabaseConfig) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%d pool_max_conn_idle_time=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.Name,
		cfg.MaxConnections,
		cfg.MaxConnectionIdleTime,
	)

	connConfig, _ := pgxpool.ParseConfig(connString)

	return pgxpool.ConnectConfig(ctx, connConfig)
}

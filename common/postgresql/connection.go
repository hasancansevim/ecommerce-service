package postgresql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

func GetConnectionPool(ctx context.Context, config Config) *pgxpool.Pool {
	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable statement_cache_mode=describe pool_max_conns=%s pool_max_conn_idle_time=%s",
		config.Host,
		config.Port,
		config.Username,
		config.Password,
		config.DbName,
		config.MaxConnections,
		config.MaxConnectionIdleTime,
	)

	connConfig, parseConfigError := pgxpool.ParseConfig(connString)

	if parseConfigError != nil {
		panic(parseConfigError.Error())
	}
	connectConfig, connectConfigError := pgxpool.ConnectConfig(ctx, connConfig)
	if connectConfigError != nil {
		panic(connectConfigError.Error())
	}
	return connectConfig
}

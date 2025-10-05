package helper

import (
	"context"
	"go-ecommerce-service/persistence/common"
	"go-ecommerce-service/persistence/helper/interfaces"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type GenericScanner[T interfaces.Scannable] struct {
	dbPool   *pgxpool.Pool
	scanFunc func(pgx.Row) (T, error)
}

func NewGenericScanner[T interfaces.Scannable](dbPool *pgxpool.Pool, scanFunc func(row pgx.Row) (T, error)) *GenericScanner[T] {
	return &GenericScanner[T]{
		dbPool:   dbPool,
		scanFunc: scanFunc,
	}
}
func (gs *GenericScanner[T]) ExecuteQuery(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := gs.dbPool.Query(ctx, query, args...)
	if err != nil {
		return nil, common.WrapError("execute query", err)
	}
	return rows, nil
}

func (gs *GenericScanner[T]) ExecuteExec(ctx context.Context, query string, args ...interface{}) error {
	_, err := gs.dbPool.Exec(ctx, query, args...)
	if err != nil {
		return common.WrapError("execute exec", err)
	}
	return nil
}

func (gs *GenericScanner[T]) Scan(row pgx.Row) (T, error) {
	return gs.scanFunc(row)
}

func (gs *GenericScanner[T]) ScanAll(rows pgx.Rows) ([]T, error) {
	defer rows.Close()

	var results []T
	for rows.Next() {
		result, err := gs.Scan(rows)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (gs *GenericScanner[T]) QueryAndScan(ctx context.Context, query string, args ...interface{}) ([]T, error) {
	rows, err := gs.ExecuteQuery(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return gs.ScanAll(rows)
}

func (gs *GenericScanner[T]) QueryRowAndScan(ctx context.Context, query string, args ...interface{}) (T, error) {
	row := gs.dbPool.QueryRow(ctx, query, args...)
	return gs.Scan(row)
}

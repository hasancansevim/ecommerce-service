package testdata

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

type ProductTestDataManager struct{}

func NewProductTestDataManager() TestDataManager {
	return &ProductTestDataManager{}
}

func (p ProductTestDataManager) GetName() string {
	return "products"
}

func (p ProductTestDataManager) Initialize(ctx context.Context, dbPool *pgxpool.Pool) error {
	sql := `
		INSERT INTO products (name, price, discount, store) VALUES 
        ('Laptop', 20000.0, 10.0, 'Teknosa'),
        ('Klavye', 800.0, 0.0, 'Teknosa'),
        ('Mouse', 500.0, 10.0, 'Teknosa'),
        ('Ütü', 200.0, 0.0, 'Güzel Evim');	
	`
	result, err := dbPool.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("❌ failed to insert product test data: %w", err)
	}
	log.Info(fmt.Sprintf("✅ Product test data created: %d rows", result.RowsAffected()))
	return nil
}

func (p ProductTestDataManager) Cleanup(ctx context.Context, dbPool *pgxpool.Pool) error {
	sql := `
	DELETE FROM products WHERE name IN ('Laptop', 'Klavye', 'Mouse', 'Ütü')
	`
	result, err := dbPool.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("❌ failed to cleanup product test data: %w", err)
	}
	log.Info(fmt.Sprintf("🧹 Product test data cleaned: %d rows", result.RowsAffected()))
	return nil
}

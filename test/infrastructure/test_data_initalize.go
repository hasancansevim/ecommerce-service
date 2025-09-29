package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

var INSERT_PRODUCTS = `
INSERT INTO products (name, price, discount, store) VALUES 
	('Laptop', 20000.0, 10.0, 'Teknosa'),
	('Klavye', 800.0, 0.0, 'Teknosa'),
	('Mouse', 500.0, 10.0, 'Teknosa'),
	('Ütü', 200.0, 0.0, 'Güzel Evim');
`

func TestDataInitalize(ctx context.Context, dbPool *pgxpool.Pool) {
	insertProductsResult, insertProductsError := dbPool.Exec(ctx, INSERT_PRODUCTS)
	if insertProductsError != nil {
		log.Error(insertProductsError.Error())
	} else {
		log.Info(fmt.Sprintf("Products data created with %d rows", insertProductsResult.RowsAffected()))
	}
}

package infrastructure

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
)

func ProductTestDataInitalize(ctx context.Context, dbPool *pgxpool.Pool) {
	var INSERT_PRODUCTS = `
		INSERT INTO products (name, price, discount, store) VALUES 
			('Laptop', 20000.0, 10.0, 'Teknosa'),
			('Klavye', 800.0, 0.0, 'Teknosa'),
			('Mouse', 500.0, 10.0, 'Teknosa'),
			('Ütü', 200.0, 0.0, 'Güzel Evim');
		`
	insertProductsResult, insertProductsError := dbPool.Exec(ctx, INSERT_PRODUCTS)
	if insertProductsError != nil {
		log.Error(insertProductsError.Error())
	} else {
		log.Info(fmt.Sprintf("Products data created with %d rows", insertProductsResult.RowsAffected()))
	}
}
func UserTestDataInitialize(ctx context.Context, dbPool *pgxpool.Pool) {
	var INSERT_USERS = `
		INSERT INTO users (first_name, last_name, email, password, created_at) VALUES
			('Hasan Can', 'Sevim', 'shasancan0@gmail.com', 'hashed_password_1', CURRENT_TIMESTAMP),
			('Mustafa Murat', 'Coşkun', 'mmurat@gmail.com', 'hashed_password_2', CURRENT_TIMESTAMP);
	`
	insertUserResult, insertUserErr := dbPool.Exec(ctx, INSERT_USERS)
	if insertUserErr != nil {
		log.Error("User insert error: ", insertUserErr.Error())
	} else {
		log.Info(fmt.Sprintf("Users data created with %d rows", insertUserResult.RowsAffected()))
	}
}

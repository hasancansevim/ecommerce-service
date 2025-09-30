package main

import (
	"context"
	"go-ecommerce-service/common/app"
	"go-ecommerce-service/common/postgresql"
	"go-ecommerce-service/controller"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service"

	"github.com/labstack/echo/v4"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	e := echo.New()
	ctx := context.Background()

	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSQLConfig)
	productRepository := persistence.NewProductRepository(dbPool)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)
	productController.RegisterRoutes(e)
	e.Start("localhost:8080")
}

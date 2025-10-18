package main

import (
	"context"
	"go-ecommerce-service/common/postgresql"
	"go-ecommerce-service/config"
	"go-ecommerce-service/controller"
	"go-ecommerce-service/internal/jwt"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	jwt.Initialize(cfg.Auth.JWTSecret)

	ctx := context.Background()

	dbPool, dbPoolErr := postgresql.GetConnectionPool(ctx, cfg.Database)
	if dbPoolErr != nil {
		log.Fatal("Failed to connect to database", err)
	}
	defer dbPool.Close()

	// Dependency Injection

	productRepository := persistence.NewProductRepository(dbPool)
	userRepository := persistence.NewUserRepository(dbPool)
	cartRepository := persistence.NewCartRepository(dbPool)
	carItemRepository := persistence.NewCartItemRepository(dbPool)
	orderRepository := persistence.NewOrderRepository(dbPool)
	orderItemRepository := persistence.NewOrderItemRepository(dbPool)

	productService := service.NewProductService(productRepository)
	userService := service.NewUserService(userRepository)
	cartService := service.NewCartService(cartRepository)
	carItemService := service.NewCartItemService(carItemRepository)
	orderService := service.NewOrderService(orderRepository)
	orderItemService := service.NewOrderItemService(orderItemRepository)
	authService := service.NewAuthService(userRepository)

	productController := controller.NewProductController(productService)
	userController := controller.NewUserController(userService)
	cartController := controller.NewCartController(cartService)
	cartItemController := controller.NewCartItemController(carItemService)
	orderController := controller.NewOrderController(orderService)
	orderItemController := controller.NewOrderItemController(orderItemService)
	authController := controller.NewAuthController(authService)

	e := echo.New()

	productController.RegisterRoutes(e)
	userController.RegisterRoutes(e)
	cartController.RegisterRoutes(e)
	cartItemController.RegiesterRoutes(e)
	orderController.RegisterRoutes(e)
	orderItemController.RegisterRoutes(e)
	authController.RegisterRoutes(e)

	log.Printf("Server starting on %s", cfg.Server.Port)
	if err := e.Start("localhost:" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

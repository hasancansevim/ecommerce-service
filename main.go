package main

import (
	"context"
	"go-ecommerce-service/common/postgresql"
	"go-ecommerce-service/config"
	"go-ecommerce-service/controller"
	"go-ecommerce-service/internal/jwt"
	"go-ecommerce-service/persistence"
	customMiddleware "go-ecommerce-service/pkg/middleware"
	"go-ecommerce-service/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

func main() {

	/*
		docker-compose up -d postgres redis
		docker-compose up --build
	*/

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}
	jwt.Initialize(cfg.Auth.JWTSecret)

	ctx := context.Background()

	// Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host + ":" + cfg.Redis.Port,
	})
	if redisConnectionErr := rdb.Ping(ctx).Err(); redisConnectionErr != nil {
		log.Fatal("Failed connect to redis", redisConnectionErr)
	}

	dbPool, dbPoolErr := postgresql.GetConnectionPool(ctx, cfg.Database)
	if dbPoolErr != nil {
		log.Fatal("Failed connect to database", err)
	}
	defer dbPool.Close()

	// Dependency Injection

	productRepository := persistence.NewProductRepository(dbPool)
	userRepository := persistence.NewUserRepository(dbPool)
	cartRepository := persistence.NewCartRepository(dbPool)
	carItemRepository := persistence.NewCartItemRepository(dbPool)
	orderRepository := persistence.NewOrderRepository(dbPool)
	orderItemRepository := persistence.NewOrderItemRepository(dbPool)
	categoryRepository := persistence.NewCategoryRepository(dbPool)
	storeRepository := persistence.NewStoreRepository(dbPool)

	productService := service.NewProductService(productRepository, rdb)
	userService := service.NewUserService(userRepository)
	cartService := service.NewCartService(cartRepository)
	carItemService := service.NewCartItemService(carItemRepository)
	orderService := service.NewOrderService(orderRepository)
	orderItemService := service.NewOrderItemService(orderItemRepository)
	jwtManager := service.NewJWTService()
	authService := service.NewAuthService(userRepository, jwtManager)
	categoryService := service.NewCategoryService(categoryRepository)
	storeService := service.NewStoreService(storeRepository)

	productController := controller.NewProductController(productService)
	userController := controller.NewUserController(userService)
	cartController := controller.NewCartController(cartService)
	cartItemController := controller.NewCartItemController(carItemService)
	orderController := controller.NewOrderController(orderService)
	orderItemController := controller.NewOrderItemController(orderItemService)
	authController := controller.NewAuthController(authService)
	categoryController := controller.NewCategoryController(categoryService)
	storeController := controller.NewStoreController(storeService)

	e := echo.New()

	productController.RegisterRoutes(e)
	userController.RegisterRoutes(e)
	cartController.RegisterRoutes(e)
	cartItemController.RegiesterRoutes(e)
	orderController.RegisterRoutes(e)
	orderItemController.RegisterRoutes(e)
	authController.RegisterRoutes(e)
	categoryController.RegisterRoutes(e)
	storeController.RegisterRoutes(e)

	e.HTTPErrorHandler = customMiddleware.CustomHTTPErrorHandler

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.Logger.Fatal(e.Start(":8080"))

	log.Printf("Server starting on %s", cfg.Server.Port)
	if err := e.Start("localhost:" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server", err)
	}
}

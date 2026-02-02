package main

import (
	"context"
	"go-ecommerce-service/common/postgresql"
	"go-ecommerce-service/config"
	"go-ecommerce-service/controller"
	"go-ecommerce-service/infrastructure/elasticsearch"
	"go-ecommerce-service/infrastructure/rabbitmq"
	"go-ecommerce-service/internal/jwt"
	"go-ecommerce-service/persistence"
	customMiddleware "go-ecommerce-service/pkg/middleware"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/worker"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
)

func main() {

	/*
		docker-compose up -d postgres redis
		docker-compose up -d rabbitmq
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

	// Database Connection
	var dbPool *pgxpool.Pool
	var dbPoolErr error
	for i := 0; i < 5; i++ {
		dbPool, dbPoolErr = postgresql.GetConnectionPool(ctx, cfg.Database)
		if dbPoolErr == nil {
			log.Info("Veritabanı bağlantısı başarılı! 🚀")
			break
		}
		log.Warnf("Veritabanına bağlanılamadı, tekrar deneniyor (%d/5)... Hata: %v", i+1, dbPoolErr)
		time.Sleep(3 * time.Second)
	}

	if dbPoolErr != nil {
		log.Fatal("Veritabanına bağlanılamadı, pes ediliyor: ", dbPoolErr)
	}

	defer dbPool.Close()

	// RabbitMQ
	var rabbitClient *rabbitmq.RabbitMQClient
	var rabbitErr error

	for i := 0; i < 20; i++ {
		rabbitClient, rabbitErr = rabbitmq.NewRabbitMQClient(
			cfg.RabbitMQ.User,
			cfg.RabbitMQ.Password,
			cfg.RabbitMQ.Host,
			cfg.RabbitMQ.Port,
		)

		if rabbitErr == nil {
			log.Info("🐇 RabbitMQ bağlantısı başarılı! 🚀")
			break
		}

		log.Warnf("RabbitMQ'ya bağlanılamadı (Host: %s), tekrar deneniyor (%d/20)... Hata: %v", cfg.RabbitMQ.Host, i+1, rabbitErr)
		time.Sleep(3 * time.Second)
	}

	if rabbitErr != nil {
		log.Fatal("RabbitMQ bağlantısı kurulamadı, pes ediliyor: ", rabbitErr)
	}
	defer rabbitClient.Close()

	// ElasticSearch
	esClient := elasticsearch.NewElasticSearchClient()
	_ = esClient
	// Dependency Injection
	productRepository := persistence.NewProductRepository(dbPool, esClient)
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
	orderService := service.NewOrderService(orderRepository, rabbitClient)
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

	// Worker
	orderWorker := worker.NewOrderWorker(rabbitClient, orderRepository)
	orderWorker.Start()

	e := echo.New()

	authMiddleware := customMiddleware.AuthMiddleware(authService)

	authController.RegisterRoutes(e)
	productController.RegisterRoutes(e)
	userController.RegisterRoutes(e)
	categoryController.RegisterRoutes(e)
	storeController.RegisterRoutes(e)

	api := e.Group("/api/v1")
	api.Use(authMiddleware)
	cartController.RegisterRoutes(e)
	cartItemController.RegiesterRoutes(e)
	orderController.RegisterRoutes(e)
	orderItemController.RegisterRoutes(e)

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

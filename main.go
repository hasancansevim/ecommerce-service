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
	"go-ecommerce-service/pkg/logger"
	customMiddleware "go-ecommerce-service/pkg/middleware"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/worker"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "go-ecommerce-service/docs"
)

// @title           E-Commerce API
// @version         1.0
// @description     Go ile geliştirilmiş, Dockerize edilmiş, Elasticsearch destekli E-Ticaret Backend'i.
// @termsOfService  http://swagger.io/terms/

// @contact.name    Hasan Can Sevim
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1
func main() {

	logger.Init()

	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load config")
	}
	jwt.Initialize(cfg.Auth.JWTSecret)

	ctx := context.Background()

	// Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.Redis.Host + ":" + cfg.Redis.Port,
	})
	if redisConnectionErr := rdb.Ping(ctx).Err(); redisConnectionErr != nil {
		log.Fatal().Err(redisConnectionErr).Msg("Failed connect to redis")
	}

	// Database Connection
	var dbPool *pgxpool.Pool
	var dbPoolErr error
	for i := 0; i < 5; i++ {
		dbPool, dbPoolErr = postgresql.GetConnectionPool(ctx, cfg.Database)
		if dbPoolErr == nil {
			log.Info().Msg("Veritabanı bağlantısı başarılı! 🚀")
			break
		}
		log.Warn().Err(dbPoolErr).Msg("Veritabanına bağlanılamadı, tekrar deneniyor...")
		time.Sleep(3 * time.Second)
	}

	if dbPoolErr != nil {
		log.Fatal().Err(dbPoolErr).Msg("Veritabanına bağlanılamadı, pes ediliyor: ")
	}

	defer dbPool.Close()

	// RabbitMQ
	var rabbitClient *rabbitmq.RabbitMQClient
	var rabbitErr error

	for i := 0; i < 20; i++ {
		rabbitClient, rabbitErr = rabbitmq.NewRabbitMQClient(cfg.RabbitMQ)

		if rabbitErr == nil {
			log.Info().Msg("RabbitMQ bağlantısı başarılı! 🚀")
			break
		}

		log.Warn().Err(dbPoolErr).Msg("RabbitMQ'ya bağlanılamadı, tekrar deneniyor...")
		time.Sleep(3 * time.Second)
	}

	if rabbitErr != nil {
		log.Fatal().Err(rabbitErr).Msg("RabbitMQ bağlantısı kurulamadı, pes ediliyor.")
	}
	defer rabbitClient.Close()

	// ElasticSearch
	esClient := elasticsearch.NewElasticSearchClient(cfg.ElasticSearch)
	if _, err := esClient.Info(); err != nil {
		log.Fatal().Err(err).Msg("ElasticSearch bağlantısı kurulamadı.")
	}
	log.Info().Msg("ElasticSearch bağlantısı başarılı! 🚀")

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

	e.GET("/swagger/*", echoSwagger.WrapHandler)

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

	go func() {
		if err := e.Start(":" + cfg.Server.Port); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Sunucu başlatılamadı")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Warn().Msg("⚠️ Kapanma sinyali alındı! Uygulama zarifçe kapatılıyor...")
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Sunucu zorla kapatıldı")
	}

	log.Info().Msg("🔌 Bağlantılar kapatılıyor...")
	log.Info().Msg("👋 Güle güle! Sistem başarıyla kapandı.")
}

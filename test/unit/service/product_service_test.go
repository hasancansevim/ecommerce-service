package service

import (
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/dto"
	"go-ecommerce-service/service"
	mock_repository "go-ecommerce-service/test/mock/repository"
)

func TestProductService(t *testing.T) {
	// --- SENARYO 1: Başarılı Getirme (Redis Boş -> DB Dolu -> Redis Yaz) ---
	t.Run("GetProductById_Success_FromDB", func(t *testing.T) {
		// 1. İzolasyon: Her test için taze mocklar oluştur
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockIProductRepository(ctrl)
		db, mockRedis := redismock.NewClientMock()
		productService := service.NewProductService(mockRepo, db)

		// 2. Veri Hazırlığı
		productId := int64(1)
		// Domain'den gelen veri
		domainProduct := domain.Product{
			Id: 1, Name: "Laptop", Price: 15000,
			Description: "Test Desc", StockQuantity: 5,
		}

		// Redis'e yazılması beklenen veri (Service içindeki dönüşümü taklit ediyoruz)
		expectedResponse := dto.ProductResponse{
			Id: domainProduct.Id, Name: domainProduct.Name, Price: domainProduct.Price,
			Description: domainProduct.Description, StockQuantity: domainProduct.StockQuantity,
		}
		expectedJson, _ := json.Marshal(expectedResponse)

		// 3. Beklentiler
		// A. Redis'e sor -> YOK
		mockRedis.ExpectGet("product:1").RedisNil()

		// B. DB'ye sor -> VAR
		mockRepo.EXPECT().GetProductById(productId).Return(domainProduct, nil)

		// C. Redis'e YAZ (Düzeltme: gomock.Any() yerine gerçek JSON verisini bekliyoruz)
		mockRedis.ExpectSet("product:1", expectedJson, 10*time.Minute).SetVal("OK")

		// 4. Çalıştır
		result, err := productService.GetProductById(productId)

		// 5. Kontrol
		assert.NoError(t, err)
		assert.Equal(t, domainProduct.Name, result.Name)

		if err := mockRedis.ExpectationsWereMet(); err != nil {
			t.Error("Redis işlemleri eksik kaldı:", err)
		}
	})

	// --- SENARYO 2: Ürün Bulunamadı ---
	t.Run("GetProductById_NotFound", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockIProductRepository(ctrl)
		db, mockRedis := redismock.NewClientMock()
		productService := service.NewProductService(mockRepo, db)

		productId := int64(99)

		// Beklentiler
		mockRedis.ExpectGet("product:99").RedisNil()
		mockRepo.EXPECT().GetProductById(productId).Return(domain.Product{}, errors.New("product not found"))

		// Çalıştır
		_, err := productService.GetProductById(productId)

		// Kontrol
		assert.Error(t, err)
		assert.Equal(t, "product not found", err.Error())

		if err := mockRedis.ExpectationsWereMet(); err != nil {
			t.Error("Redis işlemleri eksik kaldı:", err)
		}
	})

	// --- SENARYO 3: Ürün Ekleme ---
	t.Run("AddProduct_Success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepo := mock_repository.NewMockIProductRepository(ctrl)
		db, _ := redismock.NewClientMock() // Redis kullanmıyoruz ama servise lazım
		productService := service.NewProductService(mockRepo, db)

		// Düzeltme: Validasyonun geçmesi için tüm zorunlu alanları doldurduk
		req := dto.CreateProductRequest{
			Name:          "Gaming Mouse",
			Price:         500,
			StockQuantity: 10,
			StoreId:       1,
			Description:   "Yüksek DPI'lı mouse",
		}

		// Mock Beklentisi
		mockRepo.EXPECT().AddProduct(gomock.Any()).Return(domain.Product{
			Id: 1, Name: req.Name, Price: req.Price,
		}, nil)

		// Çalıştır
		_, err := productService.AddProduct(req)

		// Kontrol
		assert.NoError(t, err)
	})
}

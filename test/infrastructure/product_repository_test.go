package infrastructure

import (
	"context"
	"fmt"
	"go-ecommerce-service/common/postgresql"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"
)

var productRepository persistence.IProductRepository
var dbPool *pgxpool.Pool
var product_ctx context.Context

func ProductTestMain(m *testing.M) {
	product_ctx = context.Background()
	dbPool = postgresql.GetConnectionPool(product_ctx, postgresql.Config{
		Host:                  "localhost",
		Port:                  "6432",
		Username:              "postgres",
		Password:              "123456",
		DbName:                "ecommerce",
		MaxConnections:        "10",
		MaxConnectionIdleTime: "10s",
	})
	productRepository = persistence.NewProductRepository(dbPool)

	fmt.Println("Before All Tests - Product")
	exitCode := m.Run()
	fmt.Println("After All Tests - Product")
	os.Exit(exitCode)
}

func setup(ctx context.Context, dbPool *pgxpool.Pool) {
	ProductTestDataInitalize(ctx, dbPool)
}

func clear(ctx context.Context, dbPool *pgxpool.Pool) {
	TruncateProductTestData(ctx, dbPool)
}

func TestGetAllProducts(t *testing.T) {
	setup(product_ctx, dbPool)
	t.Run("Get All Products", func(t *testing.T) {
		allProducts := productRepository.GetAllProducts()
		assert.Equal(t, 4, len(allProducts))

		expectedProducts := []domain.Product{
			{
				Id:       1,
				Name:     "Laptop",
				Price:    20000.0,
				Discount: 10.0,
				Store:    "Teknosa",
			},
			{
				Id:       2,
				Name:     "Klavye",
				Price:    800.0,
				Discount: 0.0,
				Store:    "Teknosa",
			},
			{
				Id:       3,
				Name:     "Mouse",
				Price:    500.0,
				Discount: 10.0,
				Store:    "Teknosa",
			},
			{
				Id:       4,
				Name:     "Ütü",
				Price:    200.0,
				Discount: 0.0,
				Store:    "Güzel Evim",
			},
		}

		assert.Equal(t, expectedProducts, allProducts)
	})
	clear(product_ctx, dbPool)
}

func TestGetAllProductsByStoreName(t *testing.T) {
	setup(product_ctx, dbPool)
	t.Run("Get All Products by Store Name", func(t *testing.T) {
		allProductsByStore := productRepository.GetAllByStoreName("Teknosa")

		expectedProducts := []domain.Product{
			{
				Id:       1,
				Name:     "Laptop",
				Price:    20000.0,
				Discount: 10.0,
				Store:    "Teknosa",
			},
			{
				Id:       2,
				Name:     "Klavye",
				Price:    800.0,
				Discount: 0.0,
				Store:    "Teknosa",
			},
			{
				Id:       3,
				Name:     "Mouse",
				Price:    500.0,
				Discount: 10.0,
				Store:    "Teknosa",
			},
		}

		assert.Equal(t, 3, len(allProductsByStore))
		assert.Equal(t, expectedProducts, allProductsByStore)
	})
	clear(product_ctx, dbPool)
}

func TestGetProductById(t *testing.T) {
	setup(product_ctx, dbPool)
	t.Run("Get Product by Id", func(t *testing.T) {
		actualProduct, _ := productRepository.GetProductById(3)
		_, actualProductErr := productRepository.GetProductById(5)

		expectedProduct := domain.Product{
			Id:       3,
			Name:     "Mouse",
			Price:    500.0,
			Discount: 10.0,
			Store:    "Teknosa",
		}

		assert.Equal(t, expectedProduct, actualProduct)
		assert.Equal(t, "Product not found with id : 5", actualProductErr.Error())
	})
	clear(product_ctx, dbPool)
}

func TestUpdatePrice(t *testing.T) {
	setup(product_ctx, dbPool)
	t.Run("Update price from product", func(t *testing.T) {
		productRepository.UpdatePrice(1, 18000.0)
		updatedProduct, _ := productRepository.GetProductById(1)
		expectedProducts := domain.Product{
			Id:       1,
			Name:     "Laptop",
			Price:    18000.0,
			Discount: 10.0,
			Store:    "Teknosa",
		}
		assert.Equal(t, expectedProducts, updatedProduct)
	})
	clear(product_ctx, dbPool)
}

func TestDeleteProductById(t *testing.T) {
	setup(product_ctx, dbPool)
	t.Run("Delete Product by Id", func(t *testing.T) {
		productRepository.DeleteProductById(3)
		actualProduct := productRepository.GetAllProducts()
		expectedProducts := []domain.Product{
			{
				Id:       1,
				Name:     "Laptop",
				Price:    20000.0,
				Discount: 10.0,
				Store:    "Teknosa",
			},
			{
				Id:       2,
				Name:     "Klavye",
				Price:    800.0,
				Discount: 0.0,
				Store:    "Teknosa",
			},
			{
				Id:       4,
				Name:     "Ütü",
				Price:    200.0,
				Discount: 0.0,
				Store:    "Güzel Evim",
			},
		}
		assert.Equal(t, expectedProducts, actualProduct)
		assert.Equal(t, 3, len(actualProduct))
	})
	clear(product_ctx, dbPool)
}

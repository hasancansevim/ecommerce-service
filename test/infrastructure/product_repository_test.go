package infrastructure

import (
	"go-ecommerce-service/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProducts(t *testing.T) {
	setupProductTestData()
	defer cleanupProductTestData()

	t.Run("Get All Products", func(t *testing.T) {
		actualProducts := productRepository.GetAllProducts()
		expectedProducts := []domain.Product{
			{Name: "Laptop", Price: 20000.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Klavye", Price: 800.0, Discount: 0.0, Store: "Teknosa"},
			{Name: "Mouse", Price: 500.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Ütü", Price: 200.0, Discount: 0.0, Store: "Güzel Evim"},
		}
		assert.Equal(t, 4, len(actualProducts))
		for i, expectedProduct := range expectedProducts {
			assert.Equal(t, expectedProduct.Name, actualProducts[i].Name)
			assert.Equal(t, expectedProduct.Price, actualProducts[i].Price)
			assert.Equal(t, expectedProduct.Discount, actualProducts[i].Discount)
			assert.Equal(t, expectedProduct.Store, actualProducts[i].Store)
		}
	})
}

func TestGetAllProductsByStoreName(t *testing.T) {
	setupProductTestData()
	defer cleanupProductTestData()

	t.Run("Get All Products by Store Name", func(t *testing.T) {
		actualProducts := productRepository.GetAllProductsByStoreName("Teknosa")
		expectedProducts := []domain.Product{
			{Name: "Laptop", Price: 20000.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Klavye", Price: 800.0, Discount: 0.0, Store: "Teknosa"},
			{Name: "Mouse", Price: 500.0, Discount: 10.0, Store: "Teknosa"},
		}

		assert.Equal(t, 3, len(actualProducts))
		for i, expectedProduct := range expectedProducts {
			assert.Equal(t, expectedProduct.Name, actualProducts[i].Name)
			assert.Equal(t, expectedProduct.Price, actualProducts[i].Price)
			assert.Equal(t, expectedProduct.Discount, actualProducts[i].Discount)
			assert.Equal(t, expectedProduct.Store, actualProducts[i].Store)
		}
	})
}

func TestGetProductById(t *testing.T) {
	setupProductTestData()
	defer cleanupProductTestData()
	t.Run("Get Product by Product Id", func(t *testing.T) {
		actualProduct, _ := productRepository.GetProductById(int64(8)) // id ye bak db den
		panic(actualProduct.Id)
		expectedProduct := domain.Product{Name: "Laptop", Price: 20000.0, Discount: 10.0, Store: "Teknosa"}
		assert.Equal(t, expectedProduct.Name, actualProduct.Name)
		assert.Equal(t, expectedProduct.Price, actualProduct.Price)
		assert.Equal(t, expectedProduct.Discount, actualProduct.Discount)
		assert.Equal(t, expectedProduct.Store, actualProduct.Store)
	})
}

func TestAddProduct(t *testing.T) {
	setupProductTestData()
	defer cleanupProductTestData()
	t.Run("Add Product", func(t *testing.T) {
		productRepository.AddProduct(domain.Product{
			Name:     "Laptop 2",
			Price:    20000.0,
			Discount: 10.0,
			Store:    "Teknosa",
		})
		actualProducts := productRepository.GetAllProducts()
		expectedProducts := []domain.Product{
			{Name: "Laptop", Price: 20000.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Klavye", Price: 800.0, Discount: 0.0, Store: "Teknosa"},
			{Name: "Mouse", Price: 500.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Ütü", Price: 200.0, Discount: 0.0, Store: "Güzel Evim"},
			{Name: "Laptop 2", Price: 20000.0, Discount: 10.0, Store: "Teknosa"},
		}
		assert.Equal(t, 5, len(actualProducts))
		for i, expectedProduct := range expectedProducts {
			assert.Equal(t, expectedProduct.Name, actualProducts[i].Name)
			assert.Equal(t, expectedProduct.Price, actualProducts[i].Price)
			assert.Equal(t, expectedProduct.Discount, actualProducts[i].Discount)
			assert.Equal(t, expectedProduct.Store, actualProducts[i].Store)
		}
	})
}

func TestDeleteProductById(t *testing.T) {
	setupProductTestData()
	defer cleanupProductTestData()
	t.Run("Delete Product by Product Id", func(t *testing.T) {
		productRepository.DeleteProductById(int64(4))
		actualProducts := productRepository.GetAllProducts()
		expectedProducts := []domain.Product{
			{Name: "Laptop", Price: 20000.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Klavye", Price: 800.0, Discount: 0.0, Store: "Teknosa"},
			{Name: "Mouse", Price: 500.0, Discount: 10.0, Store: "Teknosa"},
		}
		assert.Equal(t, 3, len(actualProducts))
		for i, expectedProduct := range expectedProducts {
			assert.Equal(t, expectedProduct.Name, actualProducts[i].Name)
			assert.Equal(t, expectedProduct.Price, actualProducts[i].Price)
			assert.Equal(t, expectedProduct.Discount, actualProducts[i].Discount)
			assert.Equal(t, expectedProduct.Store, actualProducts[i].Store)
		}
	})
}

func TestUpdatePrice(t *testing.T) {
	setupProductTestData()
	defer cleanupProductTestData()
	t.Run("Update Price", func(t *testing.T) {
		productRepository.UpdatePrice(int64(4), 20000.0)
		actualProducts := productRepository.GetAllProducts()
		expectedProducts := []domain.Product{
			{Name: "Laptop", Price: 20000.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Klavye", Price: 800.0, Discount: 0.0, Store: "Teknosa"},
			{Name: "Mouse", Price: 500.0, Discount: 10.0, Store: "Teknosa"},
			{Name: "Ütü", Price: 20000.0, Discount: 0.0, Store: "Güzel Evim"},
		}
		assert.Equal(t, expectedProducts[4].Price, actualProducts[4].Price)
	})
}

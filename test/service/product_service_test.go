package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/model"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var productService service.IProductService

var initialProducts = []domain.Product{
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
}

func TestMain(m *testing.M) {

	fakeProductRepository := NewFakeProductRepository(initialProducts)
	productService = service.NewProductService(fakeProductRepository)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func Test_ShouldGetAllProducts(t *testing.T) {
	t.Run("Should Getting All Products", func(t *testing.T) {
		actualProducts := productService.GetAllProducts()
		assert.Equal(t, 2, len(actualProducts))
		assert.Equal(t, initialProducts, actualProducts)
	})
}

func Test_WhenNoValidationErrorOccurred_ShouldAddProduct(t *testing.T) {
	t.Run("When No Validation Error Occurred Should Add Product", func(t *testing.T) {
		productService.AddProduct(model.ProductCreate{
			Name:     "Ütü Masası",
			Price:    1500.0,
			Discount: 20,
			Store:    "Güzel Evim",
		})

		actualProducts := productService.GetAllProducts()

		assert.Equal(t, 3, len(actualProducts))
		assert.Equal(t, domain.Product{
			Id:       3,
			Name:     "Ütü Masası",
			Price:    1500.0,
			Discount: 20,
			Store:    "Güzel Evim",
		}, actualProducts[len(actualProducts)-1])
	})
}

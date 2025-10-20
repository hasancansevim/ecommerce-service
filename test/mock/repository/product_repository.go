package repository

import (
	"go-ecommerce-service/domain"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) GetAllProducts() []domain.Product {
	args := m.Called()
	return args.Get(0).([]domain.Product)
}

func (m *MockProductRepository) GetProductById(productId int64) (domain.Product, error) {
	args := m.Called(productId)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductRepository) GetAllProductsByStoreName(storeName string) []domain.Product {
	args := m.Called(storeName)
	return args.Get(0).([]domain.Product)
}

func (m *MockProductRepository) AddProduct(product domain.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductRepository) DeleteProductById(productId int64) error {
	args := m.Called(productId)
	return args.Error(0)
}
func (m *MockProductRepository) UpdatePrice(productId int64, newPrice float32) error {
	args := m.Called(productId, newPrice)
	return args.Error(0)
}

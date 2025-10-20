package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"

	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetAllProducts() []domain.Product {
	args := m.Called()
	return args.Get(0).([]domain.Product)
}

func (m *MockProductService) GetAllProductsByStoreName(storeName string) []domain.Product {
	args := m.Called(storeName)
	return args.Get(0).([]domain.Product)
}

func (m *MockProductService) GetProductById(productId int64) (domain.Product, error) {
	args := m.Called(productId)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockProductService) AddProduct(productCreate model.ProductCreate) error {
	args := m.Called(productCreate)
	return args.Error(0)
}

func (m *MockProductService) DeleteProductById(productId int64) error {
	args := m.Called(productId)
	return args.Error(0)
}

func (m *MockProductService) UpdatePrice(productId int64, newPrice float32) error {
	args := m.Called(productId, newPrice)
	return args.Error(0)
}

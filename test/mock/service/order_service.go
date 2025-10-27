package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"

	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) CreateOrder(order model.OrderCreate) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderService) GetOrderById(orderId int64) domain.Order {
	args := m.Called(orderId)
	return args.Get(0).(domain.Order)
}

func (m *MockOrderService) GetOrdersByUserId(userId int64) ([]domain.Order, error) {
	args := m.Called(userId)
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderService) GetAllOrders() ([]domain.Order, error) {
	args := m.Called()
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderService) UpdateOrderStatus(orderId int64, status bool) error {
	args := m.Called(orderId, status)
	return args.Error(0)
}

func (m *MockOrderService) DeleteOrderById(orderId int64) error {
	args := m.Called(orderId)
	return args.Error(0)
}

func (m *MockOrderService) UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) error {
	args := m.Called(orderId, newTotalPrice)
	return args.Error(0)
}

func (m *MockOrderService) GetOrdersByStatus(status string) ([]domain.Order, error) {
	args := m.Called(status)
	return args.Get(0).([]domain.Order), args.Error(1)
}

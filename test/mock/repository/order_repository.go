package mock

import (
	"go-ecommerce-service/domain"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) CreateOrder(order domain.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrderById(orderId int64) domain.Order {
	args := m.Called(orderId)
	return args.Get(0).(domain.Order)
}

func (m *MockOrderRepository) GetOrdersByUserId(userId int64) ([]domain.Order, error) {
	args := m.Called(userId)
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderRepository) GetAllOrders() ([]domain.Order, error) {
	args := m.Called()
	return args.Get(0).([]domain.Order), args.Error(1)
}

func (m *MockOrderRepository) UpdateOrderStatus(orderId int64, status bool) error {
	args := m.Called(orderId, status)
	return args.Error(0)
}

func (m *MockOrderRepository) DeleteOrderById(orderId int64) error {
	args := m.Called(orderId)
	return args.Error(0)
}

func (m *MockOrderRepository) UpdateOrderTotalPrice(orderId int64, newTotalPrice float32) error {
	args := m.Called(orderId, newTotalPrice)
	return args.Error(0)
}

func (m *MockOrderRepository) GetOrdersByStatus(status string) ([]domain.Order, error) {
	args := m.Called(status)
	return args.Get(0).([]domain.Order), args.Error(1)
}

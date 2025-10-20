package mock

import (
	"go-ecommerce-service/domain"

	"github.com/stretchr/testify/mock"
)

type MockOrderItemRepository struct {
	mock.Mock
}

func (m *MockOrderItemRepository) AddOrderItem(orderItem domain.OrderItem) error {
	args := m.Called(orderItem)
	return args.Error(0)
}

func (m *MockOrderItemRepository) GetOrderItemById(orderItem_id int64) (domain.OrderItem, error) {
	args := m.Called(orderItem_id)
	return args.Get(0).(domain.OrderItem), args.Error(1)
}

func (m *MockOrderItemRepository) GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error) {
	args := m.Called(orderId)
	return args.Get(0).([]domain.OrderItem), args.Error(1)
}

func (m *MockOrderItemRepository) GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error) {
	args := m.Called(productId)
	return args.Get(0).([]domain.OrderItem), args.Error(1)
}

func (m *MockOrderItemRepository) UpdateOrderItem(orderItem_id int64, orderItem domain.OrderItem) error {
	args := m.Called(orderItem_id, orderItem)
	return args.Error(0)
}

func (m *MockOrderItemRepository) UpdateOrderItemQuantity(orderItem_id int64, quantity int) error {
	args := m.Called(orderItem_id, quantity)
	return args.Error(0)
}

func (m *MockOrderItemRepository) DeleteOrderItemById(orderItem_id int64) error {
	args := m.Called(orderItem_id)
	return args.Error(0)
}

func (m *MockOrderItemRepository) DeleteAllOrderItemsByOrderId(orderId int64) error {
	args := m.Called(orderId)
	return args.Error(0)
}

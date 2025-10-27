package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"

	"github.com/stretchr/testify/mock"
)

type MockOrderItemService struct {
	mock.Mock
}

func (orderItemService *MockOrderItemService) AddOrderItem(orderItemCreate model.OrderItemCreate) error {
	args := orderItemService.Called(orderItemCreate)
	return args.Error(0)
}

func (orderItemService *MockOrderItemService) GetOrderItemById(orderItem_id int64) (domain.OrderItem, error) {
	args := orderItemService.Called(orderItem_id)
	return args.Get(0).(domain.OrderItem), args.Error(1)
}

func (orderItemService *MockOrderItemService) GetOrderItemsByOrderId(orderId int64) ([]domain.OrderItem, error) {
	args := orderItemService.Called(orderId)
	return args.Get(0).([]domain.OrderItem), args.Error(1)
}

func (orderItemService *MockOrderItemService) GetOrderItemsByProductId(productId int64) ([]domain.OrderItem, error) {
	args := orderItemService.Called(productId)
	return args.Get(0).([]domain.OrderItem), args.Error(1)
}

func (orderItemService *MockOrderItemService) UpdateOrderItem(orderItem_id int64, orderItem domain.OrderItem) error {
	args := orderItemService.Called(orderItem_id, orderItem)
	return args.Error(0)
}

func (orderItemService *MockOrderItemService) UpdateOrderItemQuantity(orderItem_id int64, quantity int) error {
	args := orderItemService.Called(orderItem_id, quantity)
	return args.Error(0)
}

func (orderItemService *MockOrderItemService) DeleteOrderItemById(orderItem_id int64) error {
	args := orderItemService.Called(orderItem_id)
	return args.Error(0)
}

func (orderItemService *MockOrderItemService) DeleteAllOrderItemsByOrderId(orderId int64) error {
	args := orderItemService.Called(orderId)
	return args.Error(0)
}

package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service"
	"go-ecommerce-service/test/fixture"
	mock2 "go-ecommerce-service/test/mock/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type OrderItemServiceTestSuite struct {
	suite.Suite
	mockOrderItemRepo *mock2.MockOrderItemRepository
	orderItemService  service.IOrderItemService
}

func (suite *OrderItemServiceTestSuite) SetupTest() {
	suite.mockOrderItemRepo = new(mock2.MockOrderItemRepository)
	suite.orderItemService = service.NewOrderItemService(suite.mockOrderItemRepo)
}

func (suite *OrderItemServiceTestSuite) TestAddOrderItem_Success() {
	createOrderItem := fixture.CreateTestOrderItemCreate()

	suite.mockOrderItemRepo.On("AddOrderItem", mock.MatchedBy(func(orderItem domain.OrderItem) bool {
		return orderItem.OrderId == createOrderItem.OrderId &&
			orderItem.ProductId == createOrderItem.ProductId &&
			orderItem.Quantity == createOrderItem.Quantity &&
			orderItem.Price == createOrderItem.Price
	})).Return(nil)

	err := suite.orderItemService.AddOrderItem(createOrderItem)

	assert.NoError(suite.T(), err)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestGetOrderItemById_Success() {
	expectedOrderItem := fixture.CreateTestOrderItem()

	suite.mockOrderItemRepo.On("GetOrderItemById", expectedOrderItem.Id).Return(expectedOrderItem, nil)

	orderItem, err := suite.orderItemService.GetOrderItemById(expectedOrderItem.Id)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrderItem, orderItem)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestGetOrderItemById_OrderItemNotFound() {
	orderItemId := int64(999)

	suite.mockOrderItemRepo.On("GetOrderItemById", orderItemId).
		Return(domain.OrderItem{}, errors.New("order item not found"))

	orderItem, err := suite.orderItemService.GetOrderItemById(orderItemId)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.OrderItem{}, orderItem)
	assert.Contains(suite.T(), err.Error(), "order item not found")
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestGetOrderItemsByOrderId_Success() {
	orderId := int64(1)
	expectedOrderItems := []domain.OrderItem{
		fixture.CreateTestOrderItem(),
		{Id: 2, OrderId: orderId, ProductId: 2, Quantity: 3, Price: 75.50},
	}

	suite.mockOrderItemRepo.On("GetOrderItemsByOrderId", orderId).Return(expectedOrderItems, nil)

	orderItems, err := suite.orderItemService.GetOrderItemsByOrderId(orderId)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrderItems, orderItems)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestGetOrderItemsByProductId_Success() {
	productId := int64(1)
	expectedOrderItems := []domain.OrderItem{
		fixture.CreateTestOrderItem(),
		{Id: 3, OrderId: 2, ProductId: productId, Quantity: 1, Price: 99.99},
	}

	suite.mockOrderItemRepo.On("GetOrderItemsByProductId", productId).Return(expectedOrderItems, nil)

	orderItems, err := suite.orderItemService.GetOrderItemsByProductId(productId)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrderItems, orderItems)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestUpdateOrderItem_Success() {
	orderItemId := int64(1)
	updateData := fixture.CreateTestOrderItem()

	suite.mockOrderItemRepo.On("UpdateOrderItem", orderItemId, updateData).Return(nil)

	err := suite.orderItemService.UpdateOrderItem(orderItemId, updateData)

	assert.NoError(suite.T(), err)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestUpdateOrderItemQuantity_Success() {
	orderItemId := int64(1)
	newQuantity := 5

	suite.mockOrderItemRepo.On("UpdateOrderItemQuantity", orderItemId, newQuantity).Return(nil)

	err := suite.orderItemService.UpdateOrderItemQuantity(orderItemId, newQuantity)

	assert.NoError(suite.T(), err)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestDeleteOrderItemById_Success() {
	orderItemId := int64(1)

	suite.mockOrderItemRepo.On("DeleteOrderItemById", orderItemId).Return(nil)

	err := suite.orderItemService.DeleteOrderItemById(orderItemId)

	assert.NoError(suite.T(), err)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func (suite *OrderItemServiceTestSuite) TestDeleteAllOrderItemsByOrderId_Success() {
	orderId := int64(1)

	suite.mockOrderItemRepo.On("DeleteAllOrderItemsByOrderId", orderId).Return(nil)

	err := suite.orderItemService.DeleteAllOrderItemsByOrderId(orderId)

	assert.NoError(suite.T(), err)
	suite.mockOrderItemRepo.AssertExpectations(suite.T())
}

func TestOrderItemServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderItemServiceTestSuite))
}

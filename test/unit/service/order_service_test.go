package service

import (
	"errors"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service"
	"go-ecommerce-service/test/fixture"
	"go-ecommerce-service/test/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrderServiceTestSuite struct {
	suite.Suite
	mockOrderRepo *mock.MockOrderRepository
	orderService  service.IOrderService
}

func (suite *OrderServiceTestSuite) SetupTest() {
	suite.mockOrderRepo = new(mock.MockOrderRepository)
	suite.orderService = service.NewOrderService(suite.mockOrderRepo)
}

func (suite *OrderServiceTestSuite) TestCreateOrderSuccess() {

	suite.mockOrderRepo.On("CreateOrder", domain.Order{
		Id:         0,
		UserId:     fixture.CreateTestOrderCreate().UserId,
		TotalPrice: fixture.CreateTestOrderCreate().TotalPrice,
		Status:     fixture.CreateTestOrderCreate().Status,
	}).Return(nil)

	err := suite.orderService.CreateOrder(fixture.CreateTestOrderCreate())

	assert.NoError(suite.T(), err)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestCreateOrder_Fail() {
	testOrder := fixture.CreateTestOrder()
	testOrder.Id = int64(0)
	suite.mockOrderRepo.On("CreateOrder", domain.Order{
		Id:         testOrder.Id,
		UserId:     testOrder.UserId,
		TotalPrice: testOrder.TotalPrice,
		Status:     testOrder.Status,
	}).Return(errors.New("user not found"))
	err := suite.orderService.CreateOrder(fixture.CreateTestOrderCreate())

	assert.Error(suite.T(), err)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestGetOrderById_Success() {
	expectedOrder := fixture.CreateTestOrder()

	suite.mockOrderRepo.On("GetOrderById", expectedOrder.Id).Return(expectedOrder)

	getOrderById := suite.orderService.GetOrderById(expectedOrder.Id)

	assert.Equal(suite.T(), expectedOrder, getOrderById)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestGetOrdersByUserId_Success() {
	expectedOrders := fixture.CreateTestGetOrdersByUserId()
	suite.mockOrderRepo.On("GetOrdersByUserId", int64(1)).
		Return(expectedOrders, nil)

	getOrdersByUserId, err := suite.orderService.GetOrdersByUserId(int64(1))

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrders, getOrdersByUserId)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestGetAllOrders() {
	expectedOrders := fixture.CreateTestGetAllOrders()
	suite.mockOrderRepo.On("GetAllOrders").Return(expectedOrders, nil)

	orders, err := suite.orderService.GetAllOrders()

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrders, orders)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestUpdateOrderStatus() {

	suite.mockOrderRepo.On("UpdateOrderStatus", int64(1), false).Return(nil)
	err := suite.orderService.UpdateOrderStatus(int64(1), false)

	assert.NoError(suite.T(), err)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestDeleteOrderById_Success() {
	orderId := int64(1)

	suite.mockOrderRepo.On("DeleteOrderById", orderId).Return(nil)

	err := suite.orderService.DeleteOrderById(orderId)

	assert.NoError(suite.T(), err)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestUpdateOrderTotalPrice_Success() {
	orderId := int64(1)
	newPrice := float32(150.50)

	suite.mockOrderRepo.On("UpdateOrderTotalPrice", orderId, newPrice).Return(nil)

	err := suite.orderService.UpdateOrderTotalPrice(orderId, newPrice)

	assert.NoError(suite.T(), err)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func (suite *OrderServiceTestSuite) TestGetOrdersByStatus_Success() {
	expectedOrders := []domain.Order{
		fixture.CreateTestOrder(),
		{Id: 2, UserId: 1, TotalPrice: 200.0, Status: true},
	}

	suite.mockOrderRepo.On("GetOrdersByStatus", "true").Return(expectedOrders, nil)

	orders, err := suite.orderService.GetOrdersByStatus("true")

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedOrders, orders)
	suite.mockOrderRepo.AssertExpectations(suite.T())
}

func TestOrderTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

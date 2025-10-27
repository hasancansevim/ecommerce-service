package controller

import (
	"bytes"
	"encoding/json"
	"go-ecommerce-service/controller"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/test/mock/service"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OrderControllerTestSuite struct {
	suite.Suite
	echo             *echo.Echo
	mockOrderService *service.MockOrderService
	orderController  *controller.OrderController
}

func (suite *OrderControllerTestSuite) SetupTest() {
	suite.echo = echo.New()
	suite.mockOrderService = new(service.MockOrderService)
	suite.orderController = controller.NewOrderController(suite.mockOrderService)
}

func (suite *OrderControllerTestSuite) TestCreateOrder_Success() {
	jsonData := []byte(`{
		"user_id": 1,
		"total_price": 199.99,
		"status": true
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	expectedOrder := model.OrderCreate{
		UserId:     1,
		TotalPrice: 199.99,
		Status:     true,
	}

	suite.mockOrderService.On("CreateOrder", expectedOrder).Return(nil)

	err := suite.orderController.CreateOrder(ctx)

	suite.T().Logf("Status: %d, Body: %s", rec.Code, rec.Body.String())

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	suite.mockOrderService.AssertExpectations(suite.T())
}

func (suite *OrderControllerTestSuite) TestCreateOrder_InvalidInput() {
	jsonData := []byte(`{
		"user_id": -1,
		"total_price": -100.0,
		"status": ""
	}`)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewReader(jsonData))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	err := suite.orderController.CreateOrder(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)

	var response map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	suite.mockOrderService.AssertNotCalled(suite.T(), "CreateOrder")
}

func (suite *OrderControllerTestSuite) TestGetOrderById_Success() {
	expectedOrder := domain.Order{
		Id:         1,
		UserId:     1,
		TotalPrice: 199.99,
		Status:     true,
	}

	suite.mockOrderService.On("GetOrderById", int64(1)).Return(expectedOrder)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/orders/1", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.orderController.GetOrderById(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)

	var response domain.Order
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Equal(suite.T(), expectedOrder.Id, response.Id)
	assert.Equal(suite.T(), expectedOrder.TotalPrice, response.TotalPrice)
	suite.mockOrderService.AssertExpectations(suite.T())
}

func (suite *OrderControllerTestSuite) TestGetOrderById_InvalidID() {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/orders/invalid", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("invalid")

	err := suite.orderController.GetOrderById(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	suite.mockOrderService.AssertNotCalled(suite.T(), "GetOrderById")
}

func (suite *OrderControllerTestSuite) TestGetOrdersByUserId_Success() {
	expectedOrders := []domain.Order{
		{
			Id:         1,
			UserId:     1,
			TotalPrice: 199.99,
			Status:     true,
		},
		{
			Id:         2,
			UserId:     1,
			TotalPrice: 299.99,
			Status:     true,
		},
	}

	suite.mockOrderService.On("GetOrdersByUserId", int64(1)).Return(expectedOrders, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/orders/get-orders-by-user-id?user_id=1", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("user_id")
	ctx.SetParamValues("1")

	err := suite.orderController.GetOrdersByUserId(ctx)
	suite.T().Logf("Status: %d, Body: %s", rec.Code, rec.Body.String())
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)

	var response []map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Len(suite.T(), response, 2)
	assert.Equal(suite.T(), float64(expectedOrders[0].UserId), response[0]["user_id"])
	suite.mockOrderService.AssertExpectations(suite.T())
}

func (suite *OrderControllerTestSuite) TestGetAllOrders_Success() {
	expectedOrders := []domain.Order{
		{
			Id:         1,
			UserId:     1,
			TotalPrice: 199.99,
			Status:     true,
		},
		{
			Id:         2,
			UserId:     2,
			TotalPrice: 299.99,
			Status:     true,
		},
	}

	suite.mockOrderService.On("GetAllOrders").Return(expectedOrders, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/orders/get-all-orders", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	err := suite.orderController.GetAllOrders(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)

	var response []domain.Order
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Len(suite.T(), response, 2)
	suite.mockOrderService.AssertExpectations(suite.T())
}

func (suite *OrderControllerTestSuite) TestUpdateOrderStatus_Success() {
	suite.mockOrderService.On("UpdateOrderStatus", int64(1), true).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/orders/update-order-status/1?status=true", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.orderController.UpdateOrderStatus(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	suite.mockOrderService.AssertExpectations(suite.T())
}

func (suite *OrderControllerTestSuite) TestDeleteOrderById_Success() {
	suite.mockOrderService.On("DeleteOrderById", int64(1)).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/orders/1", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.orderController.DeleteOrderById(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	suite.mockOrderService.AssertExpectations(suite.T())
}

func (suite *OrderControllerTestSuite) TestUpdateOrderTotalPrice_Success() {
	suite.mockOrderService.On("UpdateOrderTotalPrice", int64(1), float32(250.50)).Return(nil)

	req := httptest.NewRequest(http.MethodPut, "/api/v1/orders/1?total_price=250.50", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.orderController.UpdateOrderTotalPrice(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusCreated, rec.Code)
	suite.mockOrderService.AssertExpectations(suite.T())
}

func (suite *OrderControllerTestSuite) TestUpdateOrderTotalPrice_InvalidPrice() {
	req := httptest.NewRequest(http.MethodPut, "/api/v1/orders/1?total_price=invalid", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	err := suite.orderController.UpdateOrderTotalPrice(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, rec.Code)
	suite.mockOrderService.AssertNotCalled(suite.T(), "UpdateOrderTotalPrice")
}

func (suite *OrderControllerTestSuite) TestGetOrdersByStatus_Success() {
	expectedOrders := []domain.Order{
		{
			Id:         1,
			UserId:     1,
			TotalPrice: 199.99,
			Status:     true,
		},
	}

	suite.mockOrderService.On("GetOrdersByStatus", "true").Return(expectedOrders, nil)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/orders/?status=true", nil)
	rec := httptest.NewRecorder()
	ctx := suite.echo.NewContext(req, rec)

	err := suite.orderController.GetOrdersByStatus(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, rec.Code)

	var response []domain.Order
	json.Unmarshal(rec.Body.Bytes(), &response)

	assert.Len(suite.T(), response, 1)
	assert.Equal(suite.T(), true, response[0].Status)
	suite.mockOrderService.AssertExpectations(suite.T())
}

func TestOrderControllerTestSuite(t *testing.T) {
	suite.Run(t, new(OrderControllerTestSuite))
}

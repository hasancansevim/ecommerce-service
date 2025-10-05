package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/controller/response"
	"go-ecommerce-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type OrderController struct {
	orderService service.IOrderService
}

func NewOrderController(orderService service.IOrderService) *OrderController {
	return &OrderController{orderService: orderService}
}

func (orderController *OrderController) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/v1/orders", orderController.CreateOrder)
	e.GET("/api/v1/orders/:id", orderController.GetOrderById)
	e.GET("/api/v1/orders/get-orders-by-user-id", orderController.GetOrdersByUserId)
	e.GET("/api/v1/orders/get-all-orders", orderController.GetAllOrders)
	e.PUT("/api/v1/orders/update-order-status/:id", orderController.UpdateOrderStatus)
	e.DELETE("/api/v1/orders/:id", orderController.DeleteOrderById)
	e.PUT("/api/v1/orders/:id", orderController.UpdateOrderTotalPrice)
	e.GET("/api/v1/orders/", orderController.GetOrdersByStatus)
	e.GET("/api/v1/orders/get-order-count-by-user-id/", orderController.GetOrderCountByUserId) // !
}

func (orderController *OrderController) CreateOrder(c echo.Context) error {
	var addOrderRequest request.AddOrderRequest
	bindErr := c.Bind(&addOrderRequest)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	createOrderErr := orderController.orderService.CreateOrder(addOrderRequest.ToModel())
	if createOrderErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: createOrderErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (orderController *OrderController) GetOrderById(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	getOrderById := orderController.orderService.GetOrderById(int64(id))
	return c.JSON(http.StatusOK, response.ToResponseOrderData(getOrderById))
}

func (orderController *OrderController) GetOrdersByUserId(c echo.Context) error {
	queryParam := c.QueryParam("user_id")
	user_id, convertErr := strconv.Atoi(queryParam)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	getOrdersByUserId, getOrdersByUserIdErr := orderController.orderService.GetOrdersByUserId(int64(user_id))
	if getOrdersByUserIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: getOrdersByUserIdErr.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponseListOrders(getOrdersByUserId))
}

func (orderController *OrderController) GetAllOrders(c echo.Context) error {
	orders, getAllOrdersErr := orderController.orderService.GetAllOrders()
	if getAllOrdersErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: getAllOrdersErr.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponseListOrders(orders))
}

func (orderController *OrderController) UpdateOrderStatus(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	queryParam := c.QueryParam("status")
	status := strconv.CanBackquote(queryParam)
	updateOrderStatusErr := orderController.orderService.UpdateOrderStatus(int64(id), status)
	if updateOrderStatusErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: updateOrderStatusErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (orderController *OrderController) DeleteOrderById(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	deleteOrderByIdErr := orderController.orderService.DeleteOrderById(int64(id))
	if deleteOrderByIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: deleteOrderByIdErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (orderController *OrderController) UpdateOrderTotalPrice(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	queryParam := c.QueryParam("total_price")
	totalPrice, parseFloatErr := strconv.ParseFloat(queryParam, 64)
	if parseFloatErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: parseFloatErr.Error(),
		})
	}
	updateOrderTotalPriceErr := orderController.orderService.UpdateOrderTotalPrice(int64(id), float32(totalPrice))
	if updateOrderTotalPriceErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: updateOrderTotalPriceErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (orderController *OrderController) GetOrdersByStatus(c echo.Context) error {
	status := c.QueryParam("status")
	log.Info(status)
	ordersByStatus, ordersByStatusErr := orderController.orderService.GetOrdersByStatus(status)
	if ordersByStatusErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: ordersByStatusErr.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponseListOrders(ordersByStatus))
}

func (orderController *OrderController) GetOrderCountByUserId(c echo.Context) error {
	queryParam := c.QueryParam("user_id")
	user_id, convertErr := strconv.Atoi(queryParam)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	orderCountByUserId, orderCountByUserIdErr := orderController.orderService.GetOrderCountByUserId(int64(user_id))
	if orderCountByUserIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: orderCountByUserIdErr.Error(),
		})
	}
	return c.JSON(http.StatusOK, orderCountByUserId)
}

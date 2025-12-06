package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderController struct {
	orderService service.IOrderService
	BaseController
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
}

func (orderController *OrderController) CreateOrder(c echo.Context) error {
	var addOrderRequest request.AddOrderRequest
	bindErr := c.Bind(&addOrderRequest)
	if bindErr != nil {
		return bindErr
	}

	createdOrder, serviceErr := orderController.orderService.CreateOrder(addOrderRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}
	return orderController.Success(c, createdOrder, "Sipariş Oluşturuldu")
}

func (orderController *OrderController) GetOrderById(c echo.Context) error {
	id, parseIdErr := orderController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	getOrderById := orderController.orderService.GetOrderById(id)
	return orderController.Success(c, getOrderById, "")
}

func (orderController *OrderController) GetOrdersByUserId(c echo.Context) error {
	userId, parseIdErr := orderController.ParseIdParam(c, "user_id")
	if parseIdErr != nil {
		return parseIdErr
	}

	getOrdersByUserId, serviceErr := orderController.orderService.GetOrdersByUserId(userId)
	if serviceErr != nil {
		return serviceErr
	}
	return orderController.Success(c, getOrdersByUserId, "")
}

func (orderController *OrderController) GetAllOrders(c echo.Context) error {
	orders, serviceErr := orderController.orderService.GetAllOrders()
	if serviceErr != nil {
		return serviceErr
	}
	return orderController.Success(c, orders, "")
}

func (orderController *OrderController) UpdateOrderStatus(c echo.Context) error {
	id, parseIdErr := orderController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	queryParam := orderController.StringQueryParam(c, "status")
	status := queryParam == "true"
	updatedOrder, serviceErr := orderController.orderService.UpdateOrderStatus(id, status)
	if serviceErr != nil {
		return serviceErr
	}
	return orderController.Success(c, updatedOrder, "Sipariş Durumu Güncellendi")
}

func (orderController *OrderController) DeleteOrderById(c echo.Context) error {
	id, parseIdErr := orderController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	if serviceErr := orderController.orderService.DeleteOrderById(id); serviceErr != nil {
		return serviceErr
	}
	return orderController.Created(c, nil, "Sipariş Silinid")
}

func (orderController *OrderController) UpdateOrderTotalPrice(c echo.Context) error {
	id, parseIdErr := orderController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	queryParam := orderController.StringQueryParam(c, "total_price")
	totalPrice, parseFloatErr := strconv.ParseFloat(queryParam, 64)
	if parseFloatErr != nil {
		return parseFloatErr
	}

	updatedOrder, serviceErr := orderController.orderService.UpdateOrderTotalPrice(id, float32(totalPrice))
	if serviceErr != nil {
		return serviceErr
	}
	return orderController.Success(c, updatedOrder, "Sipariş Tutarı Güncellendi")
}

func (orderController *OrderController) GetOrdersByStatus(c echo.Context) error {
	status := orderController.StringQueryParam(c, "status")
	ordersByStatus, serviceErr := orderController.orderService.GetOrdersByStatus(status)
	if serviceErr != nil {
		return serviceErr
	}

	return orderController.Success(c, ordersByStatus, "")
}

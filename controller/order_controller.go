package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/controller/response"
	"go-ecommerce-service/pkg/validation"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/model"
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
		return orderController.BadRequest(c, bindErr)
	}

	validator := validation.OrderCreateValidator{OrderReq: model.OrderCreate{
		UserId:     addOrderRequest.UserId,
		TotalPrice: addOrderRequest.TotalPrice,
		Status:     addOrderRequest.Status,
		CreatedAt:  addOrderRequest.CreatedAt,
	}}
	if validationErr := validator.Validate(); validationErr != nil {
		return orderController.BadRequest(c, validationErr)
	}
	if createOrderErr := orderController.orderService.CreateOrder(addOrderRequest.ToModel()); createOrderErr != nil {
		return orderController.BadRequest(c, createOrderErr)
	}
	return orderController.Created(c)
}

func (orderController *OrderController) GetOrderById(c echo.Context) error {
	id, err := orderController.ParseIdParam(c, "id")
	if err != nil {
		return orderController.BadRequest(c, err)
	}
	getOrderById := orderController.orderService.GetOrderById(id)
	return orderController.Success(c, getOrderById)
}

func (orderController *OrderController) GetOrdersByUserId(c echo.Context) error {
	userId, err := orderController.ParseIdParam(c, "user_id")
	if err != nil {
		return orderController.BadRequest(c, err)
	}
	getOrdersByUserId, getOrdersByUserIdErr := orderController.orderService.GetOrdersByUserId(userId)
	if getOrdersByUserIdErr != nil {
		return orderController.BadRequest(c, getOrdersByUserIdErr)
	}
	return orderController.Success(c, response.ToResponseListOrders(getOrdersByUserId))
}

func (orderController *OrderController) GetAllOrders(c echo.Context) error {
	orders, getAllOrdersErr := orderController.orderService.GetAllOrders()
	if getAllOrdersErr != nil {
		return orderController.BadRequest(c, getAllOrdersErr)
	}
	return orderController.Success(c, response.ToResponseListOrders(orders))
}

func (orderController *OrderController) UpdateOrderStatus(c echo.Context) error {
	id, err := orderController.ParseIdParam(c, "id")
	if err != nil {
		return orderController.BadRequest(c, err)
	}
	queryParam := orderController.StringQueryParam(c, "status")
	status := queryParam == "true"
	if updateOrderStatusErr := orderController.orderService.UpdateOrderStatus(id, status); updateOrderStatusErr != nil {
		return orderController.BadRequest(c, updateOrderStatusErr)
	}
	return orderController.Created(c)
}

func (orderController *OrderController) DeleteOrderById(c echo.Context) error {
	id, err := orderController.ParseIdParam(c, "id")
	if err != nil {
		return orderController.BadRequest(c, err)
	}
	if deleteOrderByIdErr := orderController.orderService.DeleteOrderById(id); deleteOrderByIdErr != nil {
		return orderController.BadRequest(c, deleteOrderByIdErr)
	}
	return orderController.Created(c)
}

func (orderController *OrderController) UpdateOrderTotalPrice(c echo.Context) error {
	id, err := orderController.ParseIdParam(c, "id")
	if err != nil {
		return orderController.BadRequest(c, err)
	}
	queryParam := orderController.StringQueryParam(c, "total_price")
	totalPrice, parseFloatErr := strconv.ParseFloat(queryParam, 64)
	if parseFloatErr != nil {
		return orderController.BadRequest(c, parseFloatErr)
	}
	updateOrderTotalPriceErr := orderController.orderService.UpdateOrderTotalPrice(id, float32(totalPrice))
	if updateOrderTotalPriceErr != nil {
		return orderController.BadRequest(c, updateOrderTotalPriceErr)
	}
	return orderController.Created(c)
}

func (orderController *OrderController) GetOrdersByStatus(c echo.Context) error {
	status := orderController.StringQueryParam(c, "status")
	ordersByStatus, ordersByStatusErr := orderController.orderService.GetOrdersByStatus(status)
	if ordersByStatusErr != nil {
		return orderController.BadRequest(c, ordersByStatusErr)
	}
	return orderController.Success(c, response.ToResponseListOrders(ordersByStatus))
}

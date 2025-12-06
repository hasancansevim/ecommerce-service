package controller

import (
	"go-ecommerce-service/controller/request"
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
	createdOrder, serviceErr := orderController.orderService.CreateOrder(addOrderRequest.ToModel())
	if serviceErr != nil {
		return orderController.BadRequest(c, serviceErr)
	}
	return orderController.Success(c, createdOrder, "Sipariş Oluşturuldu")
}

func (orderController *OrderController) GetOrderById(c echo.Context) error {
	id, err := orderController.ParseIdParam(c, "id")
	if err != nil {
		return orderController.BadRequest(c, err)
	}
	getOrderById := orderController.orderService.GetOrderById(id)
	return orderController.Success(c, getOrderById, "")
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
	return orderController.Success(c, getOrdersByUserId, "")
}

func (orderController *OrderController) GetAllOrders(c echo.Context) error {
	orders, getAllOrdersErr := orderController.orderService.GetAllOrders()
	if getAllOrdersErr != nil {
		return orderController.BadRequest(c, getAllOrdersErr)
	}
	return orderController.Success(c, orders, "")
}

func (orderController *OrderController) UpdateOrderStatus(c echo.Context) error {
	id, parseIdErr := orderController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return orderController.BadRequest(c, parseIdErr)
	}

	queryParam := orderController.StringQueryParam(c, "status")
	status := queryParam == "true"
	updatedOrder, serviceErr := orderController.orderService.UpdateOrderStatus(id, status)
	if serviceErr != nil {
		return orderController.BadRequest(c, serviceErr)
	}
	return orderController.Success(c, updatedOrder, "Sipariş Durumu Güncellendi")
}

func (orderController *OrderController) DeleteOrderById(c echo.Context) error {
	id, err := orderController.ParseIdParam(c, "id")
	if err != nil {
		return orderController.BadRequest(c, err)
	}
	if deleteOrderByIdErr := orderController.orderService.DeleteOrderById(id); deleteOrderByIdErr != nil {
		return orderController.BadRequest(c, deleteOrderByIdErr)
	}
	return orderController.Created(c, nil, "Sipariş Silinid")
}

func (orderController *OrderController) UpdateOrderTotalPrice(c echo.Context) error {
	id, parseIdErr := orderController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return orderController.BadRequest(c, parseIdErr)
	}

	queryParam := orderController.StringQueryParam(c, "total_price")
	totalPrice, parseFloatErr := strconv.ParseFloat(queryParam, 64)
	if parseFloatErr != nil {
		return orderController.BadRequest(c, parseFloatErr)
	}

	updatedOrder, serviceErr := orderController.orderService.UpdateOrderTotalPrice(id, float32(totalPrice))
	if serviceErr != nil {
		return orderController.BadRequest(c, serviceErr)
	}
	return orderController.Success(c, updatedOrder, "Sipariş Tutarı Güncellendi")
}

func (orderController *OrderController) GetOrdersByStatus(c echo.Context) error {
	status := orderController.StringQueryParam(c, "status")
	ordersByStatus, ordersByStatusErr := orderController.orderService.GetOrdersByStatus(status)
	if ordersByStatusErr != nil {
		return orderController.BadRequest(c, ordersByStatusErr)
	}
	return orderController.Success(c, ordersByStatus, "")
}

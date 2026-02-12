package controller

import (
	"errors"
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderItemController struct {
	orderItemService service.IOrderItemService
	BaseController
}

func NewOrderItemController(orderItemService service.IOrderItemService) *OrderItemController {
	return &OrderItemController{
		orderItemService: orderItemService,
	}
}

func (orderItemController *OrderItemController) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/v1/order-items", orderItemController.AddOrderItem)
	e.GET("/api/v1/order-items/:id", orderItemController.GetOrderItemById)
	e.GET("/api/v1/order-items", orderItemController.GetOrderItems)
	e.PUT("/api/v1/order-items/:id", orderItemController.UpdateOrderItem)
	e.PUT("/api/v1/order-items/update-order-item-quantity/:id", orderItemController.UpdateOrderItemQuantity)
	e.DELETE("/api/v1/order-items/:id", orderItemController.DeleteOrderItemById)
	e.DELETE("/api/v1/order-items", orderItemController.DeleteAllOrderItemsByOrderId)
}

func (orderItemController *OrderItemController) AddOrderItem(c echo.Context) error {
	var addOrderItemRequest request.AddOrderItemRequest
	bindErr := c.Bind(&addOrderItemRequest)
	if bindErr != nil {
		return bindErr
	}

	addedOrderItem, serviceErr := orderItemController.orderItemService.AddOrderItem(addOrderItemRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}
	return orderItemController.Success(c, addedOrderItem, "Order item added")
}

func (orderItemController *OrderItemController) GetOrderItemById(c echo.Context) error {
	id, parseIdErr := orderItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	getOrderItem, serviceErr := orderItemController.orderItemService.GetOrderItemById(id)
	if serviceErr != nil {
		return serviceErr
	}

	return orderItemController.Success(c, getOrderItem, "Order item retrieved")
}

func (orderItemController *OrderItemController) GetOrderItems(c echo.Context) error {
	orderId := orderItemController.StringQueryParam(c, "order_id")
	productId := orderItemController.StringQueryParam(c, "product_id")

	if orderId != "" {
		orderId, convertErr := strconv.Atoi(orderId)
		if convertErr != nil {
			return convertErr
		}
		getOrderItemsByOrderId, serviceErr := orderItemController.orderItemService.GetOrderItemsByOrderId(int64(orderId))
		if serviceErr != nil {
			return serviceErr
		}
		return orderItemController.Success(c, getOrderItemsByOrderId, "Order items for order retrieved")
	}
	if productId != "" {
		productId, convertErr := strconv.Atoi(productId)
		if convertErr != nil {
			return convertErr
		}
		getOrderItemsByProductId, serviceErr := orderItemController.orderItemService.GetOrderItemsByProductId(int64(productId))
		if serviceErr != nil {
			return serviceErr
		}
		return orderItemController.Success(c, getOrderItemsByProductId, "")
	}
	return orderItemController.BadRequest(c, errors.New("order_id or product_id parameters required"))
}

func (orderItemController *OrderItemController) UpdateOrderItem(c echo.Context) error {
	var updateOrderItemRequest request.UpdateOrderItemRequest
	bindErr := c.Bind(&updateOrderItemRequest)
	if bindErr != nil {
		return bindErr
	}

	id, parseIdErr := orderItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	updatedOrderItem, serviceErr := orderItemController.orderItemService.UpdateOrderItem(id, updateOrderItemRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}
	return orderItemController.Created(c, updatedOrderItem, "Order item updated")
}

func (orderItemController *OrderItemController) UpdateOrderItemQuantity(c echo.Context) error {
	id, parseIdErr := orderItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	queryParam := orderItemController.StringQueryParam(c, "new_quantity")
	newQuantity, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		return queryParamErr
	}

	updatedOrderItem, serviceErr := orderItemController.orderItemService.UpdateOrderItemQuantity(id, newQuantity)
	if serviceErr != nil {
		return serviceErr
	}
	return orderItemController.Created(c, updatedOrderItem, "Order item quantity increased")
}

func (orderItemController *OrderItemController) DeleteOrderItemById(c echo.Context) error {
	id, parseIdErr := orderItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	if serviceErr := orderItemController.orderItemService.DeleteOrderItemById(id); serviceErr != nil {
		return serviceErr
	}
	return orderItemController.Created(c, nil, "Order item deleted")
}

func (orderItemController *OrderItemController) DeleteAllOrderItemsByOrderId(c echo.Context) error {
	queryParam := orderItemController.StringQueryParam(c, "order_id")
	orderId, paramErr := strconv.Atoi(queryParam)
	if paramErr != nil {
		return paramErr
	}

	if serviceErr := orderItemController.orderItemService.DeleteAllOrderItemsByOrderId(int64(orderId)); serviceErr != nil {
		return serviceErr
	}
	return orderItemController.Created(c, nil, "All order items for order deleted")
}

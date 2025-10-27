package controller

import (
	"errors"
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/controller/response"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/model"
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
		return orderItemController.BadRequest(c, bindErr)
	}

	if addOrderItemErr := orderItemController.orderItemService.AddOrderItem(model.OrderItemCreate{
		OrderId:   addOrderItemRequest.OrderId,
		ProductId: addOrderItemRequest.ProductId,
		Quantity:  addOrderItemRequest.Quantity,
		Price:     addOrderItemRequest.Price,
	}); addOrderItemErr != nil {
		return orderItemController.BadRequest(c, addOrderItemErr)
	}
	return orderItemController.Created(c)
}

func (orderItemController *OrderItemController) GetOrderItemById(c echo.Context) error {
	id, err := orderItemController.ParseIdParam(c, "id")
	if err != nil {
		return orderItemController.BadRequest(c, err)
	}
	getOrderItemById, getOrderItemByIdErr := orderItemController.orderItemService.GetOrderItemById(id)
	if getOrderItemByIdErr != nil {
		return orderItemController.BadRequest(c, getOrderItemByIdErr)
	}
	return orderItemController.Success(c, response.ToResponseOrderItemData(getOrderItemById))
}

func (orderItemController *OrderItemController) GetOrderItems(c echo.Context) error {
	orderId := orderItemController.StringQueryParam(c, "order_id")
	productId := orderItemController.StringQueryParam(c, "product_id")

	if orderId != "" {
		orderId, convertErr := strconv.Atoi(orderId)
		if convertErr != nil {
			return orderItemController.BadRequest(c, convertErr)
		}
		getOrderItemsByOrderId, getOrderItemsByOrderIdErr := orderItemController.orderItemService.GetOrderItemsByOrderId(int64(orderId))
		if getOrderItemsByOrderIdErr != nil {
			return orderItemController.BadRequest(c, getOrderItemsByOrderIdErr)
		}
		return orderItemController.Success(c, response.ToResponseListOrderItems(getOrderItemsByOrderId))
	}
	if productId != "" {
		productId, convertErr := strconv.Atoi(productId)
		if convertErr != nil {
			return orderItemController.BadRequest(c, convertErr)
		}
		getOrderItemsByProductId, getOrderItemsByProductIdErr := orderItemController.orderItemService.GetOrderItemsByProductId(int64(productId))
		if getOrderItemsByProductIdErr != nil {
			return orderItemController.BadRequest(c, getOrderItemsByProductIdErr)
		}
		return orderItemController.Success(c, response.ToResponseListOrderItems(getOrderItemsByProductId))
	}
	return orderItemController.BadRequest(c, errors.New("order_id or product_id parameters required"))
}

func (orderItemController *OrderItemController) UpdateOrderItem(c echo.Context) error {
	var updateOrderItemRequest request.UpdateOrderItemRequest
	bindErr := c.Bind(&updateOrderItemRequest)
	if bindErr != nil {
		return orderItemController.BadRequest(c, bindErr)
	}

	id, parseIdParamErr := orderItemController.ParseIdParam(c, "id")
	if parseIdParamErr != nil {
		return orderItemController.BadRequest(c, parseIdParamErr)
	}

	if updateOrderItemErr := orderItemController.orderItemService.UpdateOrderItem(id, domain.OrderItem{
		OrderId:   updateOrderItemRequest.OrderId,
		ProductId: updateOrderItemRequest.ProductId,
		Quantity:  updateOrderItemRequest.Quantity,
		Price:     updateOrderItemRequest.Price,
	}); updateOrderItemErr != nil {
		return orderItemController.BadRequest(c, updateOrderItemErr)
	}

	return orderItemController.Created(c)
}

func (orderItemController *OrderItemController) UpdateOrderItemQuantity(c echo.Context) error {
	id, err := orderItemController.ParseIdParam(c, "id")
	if err != nil {
		return orderItemController.BadRequest(c, err)
	}
	queryParam := orderItemController.StringQueryParam(c, "new_quantity")
	newQuantity, queryParamConvertErr := strconv.Atoi(queryParam)
	if queryParamConvertErr != nil {
		return orderItemController.BadRequest(c, queryParamConvertErr)
	}

	if updateOrderItemQuantityErr := orderItemController.orderItemService.UpdateOrderItemQuantity(id, newQuantity); updateOrderItemQuantityErr != nil {
		return orderItemController.BadRequest(c, updateOrderItemQuantityErr)
	}
	return orderItemController.Created(c)
}

func (orderItemController *OrderItemController) DeleteOrderItemById(c echo.Context) error {
	id, err := orderItemController.ParseIdParam(c, "id")
	if err != nil {
		return orderItemController.BadRequest(c, err)
	}

	if deleteOrderItemByIdErr := orderItemController.orderItemService.DeleteOrderItemById(id); deleteOrderItemByIdErr != nil {
		return orderItemController.BadRequest(c, deleteOrderItemByIdErr)
	}
	return orderItemController.Created(c)
}

func (orderItemController *OrderItemController) DeleteAllOrderItemsByOrderId(c echo.Context) error {
	queryParam := orderItemController.StringQueryParam(c, "order_id")
	orderId, err := strconv.Atoi(queryParam)
	if err != nil {
		return orderItemController.BadRequest(c, err)
	}
	if deleteAllOrderItemsByOrderIdErr := orderItemController.orderItemService.DeleteAllOrderItemsByOrderId(int64(orderId)); deleteAllOrderItemsByOrderIdErr != nil {
		return orderItemController.BadRequest(c, deleteAllOrderItemsByOrderIdErr)
	}
	return orderItemController.Created(c)
}

package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/controller/response"
	"go-ecommerce-service/domain"
	"go-ecommerce-service/service"
	"go-ecommerce-service/service/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderItemController struct {
	orderItemService service.IOrderItemService
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
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	addOrderItemErr := orderItemController.orderItemService.AddOrderItem(model.OrderItemCreate{
		OrderId:   addOrderItemRequest.OrderId,
		ProductId: addOrderItemRequest.ProductId,
		Quantity:  addOrderItemRequest.Quantity,
		Price:     addOrderItemRequest.Price,
	})
	if addOrderItemErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: addOrderItemErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (orderItemController *OrderItemController) GetOrderItemById(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	getOrderItemById, getOrderItemByIdErr := orderItemController.orderItemService.GetOrderItemById(int64(id))
	if getOrderItemByIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: getOrderItemByIdErr.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.ToResponseOrderItemData(getOrderItemById))
}

func (orderItemController *OrderItemController) GetOrderItems(c echo.Context) error {
	queryParam_orderId := c.QueryParam("order_id")
	queryParam_productId := c.QueryParam("product_id")
	if queryParam_orderId != "" {
		orderId, convertErr := strconv.Atoi(queryParam_orderId)
		if convertErr != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				ErrorDescription: convertErr.Error(),
			})
		}
		getOrderItemsByOrderId, getOrderItemsByOrderIdErr := orderItemController.orderItemService.GetOrderItemsByOrderId(int64(orderId))
		if getOrderItemsByOrderIdErr != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				ErrorDescription: getOrderItemsByOrderIdErr.Error(),
			})
		}
		return c.JSON(http.StatusOK, response.ToResponseListOrderItems(getOrderItemsByOrderId))
	}
	if queryParam_productId != "" {
		productId, convertErr := strconv.Atoi(queryParam_productId)
		if convertErr != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				ErrorDescription: convertErr.Error(),
			})
		}
		getOrderItemsByProductId, getOrderItemsByProductIdErr := orderItemController.orderItemService.GetOrderItemsByProductId(int64(productId))
		if getOrderItemsByProductIdErr != nil {
			return c.JSON(http.StatusBadRequest, response.ErrorResponse{
				ErrorDescription: getOrderItemsByProductIdErr.Error(),
			})
		}
		return c.JSON(http.StatusOK, response.ToResponseListOrderItems(getOrderItemsByProductId))
	}
	return c.JSON(http.StatusBadRequest, response.ErrorResponse{
		ErrorDescription: "order_id or product_id parameters required",
	})
}

func (orderItemController *OrderItemController) UpdateOrderItem(c echo.Context) error {
	var updateOrderItemRequest request.UpdateOrderItemRequest
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	bindErr := c.Bind(&updateOrderItemRequest)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	updateOrderItemErr := orderItemController.orderItemService.UpdateOrderItem(int64(id), domain.OrderItem{
		OrderId:   updateOrderItemRequest.OrderId,
		ProductId: updateOrderItemRequest.ProductId,
		Quantity:  updateOrderItemRequest.Quantity,
		Price:     updateOrderItemRequest.Price,
	})

	if updateOrderItemErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: updateOrderItemErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (orderItemController *OrderItemController) UpdateOrderItemQuantity(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	queryParam := c.QueryParam("new_quantity")
	new_quantity, queryParamConvertErr := strconv.Atoi(queryParam)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	if queryParamConvertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: queryParamConvertErr.Error(),
		})
	}

	updateOrderItemQuantityErr := orderItemController.orderItemService.UpdateOrderItemQuantity(int64(id), new_quantity)
	if updateOrderItemQuantityErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: updateOrderItemQuantityErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (orderItemController *OrderItemController) DeleteOrderItemById(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	deleteOrderItemByIdErr := orderItemController.orderItemService.DeleteOrderItemById(int64(id))
	if deleteOrderItemByIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: deleteOrderItemByIdErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (orderItemController *OrderItemController) DeleteAllOrderItemsByOrderId(c echo.Context) error {
	queryParam := c.QueryParam("order_id")
	order_id, convertErr := strconv.Atoi(queryParam)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	deleteAllOrderItemsByOrderIdErr := orderItemController.orderItemService.DeleteAllOrderItemsByOrderId(int64(order_id))
	if deleteAllOrderItemsByOrderIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: deleteAllOrderItemsByOrderIdErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

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

type CartItemController struct {
	cartItemService service.ICartItemService
}

func NewCartItemController(cartItemService service.ICartItemService) *CartItemController {
	return &CartItemController{cartItemService: cartItemService}
}

func (cartItemController *CartItemController) RegiesterRoutes(e *echo.Echo) {
	e.GET("/api/v1/cart_items/:id", cartItemController.GetItemsByCartId)
	e.POST("/api/v1/cart_items/", cartItemController.AddItemToCart)
	e.PUT("/api/v1/cart_items/:id", cartItemController.UpdateItemQuantity)
	e.DELETE("/api/v1/cart_items/:id", cartItemController.RemoveItemFromCart)
	e.DELETE("/api/v1/cart_items/", cartItemController.ClearCartItems)
	e.PUT("/api/v1/cart_items/increase/:id", cartItemController.IncreaseItemQuantity)
	e.PUT("/api/v1/cart_items/decrease/:id", cartItemController.DecreaseItemQuantity)
}

func (cartItemController *CartItemController) GetItemsByCartId(c echo.Context) error {
	param := c.Param("id")
	cartId, convertIdErr := strconv.Atoi(param)

	if convertIdErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertIdErr.Error(),
		})
	}
	cart_items := cartItemController.cartItemService.GetItemsByCartId(int64(cartId))
	return c.JSON(http.StatusOK, response.ToResponseListCartItems(cart_items))
}

func (cartItemController *CartItemController) AddItemToCart(c echo.Context) error {
	var addCartItemRequest request.AddCartItemRequest
	bindErr := c.Bind(&addCartItemRequest)

	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}
	addCartItemErr := cartItemController.cartItemService.AddItemToCart(addCartItemRequest.ToModel())
	if addCartItemErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: addCartItemErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (cartItemController *CartItemController) UpdateItemQuantity(c echo.Context) error {
	param := c.Param("id")
	id, paramConvertErr := strconv.Atoi(param)

	if paramConvertErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: paramConvertErr.Error(),
		})
	}

	queryParam := c.QueryParam("newQuantity")
	newQuantity, queryParamConvertErr := strconv.Atoi(queryParam)
	log.Info(int64(id))
	log.Info(newQuantity)
	if queryParamConvertErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: queryParamConvertErr.Error(),
		})
	}
	updateQuantityErr := cartItemController.cartItemService.UpdateItemQuantity(int64(id), newQuantity)
	if updateQuantityErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: updateQuantityErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (cartItemController *CartItemController) RemoveItemFromCart(c echo.Context) error {
	param := c.Param("id")
	id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	removeItemFromCartErr := cartItemController.cartItemService.RemoveItemFromCart(int64(id))
	if removeItemFromCartErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: removeItemFromCartErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (cartItemController *CartItemController) ClearCartItems(c echo.Context) error {
	param := c.QueryParam("cart_id")
	cart_id, convertErr := strconv.Atoi(param)
	if convertErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}
	clearCartItemsErr := cartItemController.cartItemService.ClearCartItems(int64(cart_id))
	if clearCartItemsErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: clearCartItemsErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (cartItemController *CartItemController) IncreaseItemQuantity(c echo.Context) error {
	param := c.Param("id")
	id, paramConvertErr := strconv.Atoi(param)
	if paramConvertErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: paramConvertErr.Error(),
		})
	}

	queryParam := c.QueryParam("amount")
	amount, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: queryParamErr.Error(),
		})
	}

	increaseItemQuantityErr := cartItemController.cartItemService.IncreaseItemQuantity(int64(id), amount)
	if increaseItemQuantityErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: increaseItemQuantityErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (cartItemController *CartItemController) DecreaseItemQuantity(c echo.Context) error {
	param := c.Param("id")
	id, paramConvertErr := strconv.Atoi(param)
	if paramConvertErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: paramConvertErr.Error(),
		})
	}

	queryParam := c.QueryParam("amount")
	amount, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: queryParamErr.Error(),
		})
	}

	increaseItemQuantityErr := cartItemController.cartItemService.DecreaseItemQuantity(int64(id), amount)
	if increaseItemQuantityErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: increaseItemQuantityErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

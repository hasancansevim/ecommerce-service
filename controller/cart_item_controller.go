package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartItemController struct {
	cartItemService service.ICartItemService
	BaseController
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
	cartId, parseIdErr := cartItemController.BaseController.ParseIdParam(c, "cart_id")
	if parseIdErr != nil {
		return parseIdErr
	}
	cartItems := cartItemController.cartItemService.GetItemsByCartId(cartId)
	return cartItemController.Success(c, cartItems, "")
}

func (cartItemController *CartItemController) AddItemToCart(c echo.Context) error {
	var addCartItemRequest request.AddCartItemRequest
	bindErr := c.Bind(&addCartItemRequest)

	if bindErr != nil {
		return bindErr
	}
	cartItem, serviceErr := cartItemController.cartItemService.AddItemToCart(addCartItemRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}

	return cartItemController.Created(c, cartItem, "")
}

func (cartItemController *CartItemController) UpdateItemQuantity(c echo.Context) error {
	id, parseIdErr := cartItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	queryParam := cartItemController.StringQueryParam(c, "newQuantity")
	newQuantity, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		return queryParamErr
	}

	cartItem, serviceErr := cartItemController.cartItemService.UpdateItemQuantity(id, newQuantity)
	if serviceErr != nil {
		return queryParamErr
	}
	return cartItemController.Created(c, cartItem, "Cart item updated")
}

func (cartItemController *CartItemController) RemoveItemFromCart(c echo.Context) error {
	id, parseIdErr := cartItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}
	if serviceErr := cartItemController.cartItemService.RemoveItemFromCart(id); serviceErr != nil {
		return serviceErr
	}
	return cartItemController.Created(c, nil, "Product removed from cart")
}

func (cartItemController *CartItemController) ClearCartItems(c echo.Context) error {
	cartId, parseIdErr := cartItemController.ParseIdParam(c, "cart_id")
	if parseIdErr != nil {
		return parseIdErr
	}
	if serviceErr := cartItemController.cartItemService.ClearCartItems(cartId); serviceErr != nil {
		return serviceErr
	}
	return cartItemController.Created(c, nil, "Cart items cleared")
}

func (cartItemController *CartItemController) IncreaseItemQuantity(c echo.Context) error {
	id, parseIdErr := cartItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}
	queryParam := cartItemController.StringQueryParam(c, "amount")
	amount, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		return queryParamErr
	}

	if serviceErr := cartItemController.cartItemService.IncreaseItemQuantity(id, amount); serviceErr != nil {
		return serviceErr
	}
	return cartItemController.Created(c, nil, "")
}

func (cartItemController *CartItemController) DecreaseItemQuantity(c echo.Context) error {
	id, parseIdErr := cartItemController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}
	queryParam := cartItemController.StringQueryParam(c, "amount")
	amount, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		return queryParamErr
	}

	if serviceErr := cartItemController.cartItemService.DecreaseItemQuantity(id, amount); serviceErr != nil {
		return serviceErr
	}
	return cartItemController.Created(c, nil, "")
}

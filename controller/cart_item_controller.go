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
	cartId, err := cartItemController.BaseController.ParseIdParam(c, "cart_id")
	if err != nil {
		return cartItemController.BadRequest(c, err)
	}
	cartItems := cartItemController.cartItemService.GetItemsByCartId(cartId)
	return cartItemController.Success(c, cartItems, "")
}

func (cartItemController *CartItemController) AddItemToCart(c echo.Context) error {
	var addCartItemRequest request.AddCartItemRequest
	bindErr := c.Bind(&addCartItemRequest)

	if bindErr != nil {
		return cartItemController.BadRequest(c, bindErr)
	}
	if addCartItemErr := cartItemController.cartItemService.AddItemToCart(addCartItemRequest.ToModel()); addCartItemErr != nil {
		return cartItemController.BadRequest(c, addCartItemErr)
	}
	return cartItemController.Created(c, addCartItemRequest, "")
}

func (cartItemController *CartItemController) UpdateItemQuantity(c echo.Context) error {
	id, err := cartItemController.ParseIdParam(c, "id")
	if err != nil {
		return cartItemController.BadRequest(c, err)
	}
	queryParam := cartItemController.StringQueryParam(c, "newQuantity")
	newQuantity, queryParamConvertErr := strconv.Atoi(queryParam)
	if queryParamConvertErr != nil {
		return cartItemController.BadRequest(c, queryParamConvertErr)
	}
	if updateQuantityErr := cartItemController.cartItemService.UpdateItemQuantity(id, newQuantity); updateQuantityErr != nil {
		return cartItemController.BadRequest(c, updateQuantityErr)
	}
	return cartItemController.Created(c, nil, "Sepet Öğesi Güncellendi")
}

func (cartItemController *CartItemController) RemoveItemFromCart(c echo.Context) error {
	id, err := cartItemController.ParseIdParam(c, "id")
	if err != nil {
		return cartItemController.BadRequest(c, err)
	}
	if removeItemFromCartErr := cartItemController.cartItemService.RemoveItemFromCart(id); removeItemFromCartErr != nil {
		return cartItemController.BadRequest(c, removeItemFromCartErr)
	}
	return cartItemController.Created(c, nil, "Sepetten Ürün Silindi")
}

func (cartItemController *CartItemController) ClearCartItems(c echo.Context) error {
	cartId, err := cartItemController.ParseIdParam(c, "cart_id")
	if err != nil {
		return cartItemController.BadRequest(c, err)
	}
	if clearCartItemsErr := cartItemController.cartItemService.ClearCartItems(cartId); clearCartItemsErr != nil {
		return cartItemController.BadRequest(c, clearCartItemsErr)
	}
	return cartItemController.Created(c, nil, "Sepetteki Ürünler Silindi")
}

func (cartItemController *CartItemController) IncreaseItemQuantity(c echo.Context) error {
	id, err := cartItemController.ParseIdParam(c, "id")
	if err != nil {
		return cartItemController.BadRequest(c, err)
	}
	queryParam := cartItemController.StringQueryParam(c, "amount")
	amount, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		return cartItemController.BadRequest(c, queryParamErr)
	}

	if increaseItemQuantityErr := cartItemController.cartItemService.IncreaseItemQuantity(id, amount); increaseItemQuantityErr != nil {
		return cartItemController.BadRequest(c, increaseItemQuantityErr)
	}
	return cartItemController.Created(c, nil, "")
}

func (cartItemController *CartItemController) DecreaseItemQuantity(c echo.Context) error {
	id, err := cartItemController.ParseIdParam(c, "id")
	if err != nil {
		return cartItemController.BadRequest(c, err)
	}
	queryParam := cartItemController.StringQueryParam(c, "amount")
	amount, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		return cartItemController.BadRequest(c, queryParamErr)
	}

	if increaseItemQuantityErr := cartItemController.cartItemService.DecreaseItemQuantity(id, amount); increaseItemQuantityErr != nil {
		return cartItemController.BadRequest(c, increaseItemQuantityErr)
	}
	return cartItemController.Created(c, nil, "")
}

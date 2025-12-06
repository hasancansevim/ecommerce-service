package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	cartService service.ICartService
	BaseController
}

func NewCartController(cartService service.ICartService) *CartController {
	return &CartController{cartService: cartService}
}

func (cartController *CartController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/carts/:id", cartController.GetCartById)
	e.GET("/api/v1/carts", cartController.GetCartsByUserId)
	e.POST("/api/v1/carts", cartController.CreateCart)
	e.DELETE("/api/v1/carts/:id", cartController.DeleteCartById)
	e.DELETE("/api/v1/carts/", cartController.ClearUserCarts)
}

func (cartController *CartController) GetCartById(c echo.Context) error {
	id, parseIdErr := cartController.BaseController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}

	getCartById := cartController.cartService.GetCartById(id)
	return cartController.Success(c, getCartById, "")
}

func (cartController *CartController) CreateCart(c echo.Context) error {
	var addCartRequest request.AddCartRequest
	bindErr := c.Bind(&addCartRequest)
	if bindErr != nil {
		return bindErr
	}

	cart, serviceErr := cartController.cartService.CreateCart(addCartRequest.ToModel())
	if serviceErr != nil {
		return serviceErr
	}
	return cartController.Created(c, cart, "")
}

func (cartController *CartController) GetCartsByUserId(c echo.Context) error {
	userId, parseIdErr := cartController.ParseIdParam(c, "user_id")
	if parseIdErr != nil {
		return parseIdErr
	}

	getCartByUserId := cartController.cartService.GetCartsByUserId(userId)
	return cartController.Success(c, getCartByUserId, "")
}

func (cartController *CartController) DeleteCartById(c echo.Context) error {
	id, parseIdErr := cartController.ParseIdParam(c, "id")
	if parseIdErr != nil {
		return parseIdErr
	}
	if serviceErr := cartController.cartService.DeleteCartById(id); serviceErr != nil {
		return serviceErr
	}

	return cartController.Created(c, nil, "")
}

func (cartController *CartController) ClearUserCarts(c echo.Context) error {
	userId, parseIdErr := cartController.ParseIdParam(c, "user_id")
	if parseIdErr != nil {
		return parseIdErr
	}
	if serviceErr := cartController.cartService.ClearUserCart(userId); serviceErr != nil {
		return serviceErr
	}

	return cartController.Created(c, nil, "Kullanıcının Sepeti Silindi")
}

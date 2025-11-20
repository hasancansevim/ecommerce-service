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
	id, err := cartController.BaseController.ParseIdParam(c, "id")
	if err != nil {
		return cartController.BadRequest(c, err)
	}

	getCartById := cartController.cartService.GetCartById(id)
	return cartController.Success(c, getCartById, "")
}

func (cartController *CartController) CreateCart(c echo.Context) error {
	var addCartRequest request.AddCartRequest
	bindErr := c.Bind(&addCartRequest)
	if bindErr != nil {
		return cartController.BadRequest(c, bindErr)
	}

	if toModelErr := cartController.cartService.CreateCart(addCartRequest.ToModel()); toModelErr != nil {
		return cartController.BadRequest(c, toModelErr)
	}

	return cartController.Created(c, addCartRequest, "")
}

func (cartController *CartController) GetCartsByUserId(c echo.Context) error {
	userId, err := cartController.ParseIdParam(c, "user_id")
	if err != nil {
		return cartController.BadRequest(c, err)
	}

	getCartByUserId := cartController.cartService.GetCartsByUserId(userId)
	return cartController.Success(c, getCartByUserId, "")
}

func (cartController *CartController) DeleteCartById(c echo.Context) error {
	id, err := cartController.ParseIdParam(c, "id")
	if err != nil {
		return cartController.BadRequest(c, err)
	}
	if deleteCartByIdErr := cartController.cartService.DeleteCartById(id); deleteCartByIdErr != nil {
		return cartController.BadRequest(c, deleteCartByIdErr)
	}

	return cartController.Created(c, nil, "")
}

func (cartController *CartController) ClearUserCarts(c echo.Context) error {
	userId, err := cartController.ParseIdParam(c, "user_id")
	if err != nil {
		return cartController.BadRequest(c, err)
	}
	if clearUserCartErr := cartController.cartService.ClearUserCart(userId); clearUserCartErr != nil {
		return cartController.BadRequest(c, clearUserCartErr)
	}

	return cartController.Created(c, nil, "Kullanıcının Sepeti Silindi")
}

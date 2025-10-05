package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/controller/response"
	"go-ecommerce-service/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartController struct {
	cartService service.ICartService
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
	queryParam := c.Param("id")
	id, convertErr := strconv.Atoi(queryParam)
	if convertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: convertErr.Error(),
		})
	}

	getCartById := cartController.cartService.GetCartById(int64(id))
	return c.JSON(http.StatusOK, response.ToResponseCartData(getCartById))
}

func (cartController *CartController) CreateCart(c echo.Context) error {
	var addCartRequest request.AddCartRequest
	bindErr := c.Bind(&addCartRequest)
	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}

	toModelErr := cartController.cartService.CreateCart(addCartRequest.ToModel())
	if toModelErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: toModelErr.Error(),
		})
	}
	return c.NoContent(http.StatusCreated)
}

func (cartController *CartController) GetCartsByUserId(c echo.Context) error {
	queryParam := c.QueryParam("user_id")
	user_id, queryParamErr := strconv.Atoi(queryParam)
	if queryParamErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: queryParamErr.Error(),
		})
	}

	getCartByUserId := cartController.cartService.GetCartsByUserId(int64(user_id))
	return c.JSON(http.StatusOK, response.ToResponseListCarts(getCartByUserId))
}

func (cartController *CartController) DeleteCartById(c echo.Context) error {
	param := c.Param("id")
	id, paramConvertErr := strconv.Atoi(param)
	if paramConvertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: paramConvertErr.Error(),
		})
	}
	deleteCartByIdErr := cartController.cartService.DeleteCartById(int64(id))
	if deleteCartByIdErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: deleteCartByIdErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

func (cartController *CartController) ClearUserCarts(c echo.Context) error {
	queryParam := c.QueryParam("user_id")
	user_id, queryParamConvertErr := strconv.Atoi(queryParam)
	if queryParamConvertErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: queryParamConvertErr.Error(),
		})
	}
	clearUserCartErr := cartController.cartService.ClearUserCart(int64(user_id))
	if clearUserCartErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: clearUserCartErr.Error(),
		})
	}
	return c.NoContent(http.StatusNoContent)
}

package controller

import (
	"go-ecommerce-service/controller/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BaseController struct{}

func (bc *BaseController) ParseIdParam(c echo.Context, paramName string) (int64, error) {
	param := c.Param(paramName)
	return strconv.ParseInt(param, 10, 64)
}

func (bc *BaseController) StringQueryParam(c echo.Context, paramName string) string {
	queryParam := c.QueryParam(paramName)
	return queryParam
}

func (bc *BaseController) Success(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func (bc *BaseController) Created(c echo.Context, data interface{}, message string) error {
	return c.JSON(http.StatusCreated, response.ApiResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func (bc *BaseController) BadRequest(c echo.Context, err error) error {
	return c.JSON(http.StatusBadRequest, response.ApiResponse{
		Success: false,
		Message: err.Error(),
		Data:    nil,
	})
}

func (bc *BaseController) InternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusInternalServerError, response.ApiResponse{
		Success: false,
		Message: "Internal Server Error: " + err.Error(),
		Data:    nil,
	})
}

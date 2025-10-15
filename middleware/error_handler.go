package middleware

import (
	_errors "go-ecommerce-service/pkg/errors"
	"go-ecommerce-service/service/validation"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorHandler() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				return handleError(c, err)
			}
			return nil
		}
	}
}

func handleError(c echo.Context, err error) error {
	if appErr, ok := err.(*_errors.AppError); ok {
		return c.JSON(getHTTPStatus(appErr.Code), appErr)
	}

	if validationErr, ok := err.(*validation.ValidationError); ok {
		return c.JSON(http.StatusBadRequest, validationErr)
	}

	return c.JSON(http.StatusInternalServerError, map[string]string{
		"error": "Internal server error",
	})
}

func getHTTPStatus(errorCode string) int {
	switch errorCode {
	case "USER_NOT_FOUND", "INVALID_CREDENTIALS":
		return http.StatusNotFound

	case "USER_ALREADY_EXISTS":
		return http.StatusConflict

	case "INVALID_TOKEN":
		return http.StatusUnauthorized

	default:
		return http.StatusInternalServerError
	}
}

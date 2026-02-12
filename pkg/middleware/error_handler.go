package middleware

import (
	_errors "go-ecommerce-service/pkg/errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var response interface{} = "Internal Server Error"

	switch e := err.(type) {

	case *_errors.AppError:
		code = e.Code
		response = map[string]interface{}{
			"code":    e.Code,
			"message": e.Message,
		}

	case *echo.HTTPError:
		code = e.Code
		response = map[string]interface{}{
			"code":    e.Code,
			"message": e.Message,
		}

	default:
		response = map[string]interface{}{
			"message": "An unknown error occurred",
			"details": err.Error(),
		}
	}

	if !c.Response().Committed {
		c.JSON(code, response)
	}
}

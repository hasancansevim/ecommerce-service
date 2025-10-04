package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/controller/response"
	"go-ecommerce-service/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (userController *UserController) RegisterRoutes(e *echo.Echo) {
	e.GET("/api/v1/users", userController.GetAllProducts)
	e.POST("/api/v1/users", userController.AddProduct)
}

func (userController *UserController) GetAllProducts(c echo.Context) error {
	allUsers := userController.userService.GetAllUsers()

	return c.JSON(http.StatusOK, allUsers)
}

func (userController *UserController) AddProduct(c echo.Context) error {
	var addUserRequest request.AddUserRequest
	bindErr := c.Bind(&addUserRequest)

	if bindErr != nil {
		return c.JSON(http.StatusBadRequest, response.ErrorResponse{
			ErrorDescription: bindErr.Error(),
		})
	}

	addUserErr := userController.userService.AddUser(addUserRequest.ToModel())
	if addUserErr != nil {
		return c.JSON(http.StatusInternalServerError, response.ErrorResponse{
			ErrorDescription: addUserErr.Error(),
		})
	}

	return c.NoContent(http.StatusCreated)
}

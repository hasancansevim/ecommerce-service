package controller

import (
	"go-ecommerce-service/controller/request"
	"go-ecommerce-service/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService service.IUserService
	BaseController
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

	return userController.Success(c, allUsers)
}

func (userController *UserController) AddProduct(c echo.Context) error {
	var addUserRequest request.AddUserRequest
	bindErr := c.Bind(&addUserRequest)

	if bindErr != nil {
		return userController.BadRequest(c, bindErr)
	}

	if addUserErr := userController.userService.AddUser(addUserRequest.ToModel()); addUserErr != nil {
		return userController.BadRequest(c, addUserErr)
	}
	return userController.Created(c)
}

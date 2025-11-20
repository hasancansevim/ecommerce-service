package controller

import (
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
	e.GET("/api/v1/users", userController.GetAllUsers)
}

func (userController *UserController) GetAllUsers(c echo.Context) error {
	allUsers := userController.userService.GetAllUsers()

	return userController.Success(c, allUsers, "Tüm Kullanıcılar Getirildi")
}

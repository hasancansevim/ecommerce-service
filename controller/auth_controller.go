package controller

import (
	"go-ecommerce-service/controller/request"
	_interface "go-ecommerce-service/service/interface"
	"go-ecommerce-service/service/model"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService _interface.AuthService
	BaseController
}

func NewAuthController(authService _interface.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (authController *AuthController) RegisterRoutes(e *echo.Echo) {
	e.POST("/api/v1/auth/register", authController.Register)
	e.POST("/api/v1/auth/login", authController.Login)
}

func (authController *AuthController) Register(c echo.Context) error {
	var registerRequest request.RegisterRequest
	if bindErr := c.Bind(&registerRequest); bindErr != nil {
		return bindErr
	}

	registerRequestModel := model.RegisterCreate{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Email:     registerRequest.Email,
		Password:  registerRequest.Password,
	}

	serviceErr := authController.authService.Register(registerRequestModel)
	if serviceErr != nil {
		return serviceErr
	}
	return authController.Created(c, registerRequestModel, "Kayıt Başarılı")
}

func (authController *AuthController) Login(c echo.Context) error {
	var loginRequest request.LoginRequest
	if bindErr := c.Bind(&loginRequest); bindErr != nil {
		return bindErr
	}
	token, serviceErr := authController.authService.Login(model.LoginCreate{
		Email:    loginRequest.Email,
		Password: loginRequest.Password,
	})
	if serviceErr != nil {
		return serviceErr
	}
	return authController.Success(c, token, "Giriş Başarılı")
}

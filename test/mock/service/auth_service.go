package service

import (
	"go-ecommerce-service/service/model"

	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (authService *MockAuthService) Register(registerModel model.RegisterCreate) error {
	args := authService.Called(registerModel)
	return args.Error(0)
}

func (authService *MockAuthService) Login(loginModel model.LoginCreate) (string, error) {
	args := authService.Called(loginModel)
	return args.String(0), args.Error(1)
}

func (authService *MockAuthService) ValidateToken(token string) (jwt2.Claims, error) {
	args := authService.Called(token)
	return args.Get(0).(jwt2.Claims), args.Error(1)
}

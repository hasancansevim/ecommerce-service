package _interface

import (
	"go-ecommerce-service/service/model"
)

type AuthService interface {
	Register(registerModel model.RegisterCreate) error
	Login(loginModel model.LoginCreate) (string, error)
	ValidateToken(token string) (int64, error) // return userId
}

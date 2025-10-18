package service

import (
	"go-ecommerce-service/internal/jwt"
	_interface "go-ecommerce-service/service/interface"

	jwt2 "github.com/golang-jwt/jwt/v4"
)

type JWTService struct {
}

func NewJWTService() _interface.JWTManager {
	return &JWTService{}
}

func (j *JWTService) GenerateToken(userId int64, email string) (string, error) {
	return jwt.GenerateToken(userId, email)
}

func (j *JWTService) ValidateToken(token string) (jwt2.Claims, error) {
	return jwt.ValidateToken(token)
}

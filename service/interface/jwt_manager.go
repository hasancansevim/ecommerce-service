package _interface

import (
	jwt2 "github.com/golang-jwt/jwt/v4"
)

type JWTManager interface {
	GenerateToken(userId int64, email string) (string, error)
	ValidateToken(token string) (jwt2.Claims, error)
}

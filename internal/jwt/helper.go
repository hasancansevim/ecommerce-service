package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret []byte

func Initialize(secret string) {
	jwtSecret = []byte(secret)
}

type Claim struct {
	UserId int64  `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(userId int64, email string) (string, error) {
	claim := &Claim{
		UserId: userId,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claim, error) {
	claim := &Claim{}
	token, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	return claim, nil
}

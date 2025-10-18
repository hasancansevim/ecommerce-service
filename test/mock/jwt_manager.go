package mock

import (
	jwt2 "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/mock"
)

type MockJWTManager struct {
	mock.Mock
}

func (m *MockJWTManager) GenerateToken(userId int64, email string) (string, error) {
	args := m.Called(userId, email)
	return args.String(0), args.Error(1)
}

func (m *MockJWTManager) ValidateToken(token string) (jwt2.Claims, error) {
	args := m.Called(token)
	if args.Get(0) != nil {
		return args.Get(0).(jwt2.Claims), args.Error(1)
	}
	return nil, args.Error(1)
}

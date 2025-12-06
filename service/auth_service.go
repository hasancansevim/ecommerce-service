package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/auth"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
	_interface "go-ecommerce-service/service/interface"
	"go-ecommerce-service/service/model"

	jwt2 "github.com/golang-jwt/jwt/v4"
)

type AuthService struct {
	userRepository persistence.IUserRepository
	jwtManager     _interface.JWTManager
}

func NewAuthService(userRepository persistence.IUserRepository, jwtManager _interface.JWTManager) _interface.AuthService {
	return &AuthService{
		userRepository: userRepository,
		jwtManager:     jwtManager,
	}
}

func (authService *AuthService) Register(registerModel model.RegisterCreate) error {

	existingUser, getUserByEmailErr := authService.userRepository.GetUserByEmail(registerModel.Email)
	if getUserByEmailErr == nil && existingUser.Email != "" {
		return _errors.NewBadRequest(getUserByEmailErr.Error())
	}

	passwordHash, passwordErr := auth.HashPassword(registerModel.Password)
	if passwordErr != nil {
		return _errors.NewBadRequest(passwordErr.Error())
	}
	return authService.userRepository.CreateUser(domain.User{
		FirstName:    registerModel.FirstName,
		LastName:     registerModel.LastName,
		Email:        registerModel.Email,
		PasswordHash: passwordHash,
	})
}

func (authService *AuthService) Login(loginModel model.LoginCreate) (string, error) {
	userByEmail, userByEmailErr := authService.userRepository.GetUserByEmail(loginModel.Email)
	if userByEmailErr != nil {
		return "", _errors.NewBadRequest(userByEmailErr.Error())
	}
	checkPasswordHash := auth.CheckPasswordHash(loginModel.Password, userByEmail.PasswordHash)
	if checkPasswordHash == false {
		return "", _errors.NewBadRequest("Password Error")
	}
	token, tokenErr := authService.jwtManager.GenerateToken(userByEmail.Id, userByEmail.Email)
	if tokenErr != nil {
		return "", _errors.NewBadRequest(tokenErr.Error())
	}
	return token, nil
}

func (authService *AuthService) ValidateToken(token string) (jwt2.Claims, error) {
	claim, err := authService.jwtManager.ValidateToken(token)
	if err != nil {
		return nil, _errors.NewBadRequest(err.Error())
	}
	return claim, nil
}

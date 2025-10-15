package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/internal/auth"
	"go-ecommerce-service/internal/jwt"
	"go-ecommerce-service/persistence"
	_errors "go-ecommerce-service/pkg/errors"
	_interface "go-ecommerce-service/service/interface"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/service/validation"
)

type AuthService struct {
	userRepository persistence.IUserRepository
}

func NewAuthService(userRepository persistence.IUserRepository) _interface.AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (authService *AuthService) Register(registerModel model.RegisterCreate) error {
	if validationErr := validation.ValidateRegisterModel(registerModel); validationErr != nil {
		return validationErr
	}

	existingUser, getUserByEmailErr := authService.userRepository.GetUserByEmail(registerModel.Email)
	if getUserByEmailErr == nil && existingUser.Email != "" {
		return _errors.UserAlreadyExists
	}

	passwordHash, err := auth.HashPassword(registerModel.Password)
	if err != nil {
		return _errors.InternalServerError
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
		return "", _errors.InternalServerError
	}
	checkPasswordHash := auth.CheckPasswordHash(loginModel.Password, userByEmail.PasswordHash)
	if checkPasswordHash == false {
		return "", _errors.InvalidCredentials
	}
	token, userByEmailErr := jwt.GenerateToken(userByEmail.Id, userByEmail.Email)
	if userByEmailErr != nil {
		return "", _errors.InternalServerError
	}
	return token, nil
}

func (authService *AuthService) ValidateToken(token string) (int64, error) {
	claim, err := jwt.ValidateToken(token)
	if err != nil {
		return 0, err
	}
	return claim.UserId, _errors.InternalServerError
}

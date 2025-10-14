package service

import (
	"go-ecommerce-service/domain"
	"go-ecommerce-service/persistence"
	"go-ecommerce-service/service/model"
	"go-ecommerce-service/service/validation"
	"time"
)

type IUserService interface {
	GetAllUsers() []domain.User
	CreateUser(userCreate model.UserCreate) error
	GetUserByEmail(email string) (domain.User, error)
	GetUserById(id int64) (domain.User, error)
}

type UserService struct {
	userRepository persistence.IUserRepository
}

func NewUserService(userRepository persistence.IUserRepository) IUserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) GetAllUsers() []domain.User {
	return userService.userRepository.GetAllUser()
}

func (userService *UserService) CreateUser(userCreate model.UserCreate) error {
	if validationError := validation.ValidateUserCreate(userCreate); validationError != nil {
		return validationError
	}

	return userService.userRepository.CreateUser(domain.User{
		FirstName:    userCreate.FirstName,
		LastName:     userCreate.LastName,
		Email:        userCreate.Email,
		PasswordHash: userCreate.PasswordHash,
		CreatedAt:    time.Now(),
	})
}

func (userService *UserService) GetUserById(id int64) (domain.User, error) {
	getUserById, err := userService.userRepository.GetUserById(id)
	if err != nil {
		return domain.User{}, err
	}
	return getUserById, nil
}

func (userService *UserService) GetUserByEmail(email string) (domain.User, error) {
	getUserByEmail, err := userService.userRepository.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	return getUserByEmail, nil
}
